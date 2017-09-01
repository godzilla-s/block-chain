// 简单代币实现

contract MyToken {
  mapping (address => uint256) public balanceOf; // 设置一个mapping存储每个账户的代币信息
  
  string public name;   // 代币名称
  string public symbol; // 代币图案
  uint8 public decimals; // 代币小数点
  
  event Transfer(address indexed from, address indexed to, uint256 amount);
  
  function MyToken(uint256 initSupply, string tokenName, uint8 decimalUnits, string tokenSymbol) {
    balanceOf[msg.sender] = initSupply;
    name = tokenName;
    symbol = tokenSymbol;
    decimals = decimalsUnits;
  }
  
  // 代币交易
  function transfer(address to, uint256 amount) {
    if(balanceOf[msg.sender] < amount || balanceOf[to] + amount < balanceOf[to])
      throw;
     
     balanceOf[msg.sender] -= amount;
     balanceOf[to] += amount;
     
     Transfer(msg.sender, to, amount);
  }
}
