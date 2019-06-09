package genesis

// GenesisFNAccounts are the ECSDA accounts for the foundational nodes.
var GenesisFNAccounts = [...]DeployAccount{
	{Address: "0x35D29200aFC9A4cDC05166096059a042078CB53e", Public: "0x35D29200aFC9A4cDC05166096059a042078CB53e", BLSKey: "f82fcd7b5dc07f3dd0259701300bb6746a789f8e0e7d72ac696108e075ff89d"},
	{Address: "0xe4a69826534aD3f6ec6E432474B0380E7F9a9C3d", Public: "0xe4a69826534aD3f6ec6E432474B0380E7F9a9C3d", BLSKey: "8128924a101798a1c07679c21ad30f1c30bddc672f31bc3077e7cb0db24b86a"},

	// 0 - 9
	{Address: "0x04c3636dF766ad2d3E74424c016842f5704FAE3A", Public: "0x04c3636dF766ad2d3E74424c016842f5704FAE3A", BLSKey: "5837f197c5e01316d47e9e4d0207cb58f9a5addafe59122e81b3c21ea3f96fee"},
	{Address: "0x053515CC2CAae77F7e2F0A9C48A27c8f6D76E99d", Public: "0x053515CC2CAae77F7e2F0A9C48A27c8f6D76E99d", BLSKey: "6830748bd831edeb48c1c292923f279f2a81c5c382c77b9aebcfe63620704ac"},
	{Address: "0x0850243810E77fC6261965d2F163d36628E77E05", Public: "0x0850243810E77fC6261965d2F163d36628E77E05", BLSKey: "a30e9d8999787b5acfd72af1702f166d668022b62b5c44b56a882d4f60dfb81"},
	{Address: "0x08aB87F3A8EB0b69a833575B6400670f3F330302", Public: "0x08aB87F3A8EB0b69a833575B6400670f3F330302", BLSKey: "3a9825d9fca01395737165b872297774283738270bd5795da518fcc12dc006d4"},
	{Address: "0x0d51F2d1EB1716F30c6f72673a4A89a0A10cdf64", Public: "0x0d51F2d1EB1716F30c6f72673a4A89a0A10cdf64", BLSKey: "37e97aed614f5086f2fcee8c3152725440066c03f76e0d916a3e567056d31744"},
	{Address: "0x144B2Fd168147311f749B0f9573664676C333e2A", Public: "0x144B2Fd168147311f749B0f9573664676C333e2A", BLSKey: "2bae3aeb4653657c1b30e4204e8c093bbab8e3ba93b421325dc0a3f60e53549c"},
	{Address: "0x22117D26611161b1b1f4EBB06C441aeeA102261c", Public: "0x22117D26611161b1b1f4EBB06C441aeeA102261c", BLSKey: "588ec71b9999a4007663432e8e3fa50aae972ebe3bad89148f42e588cf101701"},
	{Address: "0x24A8cD56bABef297F1C7234F830362466d01ff5d", Public: "0x24A8cD56bABef297F1C7234F830362466d01ff5d", BLSKey: "50a40637369e7d9d704f063c2c818a53e628d087b8d5a57a9661f4017c5ba0"},
	{Address: "0x25347d09373B2644191f1DC4beDEFEBE26a5b2d1", Public: "0x25347d09373B2644191f1DC4beDEFEBE26a5b2d1", BLSKey: "72d41ddbfadf9b74f8ec4d6bf2c911b90325bf5bf6620ddd891c78fcf1b358d7"},
	{Address: "0x25441821ecA41DEc79578aAB866d3627A2e9BB9f", Public: "0x25441821ecA41DEc79578aAB866d3627A2e9BB9f", BLSKey: "570dce63c362a8222fd564516538c84c81fdbc3d8dea0bd11129e0080c9d6486"},

	// 10 - 19
	{Address: "0x27930D539fA8B118B5547a81Fd4cd0f0Fd295503", Public: "0x27930D539fA8B118B5547a81Fd4cd0f0Fd295503", BLSKey: "71ccfdcf700678bcdcec92c5dda72a90f242b48dda9e91b988dc9403011bed77"},
	{Address: "0x28085D40501df849246040Ea815fbD71F08c2fc4", Public: "0x28085D40501df849246040Ea815fbD71F08c2fc4", BLSKey: "3062268458934841a67e4d3ea5bc98756736194b62e33af0706e959125cd6296"},
	{Address: "0x28dA1beF8F5361863DcD427B6264f9DdF05B5D14", Public: "0x28dA1beF8F5361863DcD427B6264f9DdF05B5D14", BLSKey: "27af5f58cb6c97e823a86ad4c9a4109adec1dc41bd6c5bd035e4c5ff7021ac9a"},
	{Address: "0x2FB4584233B07d99ed7215c2E32dFCac8A2d5575", Public: "0x2FB4584233B07d99ed7215c2E32dFCac8A2d5575", BLSKey: "71d864883d2bd9286c41a5c4be282b6e6dcacb8cef75f6eca92f7141ba1eb7de"},
	{Address: "0x2b3234Ee92270A486a1598c5Bd74e739EC26fd9b", Public: "0x2b3234Ee92270A486a1598c5Bd74e739EC26fd9b", BLSKey: "5000fc83d8074a44a6e2adcb2aba814f9464f05c15f0fcdd90ee4bf102b0f665"},
	{Address: "0x2bC858D0967384C0093e12824Bb3d6486d51c30D", Public: "0x2bC858D0967384C0093e12824Bb3d6486d51c30D", BLSKey: "3a995a7f55096a2d0142106ca97d6d63cccaa62259f7e72b98629218212d1835"},
	{Address: "0x324c741430F5B970b61E398434B4F3957a6BC6E0", Public: "0x324c741430F5B970b61E398434B4F3957a6BC6E0", BLSKey: "e1158bd530e6f57f8404be1116dc90dc5d41d638be2b9c583a6ba500d9d98c6"},
	{Address: "0x3413e7e39eE7394b692FB04c12f5671d5Bb43e0b", Public: "0x3413e7e39eE7394b692FB04c12f5671d5Bb43e0b", BLSKey: "3d00e49943931598501e894c7a3c16349dad0ce4f8197150aa62cbf0bbef2c1a"},
	{Address: "0x386333bfe5Dbdb4c0b5633E71190F3F822b3C0bC", Public: "0x386333bfe5Dbdb4c0b5633E71190F3F822b3C0bC", BLSKey: "3bfd4ccee42c931445103f114656502aa6ce1e81c83e614bb2a6510f52544abf"},
	{Address: "0x3BF69655b3cE5212A3d56f0D78064Cb6F124a60B", Public: "0x3BF69655b3cE5212A3d56f0D78064Cb6F124a60B", BLSKey: "447d5a94aded9675fdf4729f3debe204a1c7cb2b7553e1dde17736e8c6d2a3ce"},

	// 20 - 29
	{Address: "0x3D88FF444D18F7bcC530F5f5171048e725AEc79C", Public: "0x3D88FF444D18F7bcC530F5f5171048e725AEc79C", BLSKey: "6f13f8f7461d670a263f91c8603271718dbb36d99a4a300b75315e4f6e25b489"},
	{Address: "0x40d6f48c7b27BA7544b04456445Cf19B680F5484", Public: "0x40d6f48c7b27BA7544b04456445Cf19B680F5484", BLSKey: "385dc8a1eeb01d255b207dea926c03491f51f6928474e69c03f0ab0a28ae87c9"},
	{Address: "0x43bcBa1c3c3Bf76790d04cad7357229ECD71BDAD", Public: "0x43bcBa1c3c3Bf76790d04cad7357229ECD71BDAD", BLSKey: "37efb861596a44781ec5d0754e113fd68ec6d02129087633b87209309533f1af"},
	{Address: "0x52D77E90caE790ad2bA9DE138Ea8B65cCC5EF652", Public: "0x52D77E90caE790ad2bA9DE138Ea8B65cCC5EF652", BLSKey: "3f9b75217a80be57b852ddd8018ea2fcb31d6199e3b0db4064f631e7dc2d2807"},
	{Address: "0x583B5d4a45E2ce2E29F2Dc6c0645344Bad901755", Public: "0x583B5d4a45E2ce2E29F2Dc6c0645344Bad901755", BLSKey: "5a65f60194e3fcec97a0bffc7104ae8b2319b6c0feea73a25b952ed8f844c7a1"},
	{Address: "0x593815C55fC25B4BcC84473038A121a22796aAA8", Public: "0x593815C55fC25B4BcC84473038A121a22796aAA8", BLSKey: "3e42c6c419fdfa897ba715b6c2149d66e944d0afe311c05f70ae67691820cc29"},
	{Address: "0x59ebA70c8D8B3d4157432815c2A2DA774bA63aa8", Public: "0x59ebA70c8D8B3d4157432815c2A2DA774bA63aa8", BLSKey: "370560f0549e7cbab83a9693b803eadc1a9c03b3445a9277aafd5a895a4e7a64"},
	{Address: "0x5E49BB8be4e199e8ddDe3A09E67D3c23239AC16c", Public: "0x5E49BB8be4e199e8ddDe3A09E67D3c23239AC16c", BLSKey: "49f2f5793502eb401f01b18b6d4440c1e8dfc8ba978a70baccc0a712d4d4a30"},
	{Address: "0x5dc4D61A44EBEb41549021342a290bd726623A38", Public: "0x5dc4D61A44EBEb41549021342a290bd726623A38", BLSKey: "4df4f1031a076da8f4dc7eed2bc7c338170881a6837c2b08826f9745b973d172"},
	{Address: "0x609e8747cdcE518fB86C5990dCE325649C487133", Public: "0x609e8747cdcE518fB86C5990dCE325649C487133", BLSKey: "1df06efea571b52a420bb9b91c492e827b869be5a9039d2f49251daf7fb908aa"},

	// 30 - 39
	{Address: "0x638Ff0c3c291eA08c2653Bb993E3360D63038678", Public: "0x638Ff0c3c291eA08c2653Bb993E3360D63038678", BLSKey: "3afd1c6543c614cce560f40d2c1e97cd08f67f92c4b39eb658b9346baf0309e6"},
	{Address: "0x65c123f9492De546B6e5852f58eB0Df39307Bf45", Public: "0x65c123f9492De546B6e5852f58eB0Df39307Bf45", BLSKey: "300182936b8b4fbb462b638f7abebaaac97f5f78f64e8ef36ea842dc433a45ad"},
	{Address: "0x689a35324d6B8DDDfa3bF5E7b26A23E704dD0100", Public: "0x689a35324d6B8DDDfa3bF5E7b26A23E704dD0100", BLSKey: "17847ab540a1b364064190fead44024913d50e3a7ea06ec370c4b50689612384"},
	{Address: "0x6A6A5FBfA9923EBB76f9E42013e7C4f3CfDC145C", Public: "0x6A6A5FBfA9923EBB76f9E42013e7C4f3CfDC145C", BLSKey: "59c9561d07d68e392a792bf7e8746fb89644addc2d683210ce28a1960ed5df65"},
	{Address: "0x6b9E03aB56786f4F228eE11D965A1a81ed7dA1D4", Public: "0x6b9E03aB56786f4F228eE11D965A1a81ed7dA1D4", BLSKey: "7308cfc0eedf5393f64c9104b4dd841524dba8c6270a9c07d126e21105eec60"},
	{Address: "0x6c11b83856804D1eae8823beB697d09569fE87A0", Public: "0x6c11b83856804D1eae8823beB697d09569fE87A0", BLSKey: "476e9ba5930dac70afb6435a8e50ed37b3a7bf9707d335fd9dd05ebcc73d9251"},
	{Address: "0x72B6aefe8aC9B8873Ab854e6f4fD4801A3F4B2f0", Public: "0x72B6aefe8aC9B8873Ab854e6f4fD4801A3F4B2f0", BLSKey: "3c09d4e7bcfb4a3629b6172374076001ecbdf208106d0e9c71c46963b81e0d00"},
	{Address: "0x76f8d12F6624f713B2D8894A749ad926F7812350", Public: "0x76f8d12F6624f713B2D8894A749ad926F7812350", BLSKey: "2cfdd3c95ffe3b7b70f0b87490ade1c5bd13799e011feb18cf7bdd958fa10453"},
	{Address: "0x78A8D29D81dD02c13a2a6077d887CF661B67E2c0", Public: "0x78A8D29D81dD02c13a2a6077d887CF661B67E2c0", BLSKey: "1829f2c1f13b86744d801ed742393f11915ac555a1db0d9dde1819655590fbf8"},
	{Address: "0x79f8E1B732bA63987873d5eB86C81364C2cF5021", Public: "0x79f8E1B732bA63987873d5eB86C81364C2cF5021", BLSKey: "5ea77c130f52e4fd7437f8e25b7221ba9694a2e4628654b3677e75eee4ebdc4e"},

	// 40 - 49
	{Address: "0x7A4306d4D0A4f15A5fA54486cE4e6403E313805A", Public: "0x7A4306d4D0A4f15A5fA54486cE4e6403E313805A", BLSKey: "53af2abd846ac55ed952cf21b341bfc5d3a1e00ac12129dd9e41972071af234c"},
	{Address: "0x7ACDCB2BAcA2911BdcE98e308515A289ac60b7d2", Public: "0x7ACDCB2BAcA2911BdcE98e308515A289ac60b7d2", BLSKey: "1acaf7104a48a9da4a6d93bbd050230dd803a260e5b1e9bda782e5019885f5b9"},
	{Address: "0x7f42f7a4d66f0387AE77A219d0742E8a706231CA", Public: "0x7f42f7a4d66f0387AE77A219d0742E8a706231CA", BLSKey: "31020fbc52611b249562f8ec723889da9e4aeef1e3bb37e5e90e39c7c474838b"},
	{Address: "0x802bEcc3615Fb8b751ADebA452A30C57F351e8D1", Public: "0x802bEcc3615Fb8b751ADebA452A30C57F351e8D1", BLSKey: "4db9db91dda8c9967f74dfa4b009c76158dcf8006031cb0098d4f77baea8f2e2"},
	{Address: "0x82301962Afa7328FDC34e3610B48D899F031e15F", Public: "0x82301962Afa7328FDC34e3610B48D899F031e15F", BLSKey: "6dc3799065d016ab96d7dd362b82793616399fc3e9615b2d86713905834348a3"},
	{Address: "0x86b4b2dEEE393eBb9633e6F0FEd74F39638A7B4e", Public: "0x86b4b2dEEE393eBb9633e6F0FEd74F39638A7B4e", BLSKey: "5238e72345c82d9e529cdb79a16b0cde505fb7b10e65f9601d2a18dba0396752"},
	{Address: "0x87a157db95dc3517Eb578d4cedee92a5ab275BD5", Public: "0x87a157db95dc3517Eb578d4cedee92a5ab275BD5", BLSKey: "37a61d9c128add82291c7f304f1275ad4abe91413b40f569096b2a2ee18e658d"},
	{Address: "0x880D5c6aD4117D26126543Af48f2f9bCDd4DaA0A", Public: "0x880D5c6aD4117D26126543Af48f2f9bCDd4DaA0A", BLSKey: "3154dfa2bfa5cb6b711424c0a1098aad2b16b75c9a31206e578d87c2b53aa94c"},
	{Address: "0x8cF87bB4BE77d8Dbf16fF61273e02E046a18D716", Public: "0x8cF87bB4BE77d8Dbf16fF61273e02E046a18D716", BLSKey: "6fbebacc62ceb8417a6457e53157dd2d2165d950970db840c62a95e01cc4543c"},
	{Address: "0x8dc63cCA875eAd38d9554bB97171a4f18AbE92E7", Public: "0x8dc63cCA875eAd38d9554bB97171a4f18AbE92E7", BLSKey: "6d67e5640c0dce8c354eeace0a66df912f9e6011d5528b79fc26aa19f74ee8a"},

	// 50 - 59
	{Address: "0x93570Dcb1Bf1a0bD1d476a542309754a6dbCE632", Public: "0x93570Dcb1Bf1a0bD1d476a542309754a6dbCE632", BLSKey: "39f372c4ae8b66aee188d6b93faf5ffc027637c2a7352c850a386286cb1d0322"},
	{Address: "0x9668c58b282f954EA8B732e0D72045bdF19df8B3", Public: "0x9668c58b282f954EA8B732e0D72045bdF19df8B3", BLSKey: "5e026a694cf2610b4a89c1f4b163f300f41f7e9d243f622da36a476e588f07e0"},
	{Address: "0x97b834277538e4517f43f9E11fa0BbebaD7c0d3e", Public: "0x97b834277538e4517f43f9E11fa0BbebaD7c0d3e", BLSKey: "5ad3e33836bb2626279b0f75a70d53becf11362aee66195fca3b79f897229aab"},
	{Address: "0x9c23fE8cdcA1a8E529deeE8eD8492575cc3F9129", Public: "0x9c23fE8cdcA1a8E529deeE8eD8492575cc3F9129", BLSKey: "170a9940960855169235886270a2be5c1033dd5555a8e09b88b20297932ee495"},
	{Address: "0xA28e6f8D23cc3Fe77D531c7D60bd73F8fD71C5c7", Public: "0xA28e6f8D23cc3Fe77D531c7D60bd73F8fD71C5c7", BLSKey: "4ca252941e795dcae69fe7b08744d30e80f1fcd16683934831a1c2427af2a538"},
	{Address: "0xA41F4dDd1b11A6107f1973037D070869495e71E4", Public: "0xA41F4dDd1b11A6107f1973037D070869495e71E4", BLSKey: "1deb567947914013b3ec5b75b098953866af2e09fc0c5d497668d5703ec1f52b"},
	{Address: "0xA721ad85fFfAb28115e1b4B8347A5B42AEA26aA1", Public: "0xA721ad85fFfAb28115e1b4B8347A5B42AEA26aA1", BLSKey: "1b6c1a66406ed091243e5af88c8ca3cfb2d37cccfa2064dec4f615d18f14b2d3"},
	{Address: "0xB18e698BE8698346f7929F4E019D8B1aFE3D04b7", Public: "0xB18e698BE8698346f7929F4E019D8B1aFE3D04b7", BLSKey: "8bfb75d011e179bc7a66915e81849cce26854ffc0ff4ae587e1013cf5be4632"},
	{Address: "0xB1Fa8F1CEa1A78d8887609CebEA592313dD659C1", Public: "0xB1Fa8F1CEa1A78d8887609CebEA592313dD659C1", BLSKey: "6d06b60efbb94aa7a00d55122c1930d29d73d18ba9d5875c0eb5fc8cbaa14959"},
	{Address: "0xB4018FF5B888e902bD952D6e55A5cDbd8C73Ac1A", Public: "0xB4018FF5B888e902bD952D6e55A5cDbd8C73Ac1A", BLSKey: "5b566bafbcbd319f2be986c48122de1bf1ff59af0c67899673cf65c2ef17c2dc"},

	// 60 - 69
	{Address: "0xB68751A436f287CE3DA347277259af5c7bA84e38", Public: "0xB68751A436f287CE3DA347277259af5c7bA84e38", BLSKey: "373f20e831998d9040dec116ed7f5d16b78367ffdf206b13c3c437bdf2869252"},
	{Address: "0xB88AB7A6678c87aeBE7b753459258012eb2Cc76c", Public: "0xB88AB7A6678c87aeBE7b753459258012eb2Cc76c", BLSKey: "5a5b29539e0a635b3c5f3a9064d8f375c7badf5091eace3c91fa2420f00bd4be"},
	{Address: "0xB99Ad8B391eDD1F15c51f773F4bc23Bba7dF45F3", Public: "0xB99Ad8B391eDD1F15c51f773F4bc23Bba7dF45F3", BLSKey: "706d9706f42f7c1229cdf499af4155ebbe9f1636f7df2d91e6640f963069a8f6"},
	{Address: "0xC3FBdE6a171aCc0466614D09b58E013058e7c0d2", Public: "0xC3FBdE6a171aCc0466614D09b58E013058e7c0d2", BLSKey: "670cc2a2cf405c4fa7a560ae2edd2e193b35c73583a494363828dfa18fd8ab6c"},
	{Address: "0xC6b6a71d6f0C5b98E25FCf14b5378c807B0d475a", Public: "0xC6b6a71d6f0C5b98E25FCf14b5378c807B0d475a", BLSKey: "6ac9db11d32288d155cdff1a29b07ccb52a6be140696bdbe0767ef32fae4d630"},
	{Address: "0xC6fDB78B643E8eBaC472dB61c1e84B6Fe0Abe185", Public: "0xC6fDB78B643E8eBaC472dB61c1e84B6Fe0Abe185", BLSKey: "67cd0a36e8d3f1d61a13825140dce51e896d9eb307178074a727f30ac587fee3"},
	{Address: "0xD0F9AD2b60792fAff02f8Bd0F2D9cE2790722706", Public: "0xD0F9AD2b60792fAff02f8Bd0F2D9cE2790722706", BLSKey: "5a08371e86970c2d7c15e9a0d404ab3e4503dc79fdaa5c3ab73b9f37c231b55"},
	{Address: "0xD28B4bC96020De252A0ee817767B6Cdb26A47d73", Public: "0xD28B4bC96020De252A0ee817767B6Cdb26A47d73", BLSKey: "42d76f1f5c6878524f2bb61d58ab0ad3d48fda2ef53792da03208933892f2c98"},
	{Address: "0xD31095BE15D4b0b16657EEB72e0cc81e24EAc101", Public: "0xD31095BE15D4b0b16657EEB72e0cc81e24EAc101", BLSKey: "2edfef59b86f638653c61e5da7d6a88fb481fd1ea5c937509690af2c46329e4a"},
	{Address: "0xD499fAC5afa17b5705B91838753Bfbf2e20138e4", Public: "0xD499fAC5afa17b5705B91838753Bfbf2e20138e4", BLSKey: "485bc6c4979a0e586a77aa93a3ad677a9908ff042d958cdfb275a586fbfc1061"},

	// 70 - 79
	{Address: "0xDfC0B00B629dDD5704a156A0D932F78692fc842F", Public: "0xDfC0B00B629dDD5704a156A0D932F78692fc842F", BLSKey: "417da445b2653caf6b094fd95182548bf97a0c386a4f8ad24b93a8ea96b44073"},
	{Address: "0xE2ab78ecf325084485957B2599d53Bcf944Cbca8", Public: "0xE2ab78ecf325084485957B2599d53Bcf944Cbca8", BLSKey: "45c2600a71d2e538134777477940cddded763f4734d76678e32ec87d31126943"},
	{Address: "0xEC7C495866689d6b7E335D810645F440f16F86d0", Public: "0xEC7C495866689d6b7E335D810645F440f16F86d0", BLSKey: "40797d43fed357ccf2d643317c70cb5177aa289f3ee81bcf6c2e23933f65fd8d"},
	{Address: "0xEd677E021df3542998e407970E1127d334Be0285", Public: "0xEd677E021df3542998e407970E1127d334Be0285", BLSKey: "221fef10b540078b9dc1c1254626b5c4a2eecc29d756b8afaf13d6937629a8f7"},
	{Address: "0xF262609617c202087B31aCf364C00967f5Cd85De", Public: "0xF262609617c202087B31aCf364C00967f5Cd85De", BLSKey: "27df38aaff7fcdeb673eeae0282362704506bc54ec6856ccd888cc47af54938d"},
	{Address: "0xF7E33ef7132bcc716C2242385d9862c3c43baB7E", Public: "0xF7E33ef7132bcc716C2242385d9862c3c43baB7E", BLSKey: "6e2479d510f0f6058d2b277c4a48fd6173bebf4bf34619ac473738d4da31f3c7"},
	{Address: "0xa3B34f4E21C6c44A603E3c53abbF8b10C7BdaF59", Public: "0xa3B34f4E21C6c44A603E3c53abbF8b10C7BdaF59", BLSKey: "433ce09e054052678ac5f1c5021d07f181496116b021ec13405b98988fa9f2d0"},
	{Address: "0xa4BF67e67910225aA1C3Cd65595d8a1b1227F42E", Public: "0xa4BF67e67910225aA1C3Cd65595d8a1b1227F42E", BLSKey: "13d2218fa9236f4c6941e0c71e6cc8d1911f6c6289e11db0e2125820e370672"},
	{Address: "0xa61CA9f1EB26787EEd89dAEE4A326C4e1cb5eCdB", Public: "0xa61CA9f1EB26787EEd89dAEE4A326C4e1cb5eCdB", BLSKey: "595dd876e4e2c1f4e0ccfe2fd02440bdbe03379333b98ef94e906ba66b1d0c"},
	{Address: "0xa714cd269A0ca23131C8cD5aeFC49F450578C4B4", Public: "0xa714cd269A0ca23131C8cD5aeFC49F450578C4B4", BLSKey: "387811013e8579d51b4625368dbefd4c2ae7205de6d6593a72b06c44c51e8629"},

	// 80 - 89
	{Address: "0xb108BF4945Bd7975cF974f47476e689ACd542F23", Public: "0xb108BF4945Bd7975cF974f47476e689ACd542F23", BLSKey: "2e48ddb0cb3d6c816e8faff0661fc083b1a70c345a3130049af64ab25f12c822"},
	{Address: "0xb69509BFf7Ac53eA998e16FBC247f24F88eE8572", Public: "0xb69509BFf7Ac53eA998e16FBC247f24F88eE8572", BLSKey: "6a6336b9e939e44caaac541829ebf8443b60bbb39359f06becb7cd2a432b9597"},
	{Address: "0xc55c56F661eD185103839FdFeFd80DC38938913b", Public: "0xc55c56F661eD185103839FdFeFd80DC38938913b", BLSKey: "6ed95c176e3805957c1781466ef59fbb45ab134294e5d6e84c51b53393db4c3e"},
	{Address: "0xd021c9a6A8816FE57a3A4CBd02fA824e0e92421B", Public: "0xd021c9a6A8816FE57a3A4CBd02fA824e0e92421B", BLSKey: "27c5867d1609b897dbd6858d82d21b7b875927a1a9dcc27f0dd53ac2d8223eee"},
	{Address: "0xdA1DF648bC047546326D05dF370ec0ee3D84642A", Public: "0xdA1DF648bC047546326D05dF370ec0ee3D84642A", BLSKey: "2742de48065f12da2605606972f7d724d645cd9154c843697699c488f3e5b532"},
	{Address: "0xeaD1fAa7E5Fdb6136057d4BfCa1f05D220D1441f", Public: "0xeaD1fAa7E5Fdb6136057d4BfCa1f05D220D1441f", BLSKey: "3d72a88df0677822b9a607b9fc55d2d8d3f088363bf380a3c334c65a24c49411"},
	{Address: "0xf10f63f5Bd46c58d2e9530E7F8cb6b4336D05d4E", Public: "0xf10f63f5Bd46c58d2e9530E7F8cb6b4336D05d4E", BLSKey: "613f101fd405daa52c50416517ee151e4f6161fed5f7571b64d2470f86145aed"},
	{Address: "0xfB81EFd254Fe117047872146806153539F89669E", Public: "0xfB81EFd254Fe117047872146806153539F89669E", BLSKey: "53b1951f2c9976fa5e92e936da3dd7a6a356fb387bb487489703347f708f9d5f"},
	{Address: "0xfdc963E875Ea99E434e4B815b7d8Bf506dAA9222", Public: "0xfdc963E875Ea99E434e4B815b7d8Bf506dAA9222", BLSKey: "3606585dd5e71c3117da6f08d9c0fb55dac5caf67b22f5f0bdd1605e2e0132a2"},
}
