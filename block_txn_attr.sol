pragma solidity ^0.4.10;

// 区块与交易属性
contract Block_Txn_Attr {
    // 获取当前矿工的地址
    function getCoinbase() returns (address) {
        return block.coinbase;
    }
    
    // 当前块的难度
    function getCurrDifficulty()  returns (uint) {
        return block.difficulty;
    }
    
    // 当前块的gaslimit
    function getCurrGaslimit() returns (uint) {
        return block.gaslimit;
    }
    
    // 完整的调用数据（calldata）
    function getMsgData() returns (bytes) {
        return msg.data;
    }
 
    // 获取区块hash
    function getBlockHash() returns (bytes32) {
        return block.blockhash(block.number-1);
    }
    
    // 获取未使用gas
    function getUnusedGas() returns (uint) {
        return msg.gas;
    }
    
    // 获取当前地址
    function getMsgSender() returns (address) {
        return msg.sender;
    }
    
    // 这个消息所附带的货币量，单位为wei
    function getMsgValue() payable returns (uint) {
        return msg.value;
    }
    
    // 获取交易的gas价格
    function getTxPrices() returns (uint) {
        return tx.gasprice;
    }  
}
