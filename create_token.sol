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
  uint minBalForAccount;           // 账户最低余额(私有)
  
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
    // 判断是否冻结
    if(frozenAccount[msg.sender])
      throw;
      
    // 判断是否足够转出
    if(balanceOf[msg.sender] < amount) 
      throw;
    
    // ??
    if(balanceOf[to] + amount < balanceOf[to])
      throw;
    
    Transfer(msg.sender, to, amount);
  }
  
  // 买入
  function buy() returns (uint amount) {
    amount = msg.value / buyPrice;  // 买入数量, 以当前汇率来计算
    if (balanceOf[this] < amount) // 查看是否有足够的出售
      throw;    
    balanceOf[msg.sender] += amount; // 购买者增加
    balanceOf[this] -= amount;   // 总量减少
    
    Transfer(this, msg.sender, amount); // 触发转账事件
    return amount;
  }
  
  // 卖出
  function sell(uint amount) returns (uint revenue) {
    if (balanceOf[msg.sender] < amount)
      throw;
    balanceOf[this] += amount;
    balanceOf[msg.sender] -= amount;
    
    revenue = amount * sellPrice; // 计算卖出总量
    msg.sender.send(revenue); // 用户获得因为输出代币得到的以太币 ？
    Transfer(msg.sender, this, amount);
    return revenue;
  }
  
  // 设置买入，卖出汇率
  function setPrice(uint256 newBuyPrice, uint256 newSellPrice) onlyOwner {
    sellPrice = newSellPrice;
    buyPrice = newBuyPrice;
  }
  
  function setMinBalance(uint256 newMinBalance) onlyOwner {
    minBalForAccount = newMinBalance * 1 finney;
  }
}
