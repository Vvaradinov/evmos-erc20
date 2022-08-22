pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

// This contract creates a new token called VladToken with the symbol VLAD.
contract VladToken is ERC20 {
    constructor() ERC20("Vladcoin", "VLAD") {
        _mint(msg.sender, 10000000000000000000000000000);
    }
}
