function testRandomWord() public {
    // Get the owner's private key from an environment variable.
    uint256 ownerKey = vm.envUint("PRIVATE_KEY");
    // Start broadcasting from the owner's account.
    vm.startBroadcast(ownerKey);
    uint256 requestId = consumer.requestRandomNumber();
    vm.stopBroadcast();

    // Now, after the randomness is (ideally) fulfilled,
    // you can fetch and log the stored random number.
    uint256 randomWord = consumer.s_randomWord();
    console.log("Random Word:", randomWord);
    
    // Optionally, add your assertion here based on your expected logic.
} 