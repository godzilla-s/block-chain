pragma solidity ^0.4.10;

// 区块与交易属性
contract Block_Txn_Attr {
    // 获取区块hash
    function getBlockHash() returns (bytes32) {
        return block.blockhash(block.number-1);
    }
    
    // 获取未使用gas
    function getUnusedGas() returns (uint) {
        return msg.gas;
    }
    
    // 获取当前地址
    function getCurrAddr() returns (address) {
        return msg.sender;
    }
    
    // 获取发送者余额 注意 payable 修饰
    function getSenderBalance() payable returns (uint) {
        return msg.value;
    }
    
    // 获取交易的gas价格
    function getTxPrices() returns (uint) {
        return tx.gasprice;
    }  
}
