pragma solidity ^0.8.11;

contract MoneyBox {
    address public god;
    uint randNonce;

    constructor() payable public {
        require(msg.value > 0);
        god = msg.sender;
        randNonce = block.timestamp;
    }

    function pray() payable public  {
        uint random = uint(keccak256(abi.encodePacked(block.timestamp, msg.sender, randNonce))) % 10;
        if (random < 2) {
            payable(msg.sender).transfer(address(this).balance);
        }
    }

    function balance() public view returns (uint256) {
        return address(this).balance;
    }

    function godCollect() public {
        require(msg.sender == god);
        selfdestruct(payable(god));
    }

}
