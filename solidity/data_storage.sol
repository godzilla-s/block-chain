// 数据存储
// 数据存储主要两种: memory, storage

pragma solidity ^0.4.0;

contract Storage {
  uint  value_01;   // memory 存储
  mapping (uint => uint) public value_02; // storage 存储
  
  function setMemoryData() {
    value_01 = 80;
  }
  
  function setStorageData() {
    value_02[1] = 60;
  }
  
  function changeMemoryData(uint a) {
    var temp = value_01;
    temp = a;
  }
  
  function changeStorageData(uint a) {
    var temp = value_02;
    value02[0] = a;
  }
  
  function getValue() returns(uint, uint) {
    return (value_01, value_02[1]);
  }
}


// 总结:
// 函数参数（括返回参数：memory
// 所有其它的局部变量：storage
