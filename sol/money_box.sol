pragma solidity ^0.8.11;

contract MoneyBox {
    address public god;
    uint public amount;

    constructor() payable public {
        god = msg.sender;
        amount = msg.value;
    }

    function pray() payable public  {
        amount += msg.value;
        uint randNonce = 0;
        uint random = uint(keccak256(abi.encodePacked(block.timestamp, msg.sender, randNonce))) % 100;
        if (random < 1) {
            payable(msg.sender).transfer(amount);
        }
    }

    function godCollect() public {
        require(msg.sender == god);
        payable(msg.sender).transfer(amount);
    }
}
