pragma solidity ^0.4.10;

// ERC20 标准: 各个代币的标准接口
//  获得代币总供应量
//  账户余额
//  代币转账(转入， 转出）
//  批准代币花费

contract Token {
    uint256 public totalSupply;
    function balanceOf(address owner) constant returns (uint256 balance); // 余额
    function transferTo(address to, uint256 amount) returns (bool success); // 转出
    function transferFrom(address from, address to, uint256 amount) returns (bool success); // 转入至
    function approve(address spender, uint256 amount) returns (bool success); // 批准花费
    function allowance(address owner, address spender) constant returns (uint256 remaining); // 补贴

    event Transfer(address indexed from, address indexed to, uint256 amount);
    event Approval(address indexed owner, address indexed spender, uint256 amount);
}

// 合约标准
contract StandardToken is Token {
    mapping (address => uint256) balances;
    mapping (address => mapping (address =>uint256)) allowed;

    function transferTo(address to, uint256 amount) returns (bool success) {
        if (balances[msg.sender] >= amount && amount > 0) {
            balances[msg.sender] -= amount;
            balances[to] += amount;
            Transfer(msg.sender, to, amount); //触发转账
            return true;
        } else {
            return false;
        }
    }

    function transferFrom(address from, address to, uint256 amount) returns (bool success) {
        if (balances[from] >= amount && allowed[from][msg.sender] >= amount && amount > 0 ) {
            balances[from] -= amount;
            balances[to] += amount;
            allowed[from][msg.sender] -= amount;
            Transfer(from, to, amount)
            return true;
        } else {
            return false;
        }
    }

    function balance(address owner) constant returns (uint256 balance) {
        return balances[owner];
    }

    function approve(address spender, uint256 amount) return (bool success) {
        allowed[msg.sender][spender] = amount;
        Approval(msg.sender, spender, amount);
        return true;
    }

    function allowance(address owner, address spender) constant returns (uint256 remaining) {
        return allowed[owner][spender];
    }
}
