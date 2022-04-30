package test

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/royalfork/catchall-resolver/bindings"
	"github.com/royalfork/ens/enstest"
	"github.com/royalfork/ens/util"
	"github.com/royalfork/soltest"
)

func TestDirectResolver(t *testing.T) {
	chain, accts := soltest.New()

	ensTest, err := enstest.New(accts[0], chain)
	if err != nil {
		t.Fatal(err)
	}

	node, err := ensTest.RegisterETHDomain(accts[1].Addr, "royalfork")
	if err != nil {
		t.Fatal(err)
	}

	_, _, catchall, err := bindings.DeployCatchallResolver(accts[0].Auth, chain, ensTest.RegistryAddr)
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	if !chain.Succeed(catchall.SetResolver(accts[1].Auth, node, ensTest.ResolverAddr)) {
		t.Fatal("catchall can't set resolver")
	}

	t.Run("supportsInterface", func(t *testing.T) {
		for _, iID := range [][4]byte{
			{0x01, 0xff, 0xc9, 0xa7}, // supportsInterface
			{0x3b, 0x3b, 0x57, 0xde}, // addr
			{0x59, 0xd1, 0xd4, 0x3c}, // text
		} {
			if ok, err := catchall.SupportsInterface(&bind.CallOpts{}, iID); err != nil {
				t.Fatal(err)
			} else if !ok {
				t.Errorf("resolver should support interface: %x", iID)
			}
		}
	})

	t.Run("addr", func(t *testing.T) {
		if addr, err := catchall.Addr(&bind.CallOpts{}, node); err != nil {
			t.Fatal(err)
		} else if want := (common.Address{}); addr != want {
			t.Errorf("want addr: %s, got: %s", want, addr)
		}

		if !chain.Succeed(ensTest.Resolver.SetAddr0(accts[1].Auth, node, accts[1].Addr)) {
			t.Fatal("unable to set addr")
		}

		if addr, err := catchall.Addr(&bind.CallOpts{}, node); err != nil {
			t.Fatal(err)
		} else if addr != accts[1].Addr {
			t.Errorf("want addr: %s, got: %s", accts[1].Addr, addr)
		}
	})

	t.Run("text", func(t *testing.T) {
		if val, err := catchall.Text(&bind.CallOpts{}, node, "key"); err != nil {
			t.Fatal(err)
		} else if val != "" {
			t.Errorf("want text: %s, got: %s", "", val)
		}

		if !chain.Succeed(ensTest.Resolver.SetText(accts[1].Auth, node, "key", "val")) {
			t.Fatal("unable to set text")
		}

		if val, err := catchall.Text(&bind.CallOpts{}, node, "bad_key"); err != nil {
			t.Fatal(err)
		} else if val != "" {
			t.Errorf("want text: %s, got: %s", "", val)
		}

		if val, err := catchall.Text(&bind.CallOpts{}, node, "key"); err != nil {
			t.Fatal(err)
		} else if want := "val"; val != want {
			t.Errorf("want text: %s, got: %s", want, val)
		}
	})
}

func TestSetResolver(t *testing.T) {
	// TODO
}

func TestResolve(t *testing.T) {
	chain, accts := soltest.New()

	ensTest, err := enstest.New(accts[0], chain)
	if err != nil {
		t.Fatal(err)
	}

	_, _, catchall, err := bindings.DeployCatchallResolver(accts[0].Auth, chain, ensTest.RegistryAddr)
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	node, err := ensTest.RegisterETHDomain(accts[1].Addr, "royalfork")
	if err != nil {
		t.Fatal(err)
	}

	if !chain.Succeed(catchall.SetResolver(accts[1].Auth, node, ensTest.ResolverAddr)) {
		t.Fatal("catchall can't set resolver")
	}

	domain := "sub.royalfork.eth"
	encDomain, err := util.DNSEncode(domain)
	if err != nil {
		t.Fatal(err)
	}

	abi, err := bindings.CatchallResolverMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("shortData", func(t *testing.T) {
		if _, err := catchall.Resolve(&bind.CallOpts{}, []byte{}, []byte{0xDE, 0xAD, 0xBE, 0xEF}); err == nil {
			t.Fatal("expected non-nil error")
		}
	})

	t.Run("badData", func(t *testing.T) {
		badData := make([]byte, 132)
		if _, err := catchall.Resolve(&bind.CallOpts{}, []byte{}, badData); err == nil {
			t.Fatal("expected non-nil error")
		}
	})

	t.Run("addr", func(t *testing.T) {
		if !chain.Succeed(ensTest.Resolver.SetAddr0(accts[1].Auth, node, accts[1].Addr)) {
			t.Fatal("unable to set addr")
		}

		data, err := abi.Pack("addr", [32]byte{})
		if err != nil {
			t.Fatal(err)
		}

		result, err := catchall.Resolve(&bind.CallOpts{}, encDomain, data)
		if err != nil {
			t.Fatal(err)
		}

		if output, err := abi.Unpack("addr", result); err != nil {
			t.Fatal(err)
		} else if addr, ok := output[0].(common.Address); !ok {
			t.Fatal("resolve output not addr")
		} else if addr != accts[1].Addr {
			t.Errorf("want addr: %s, got: %s", accts[1].Addr, addr)
		}
	})

	t.Run("text", func(t *testing.T) {
		if !chain.Succeed(ensTest.Resolver.SetText(accts[1].Auth, node, "key", "val")) {
			t.Fatal("unable to set addr")
		}

		data, err := abi.Pack("text", [32]byte{}, "key")
		if err != nil {
			t.Fatal(err)
		}

		result, err := catchall.Resolve(&bind.CallOpts{}, encDomain, data)
		if err != nil {
			t.Fatal(err)
		}

		if output, err := abi.Unpack("text", result); err != nil {
			t.Fatal(err)
		} else if text, ok := output[0].(string); !ok {
			t.Fatal("resolve output not string")
		} else if want := "val"; text != "val" {
			t.Errorf("want text: %s, got: %s", want, text)
		}
	})
}

