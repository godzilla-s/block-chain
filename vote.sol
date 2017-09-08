// 选举

contract Vote {
  address public voteorg
  struct vote{
    bool  vote;
    bool  isVote;
  }
  
  mapping (address => vote) public voteinfo;
  
  function Vote() {
    voteorg = msg.sender;
  }
}
