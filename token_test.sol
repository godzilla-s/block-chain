pragma solidity ^0.4.10;

contract Owner {
    address public owner;

    function owner() {
        owner = msg.sender;
    }
    
    modifier ownerOnly {
        require(msg.sender == owner);
        _;
    }
    
    function transferOwnership(address newOwner) {
        owner = newOwner;
    } 
}

contract Token {
    uint256 public initSupply;
    uint8 public decimal;
    string public tokenName;
    string public tokenSymbol;
    
    uint public sellPrice;
    uint public buyPrice;
    
    mapping (address => uint256) public balanceOf;
    mapping (address => bool) public frozenAccount;
    mapping (address => mapping(address => uint256)) public allowence;
    
    event Transfer(address indexed _from, address indexed _to, uint amount);
    event FrozenFunds(address indexed _target, bool _frozen);
    
    function Token(uint256 _initSupply, string _tokenName, string _tokenSymbol, uint8 _decimal) {
        initSupply = _initSupply;
        balanceOf[msg.sender] = _initSupply;
        tokenName = _tokenName;
        tokenSymbol = _tokenSymbol;
        decimal = _decimal;
    }
    
    // test
    function getAddr() payable returns (address, address, uint){
        return (msg.sender, this, msg.value);
    }
    
    function _transfer(address _from, address _to, uint256 _amount) internal {
        require(_to != 0x00);
        require(balanceOf[_from] > _amount);
        require(balanceOf[_to] + _amount > balanceOf[_to]);
        require(!frozenAccount[_from]);
        require(!frozenAccount[_to]);
        balanceOf[_from] -= _amount;
        balanceOf[_to] += _amount;
        Transfer(_from, _to, _amount);
    }
    
    function transfer(address _to, uint256 _amount) returns (bool)  {
        _transfer(msg.sender, _to, _amount);
        return true;
    }
    
    function transferFrom(address _from, address _to, uint256 _amount) returns (uint256, bool) {
        require(allowence[msg.sender][_from] > 0 && _amount < allowence[msg.sender][_from]);
        allowence[msg.sender][_from] -= _amount;
        _transfer(_from, _to, _amount);
        return (allowence[msg.sender][_from], true);
    }
    
    function _approve(address _owner, address _target, uint256 _amount) internal {
        require(_target != 0x00);
        allowence[_owner][_target] = _amount;
    }
    
    function approve(address _target, uint256 _amount) returns (bool) {
        _approve(msg.sender, _target, _amount);
        return true;
    }
    
    function frozenAccount(address _target, bool _frozen) returns (bool) {
        frozenAccount[_target] = _frozen;
        FrozenFunds(_target, _frozen);
        return true;
    }
    
    function mintToken(address _target, uint256 _mintAmount) {
        balanceOf[_target] += _mintAmount;
        initSupply += _mintAmount;
        Transfer(0, this, _mintAmount);
        Transfer(this, _target, _mintAmount);
    }
    
    function setPrice(uint _sellPrice, uint _sellPrice) {
        buyPrice = _buyPrice;
        sellPrice = _sellPrice;
    }
    
    // 需先充钱到合约账户里面去  也就是 this 账户
    function buy() payable {
        require(buyPrice > 0);
        uint amount = msg.value / buyPrice;  
        _transfer(this, msg.sender, amount);
    }
    
    function sell(uint _amount) {
        require(sellPrice > 0);
        require(this.balance >= _amount * sellPrice);
        _transfer(msg.sender, this, _amount);
        msg.sender.transfer(amount * sellPrice);
    }
    
}