func TestResolver(t *testing.T) {
	chain, accts := soltest.New()

	ensTest, err := enstest.New(accts[0], chain)
	if err != nil {
		t.Fatal(err)
	}

	_, _, catchall, err := bindings.DeployCatchallResolver(accts[0].Auth, chain, ensTest.RegistryAddr)
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	t.Run("noResolver", func(t *testing.T) {
		for _, domain := range []string{
			"eth",
			"noexist.eth",
			"sub.noexist.eth",
		} {
			encDomain, err := util.DNSEncode(domain)
			if err != nil {
				t.Fatal(err)
			}

			resolverAddr, _, err := catchall.Resolver(&bind.CallOpts{}, encDomain)
			if err != nil {
				t.Fatal(err)
			}

			if resolverAddr != (common.Address{}) {
				t.Errorf("want resolver: %s, got: %s", common.Address{}, resolverAddr)
			}
		}
	})

	t.Run("subResolver", func(t *testing.T) {
		// Set resolver for example.eth
		node, err := ensTest.RegisterETHDomain(accts[1].Addr, "example")
		if err != nil {
			t.Fatal(err)
		}
		if !chain.Succeed(catchall.SetResolver(accts[1].Auth, node, common.Address{1})) {
			t.Fatal("unable to set catchall resolver")
		}

		// Set resolver for example.sub.example.eth
		label, err := util.LabelHash("sub")
		if err != nil {
			t.Fatal(err)
		}
		if !chain.Succeed(ensTest.Registry.SetSubnodeOwner(accts[1].Auth, node, label, accts[1].Addr)) {
			t.Fatal("unable to register sub.example.eth")
		}
		node, err = util.NameHash("sub.example.eth")
		if err != nil {
			t.Fatal(err)
		}
		label, err = util.LabelHash("example")
		if err != nil {
			t.Fatal(err)
		}
		if !chain.Succeed(ensTest.Registry.SetSubnodeOwner(accts[1].Auth, node, label, accts[1].Addr)) {
			t.Fatal("unable to register example.sub.example.eth")
		}
		node, err = util.NameHash("example.sub.example.eth")
		if err != nil {
			t.Fatal(err)
		}
		if !chain.Succeed(catchall.SetResolver(accts[1].Auth, node, common.Address{2})) {
			t.Fatal("unable to set catchall resolver")
		}

		for _, test := range []struct {
			domain       string
			resolverAddr common.Address
			resolverName string
		}{
			{"eth", common.Address{}, ""},
			{"example.eth", common.Address{1}, "example.eth"},
			{"sub.example.eth", common.Address{1}, "example.eth"},
			{"a.sub.example.eth", common.Address{1}, "example.eth"},
			{"b.sub.example.eth", common.Address{1}, "example.eth"},
			{"example.sub.example.eth", common.Address{2}, "example.sub.example.eth"},
			{"a.example.sub.example.eth", common.Address{2}, "example.sub.example.eth"},
			{"b.example.sub.example.eth", common.Address{2}, "example.sub.example.eth"},
			{"c.a.example.sub.example.eth", common.Address{2}, "example.sub.example.eth"},
		} {
			encDomain, err := util.DNSEncode(test.domain)
			if err != nil {
				t.Errorf("dnsencode err for domain %s: %v", test.domain, err)
			}

			resolverAddr, resolverName, err := catchall.Resolver(&bind.CallOpts{}, encDomain)
			if err != nil {
				t.Errorf("resolver err for domain %s: %v", test.domain, err)
			}
			if resolverAddr != test.resolverAddr {
				t.Errorf("want %s resolver: %s; got: %s", test.domain, test.resolverAddr, resolverAddr)
			}
			if test.resolverAddr == (common.Address{}) {
				continue
			}

			decDomain, err := util.DNSDecode(resolverName)
			if err != nil {
				t.Errorf("dnsdecode err for domain %s, resolver name: %v: %v", test.domain, resolverName, err)
			}

			if decDomain != test.resolverName {
				t.Errorf("want %s resolver name: %s; got: %s", test.domain, test.resolverName, decDomain)
			}
		}
	})
}
