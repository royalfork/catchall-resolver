package test

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/royalfork/catchall-resolver/bindings"
	"github.com/royalfork/ens/enstest"
	"github.com/royalfork/soltest"
)

func TestSubdomainResolver(t *testing.T) {
	chain, accts := soltest.New()

	ensTest, err := enstest.New(accts[0], chain)
	if err != nil {
		t.Fatal(err)
	}

	node, err := ensTest.RegisterETHDomain(accts[1].Addr, "royalfork")
	if err != nil {
		t.Fatal(err)
	}

	if !chain.Succeed(ensTest.Resolver.SetAddr0(accts[1].Auth, node, accts[1].Addr)) {
		t.Fatal("unable to set addr")
	}

	_, _, resolver, err := bindings.DeployCatchallResolver(accts[0].Auth, chain, ensTest.ResolverAddr, node)
	if err != nil {
		t.Fatal(err)
	}
	chain.Commit()

	abi, err := bindings.CatchallResolverMetaData.GetAbi()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("addr", func(t *testing.T) {
		if addr, err := resolver.Addr(&bind.CallOpts{}, node); err != nil {
			t.Fatal(err)
		} else if addr != accts[1].Addr {
			t.Errorf("want addr: %s, got: %s", accts[1].Addr, addr)
		}
	})

	t.Run("resolve", func(t *testing.T) {
		// TODO test data < 32 bytes
		// TODO test data isn't resolver method
		// TODO resolver has no record for the parentNode

		t.Run("addr", func(t *testing.T) {
			data, err := abi.Pack("addr", [32]byte{})
			if err != nil {
				t.Fatal(err)
			}

			result, err := resolver.Resolve(&bind.CallOpts{}, []byte{}, data)
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
	})
}
