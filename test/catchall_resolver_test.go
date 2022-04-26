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
		} else if want := (common.Address{}); addr != want {
			t.Errorf("want addr: %s, got: %s", want, addr)
		}

		if !chain.Succeed(ensTest.Resolver.SetAddr0(accts[1].Auth, node, accts[1].Addr)) {
			t.Fatal("unable to set addr")
		}

		if addr, err := resolver.Addr(&bind.CallOpts{}, node); err != nil {
			t.Fatal(err)
		} else if addr != accts[1].Addr {
			t.Errorf("want addr: %s, got: %s", accts[1].Addr, addr)
		}
	})

	t.Run("text", func(t *testing.T) {
		if val, err := resolver.Text(&bind.CallOpts{}, node, "key"); err != nil {
			t.Fatal(err)
		} else if val != "" {
			t.Errorf("want text: %s, got: %s", "", val)
		}

		if !chain.Succeed(ensTest.Resolver.SetText(accts[1].Auth, node, "key", "val")) {
			t.Fatal("unable to set addr")
		}

		if val, err := resolver.Text(&bind.CallOpts{}, node, "bad_key"); err != nil {
			t.Fatal(err)
		} else if val != "" {
			t.Errorf("want text: %s, got: %s", "", val)
		}

		if val, err := resolver.Text(&bind.CallOpts{}, node, "key"); err != nil {
			t.Fatal(err)
		} else if want := "val"; val != want {
			t.Errorf("want text: %s, got: %s", want, val)
		}
	})

	t.Run("resolve", func(t *testing.T) {
		t.Run("shortData", func(t *testing.T) {
			if _, err := resolver.Resolve(&bind.CallOpts{}, []byte{}, []byte{0xDE, 0xAD, 0xBE, 0xEF}); err == nil {
				t.Fatal("expected non-nil error")
			}
		})

		t.Run("badData", func(t *testing.T) {
			badData := make([]byte, 132)
			if _, err := resolver.Resolve(&bind.CallOpts{}, []byte{}, badData); err == nil {
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

		t.Run("text", func(t *testing.T) {
			if !chain.Succeed(ensTest.Resolver.SetText(accts[1].Auth, node, "key", "val")) {
				t.Fatal("unable to set addr")
			}

			data, err := abi.Pack("text", [32]byte{}, "key")
			if err != nil {
				t.Fatal(err)
			}

			result, err := resolver.Resolve(&bind.CallOpts{}, []byte{}, data)
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
	})
}
