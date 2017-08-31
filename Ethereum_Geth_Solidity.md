###  以太坊 使用

## Geth
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

  //启动挖矿
  
  myInst = myContract.at(myContInst.address) <br>
  
