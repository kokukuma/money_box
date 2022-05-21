pragma solidity ^0.8.11;

contract MoneyBox {
    address public god;
    uint randNonce;

    event Pray(address prayer, uint amount, bool rewarded);

    constructor() payable public {
        require(msg.value > 0);
        god = msg.sender;
        randNonce = block.timestamp;
    }

    function pray() payable public  {
        uint random = uint(keccak256(abi.encodePacked(block.timestamp, msg.sender, randNonce))) % 10;
        bool rewarded = random < 2;
        if (rewarded) {
            payable(msg.sender).transfer(address(this).balance);
        }
        emit Pray(msg.sender, msg.value, rewarded);
    }

    function balance() public view returns (uint256) {
        return address(this).balance;
    }

    function godCollect() public {
        require(msg.sender == god);
        selfdestruct(payable(god));
    }

}
