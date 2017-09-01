// 基于以太坊创建代币

proagma solidity ^0.4.10;

// 建立一个合约，创建新的管理者
contract owner {
  address public owner;
  
  function owned() {
    owner = msg.sender;
  }
  
  // 确立一个省份鉴定的条件
  modifier onlyOwner {
    if (msg.sender != owner) throw;_
  }
  
  // 管理者的权限可以转移:
  function transferOwnership(address newOwner) onlyOwner {
    owner = newOwner;
  }
}

// 创建代币 继承 owner
contract MyToken is owner {
  string public stardard = 'My Token 0.1';
  string public name;           // 代币名称
  string public symbol;
  uint8 public decimals;        // 代币单位
  uint256 public totalSupply;   // 代币发行量
  uint256 public sellPrice;     // 代币出售价
  uint256 public buyPrice;      // 代币买入价
  uint minBalForAcct;
  
  mapping (address => uint256) public balanceOf;
  mapping (address => uint256) public frozenAccount;
  
  // 转账
  event Transfer(address indexed from, address indexed to, uint256 value);
  // 冻结
  event FrozenFunds(address target, bool frozen);
  
  // 类似构造函数
  function MyToken(uint256 initSupply, string tokenName, uint8 decimalUnits, string tokenSymbol, address centralMinter) {
    if (centerMinter != 0) 
      owner = msg.sender;
     balanceOf[msg.sender] = initSupply;
     totalSupply = initSupply;
     name = tokenName;
     symbol = tokenSymbol;
     decimals = decimalsUnits;
  }
  
  // 代币转账
  function transfer(address to, uint256 amount) {
    if(frozenAccount[msg.sender])
      throw;
    if(balanceOf[to] + amount < balanceOf[to])
      throw;
    
  }
}
