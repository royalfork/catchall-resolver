# 🅴🅽🆂 Catchall Resolver

CatchallResolver is an [ENSIP-10](https://docs.ens.domains/ens-improvement-proposals/ensip-10-wildcard-resolution) "Wildcard" resolver which proxies all subdomain resolution requests to the parent resolver.  For instance, if `royalfork.eth` has set a "text/email" record to `ens@royalfork.org`, with the CatchallResolver, `sub.royalfork.eth`, `another.royalfork.eth`, and `double.sub.royalfork.eth` will also have text/email records set to `ens@royalfork.org`.  The CatchallResolver works for all ENSIP defined resolver methods (addr, text, contenthash, etc).

CatchallResolver is deployed on mainnet at: [0xca7c4a117baef1b7a122d55ede7216ba6631ac3e](https://etherscan.io/address/0xca7c4a117baef1b7a122d55ede7216ba6631ac3e)

## Use

Like the ENS public resolver, a single deployment of CatchallResolver can be used by many ENS domains.  If you'd like to use CatchallResolver for your domain (such that *.yourdomain.eth resolves to yourdomain.eth):
1. Call the CatchallResolver's `setResolver(node, resolver)` method with:
   - node = Namehash of your domain.
   - resolver = address to the "catchall" resolver to which resolution requests are sent (for most, this will be the PublicResolver `0x4976fb03C32e5B8cfe2b6cCB31c09Ba78EBaBa41`)
2. Set your domain's ENS resolver to CatchallResolver
   - Call ens.setResolver(node, resolver) with resolver address: `0xca7c4a117baef1b7a122d55ede7216ba6631ac3e`
