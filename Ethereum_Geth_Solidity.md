##  以太坊相关

### Geth
==================
关于geth 1.6 没法使用solidity 内嵌的编译 <br>
解决方法 <br>
1. 生产 bytecode: <br>
  solc -o ./ --bin xxx.sol  <br>
2. 生产 abi: <br>
  solc -o ./ --abi xxx.sol  <br>

修改 xxx.bin  , xxx.abi <br>

3. 使用步骤 <br>
  loadScript("/path/to/xxx.bin") <br>
  loadScript("/path/to/xxx.abi") <br>
  myContract = eth.contract(abi) <br>
  myDeploy = {from:eth.coinbase, data: bytecode, gas: 1000000} <br>
  myContInst = myContract.new(txDeploy) <br>

  //启动挖矿 <br>
  
  myInst = myContract.at(myContInst.address) <br>
  
  //调用 <br>
  myInst.<function>.call(argv, ....)  // function 函数名 <br>
  
参考： https://ethereum.stackexchange.com/questions/15435/how-to-compile-solidity-contracts-with-geth-v1-6/15436 <br>
      http://blog.csdn.net/jwter87/article/details/53445709 <br>
      http://blog.csdn.net/CHENYUFENG1991/article/details/53458175?locationNum=7&fps=1 <br>
      

### 关于solc 使用
==================

参考： https://zhuanlan.zhihu.com/p/27889205 <br>
      Solidity 简易教程: http://wiki.jikexueyuan.com/project/solidity-zh/introduction.html <br> 


### ERC20 标准

参考： ERC20 ： https://github.com/ethereum/EIPs/issues/20 <br>
      标准代币接口： https://github.com/ethereum/wiki/wiki/Standardized_Contract_APIs#data-feeds


### 以太坊区块链 -- 学习
http://blog.csdn.net/fidelhl/article/category/6060944

区块链入门 ： http://www.cnblogs.com/zl03jsj/p/6819333.html

以太坊智能合约编程之菜鸟教程（简单例子） http://blog.csdn.net/xxxslinyue/article/details/70881030

以太坊研究： http://blog.csdn.net/DDFFR/article/category/6601839/1 
