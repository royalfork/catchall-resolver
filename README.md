# 🅴🅽🆂 Catchall Resolver

CatchallResolver is an [ENSIP-10](https://docs.ens.domains/ens-improvement-proposals/ensip-10-wildcard-resolution) "Wildcard" resolver which proxies all subdomain resolution requests to the parent resolver.  For instance, if `royalfork.eth` has set a "text/email" record to `ens@royalfork.org`, with the CatchallResolver, `sub.royalfork.eth`, `another.royalfork.eth`, and `double.sub.royalfork.eth` will also have text/email records set to `ens@royalfork.org`.  The CatchallResolver works for all ENSIP defined resolver methods (addr, text, contenthash, etc).

CatchallResolver is deployed on mainnet at: [0xca7c4a117baef1b7a122d55ede7216ba6631ac3e](https://etherscan.io/address/0xca7c4a117baef1b7a122d55ede7216ba6631ac3e)

## Use

Like the ENS public resolver, a single deployment of CatchallResolver can be used by many ENS domains.  If you'd like to use CatchallResolver for your domain (such that *.yourdomain.eth resolves to yourdomain.eth):
1. Call the CatchallResolver's `setResolver(node, resolver)` method with:
   - node = Namehash of your domain (eg: `namehash(yourdomain.eth)`)
   - resolver = Address to the "catchall" resolver to which resolution requests are sent (eg: the PublicResolver contract at `0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41`)
2. Set your domain's ENS resolver to "Catchall Resolver"
   - Call the [ENS Registry's](https://etherscan.io/address/0x00000000000c2e074ec69a0dfb2997ba6c7d2e1e) `setResolver(node, resolver)` with resolver address: `0xca7c4a117baef1b7a122d55ede7216ba6631ac3e`

#### Example

The Namehash of `royalfork.eth` is `f47a153bb881860e9a4390b84b063154e9623c32a1611c73aef2038a134d8eba`, and resolver is `0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41`.

0. On [0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41](https://etherscan.io/address/0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41#readContract), validate that resolver does indeed service your node.
    <img width="732" alt="Screen Shot 2022-05-01 at 7 47 00 PM" src="https://user-images.githubusercontent.com/8282941/166172082-baad535e-f407-419a-94c3-a9b9d04352b5.png">

1. On [0xca7c4a117baef1b7a122d55ede7216ba6631ac3e](https://etherscan.io/address/0xca7c4a117baef1b7a122d55ede7216ba6631ac3e#writeContract), set the parent resolver for domain:
    <img width="713" alt="Screen Shot 2022-05-01 at 8 50 10 PM" src="https://user-images.githubusercontent.com/8282941/166175354-0bb70987-9397-463e-8caf-6b432d7e4124.png">
    
2. On [0xca7c4a117baef1b7a122d55ede7216ba6631ac3e](https://etherscan.io/address/0xca7c4a117baef1b7a122d55ede7216ba6631ac3e#readContract), ensure that resolution is working:
    <img width="735" alt="Screen Shot 2022-05-01 at 8 53 33 PM" src="https://user-images.githubusercontent.com/8282941/166175924-9f772a7b-0b18-493b-ade1-96e30f99f325.png">

3. On [0x00000000000c2e074ec69a0dfb2997ba6c7d2e1e](https://etherscan.io/address/0x00000000000c2e074ec69a0dfb2997ba6c7d2e1e), set resolver to CatchallResolver:
    <img width="735" alt="Screen Shot 2022-05-01 at 8 53 33 PM" src="https://user-images.githubusercontent.com/8282941/166175527-8af9dbd9-fd83-44e4-a01e-75ead4a1af3c.png">

4. Using [ethers.js](https://docs.ethers.io/v5/), test that everything works as expected:
    ```
    const { ethers } = require("ethers");
    const provider = new ethers.providers.JsonRpcProvider("https://mainnet.infura.io/v3/3761559e11414fd8afa394df10d77d00");
    
    (async () => {
      for (const name of [
        'royalfork.eth',
        'sub.royalfork.eth',
        'a.b.c.royalfork.eth'
      ]){
        console.log(`${name}`);
        let resolver = await provider.getResolver(name);
        let resolveName = await provider.resolveName(name);
        if (resolver) {
          console.log(`resolver address ${resolver.address}`);
          console.log(`eth address ${resolveName}`);
        } else {
          console.log(`resolver not found for ${name}`);
        }
        console.log(`-----`);
      }
    })();
    ```
    Ouputs:
    ```
    royalfork.eth
    resolver address 0xca7C4a117bAef1B7A122D55Ede7216Ba6631AC3E
    eth address 0xb82875007A206D52222887B8Bc21ed309357f878
    -----
    sub.royalfork.eth
    resolver address 0xca7C4a117bAef1B7A122D55Ede7216Ba6631AC3E
    eth address 0xb82875007A206D52222887B8Bc21ed309357f878
    -----
    a.b.c.royalfork.eth
    resolver address 0xca7C4a117bAef1B7A122D55Ede7216Ba6631AC3E
    eth address 0xb82875007A206D52222887B8Bc21ed309357f878
    -----
    ```

😀
