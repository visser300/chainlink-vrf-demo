// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../src/RandomNumberConsumer.sol";

contract SepoliaIntegrationTest is Test {
    RandomNumberConsumer consumer;

    // Set up by reading the deployed contract address from an env variable.
    function setUp() public {
        // DEPLOYED_CONTRACT_ADDRESS must be set to your deployed contract's address on Sepolia.
        address deployedAddress = vm.envAddress("DEPLOYED_CONTRACT_ADDRESS");
        consumer = RandomNumberConsumer(deployedAddress);
    }

    // Example test: verify that the contract owner matches the expected deployer.
    function testOwnerIsDeployer() public {
        // DEPLOYER_ADDRESS must be set to the deployer's address when you deployed the contract.
        address expectedOwner = vm.envAddress("DEPLOYER_ADDRESS");
        assertEq(consumer.s_owner(), expectedOwner, "Owner does not match deployer");
    }

    // Example test: check the random number state (adjust as needed).  
    // Note: If no random number request has been fulfilled, the value may be 0.
    function testRandomWord() public {
        uint256 ownerKey = vm.envUint("PRIVATE_KEY"); // must match the deployer/owner
        address expectedOwner = vm.addr(ownerKey);
        assertEq(consumer.s_owner(), expectedOwner, "Deployed owner doesn't match expected owner");

        vm.startBroadcast(ownerKey);
        uint256 requestId = consumer.requestRandomNumber();
        vm.stopBroadcast();

        // Compute the expected random result based on VRFCoordinatorV2Mock logic.
        uint256 expectedRandomWord = uint256(keccak256(abi.encode(requestId, uint256(0))));
        assertEq(
            consumer.s_randomWord(),
            expectedRandomWord,
            "Random word not set correctly"
        );
        console.log("Random Word:", expectedRandomWord);
    }
} 