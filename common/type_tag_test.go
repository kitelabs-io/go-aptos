package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsTypeTagStruct(t *testing.T) {
	t.Run("it should match correct tag", func(t *testing.T) {
		tag := `0x48271d39d0b05bd6efca2278f22277d6fcc375504f9839fd73f74ace240861af::stable_pool::StablePool<0xdf3d5eb83df80dfde8ceb1edaa24d8dbc46da6a89ae134a858338e1b86a29e38::coin::Returd, 0x1::aptos_coin::AptosCoin, 0x48271d39d0b05bd6efca2278f22277d6fcc375504f9839fd73f74ace240861af::base_pool::Null, 0x48271d39d0b05bd6efca2278f22277d6fcc375504f9839fd73f74ace240861af::base_pool::Null>`

		a, ok := AsTypeTagStruct(tag)
		assert.True(t, ok)

		assert.Equal(t, "0x48271d39d0b05bd6efca2278f22277d6fcc375504f9839fd73f74ace240861af", a.ModuleAddress.ToHex())
		assert.Equal(t, "stable_pool", a.ModuleName)
		assert.Equal(t, "StablePool", a.StructName)
		assert.Equal(t, 4, len(a.TypeParams))
		assert.Equal(t, "0xdf3d5eb83df80dfde8ceb1edaa24d8dbc46da6a89ae134a858338e1b86a29e38::coin::Returd", a.TypeParams[0].String())
		assert.Equal(t, "0x1::aptos_coin::AptosCoin", a.TypeParams[1].String())
		assert.Equal(t, "0x48271d39d0b05bd6efca2278f22277d6fcc375504f9839fd73f74ace240861af::base_pool::Null", a.TypeParams[2].String())
		assert.Equal(t, "0x48271d39d0b05bd6efca2278f22277d6fcc375504f9839fd73f74ace240861af::base_pool::Null", a.TypeParams[3].String())
	})

	t.Run("it should not match correct tag", func(t *testing.T) {
		tag := `0xe52923154e25c258d9befb0237a30b4001c63dc3bb73011c29cb3739befffcef::swap_v2dot1::TokenPairMetadata<0x1::aptos_coin::AptosCoin, 0x6ee5ff12d9af89de4cb9f127bc4c484d26acda56c03536b5e3792eac94da0a36::swap_v2::LPToken<0x1::aptos_coin::AptosCoin, 0x83b619e2d9e6e10d15ed4b714111a4cd9526c1c2ae0eec4b252a619d3e8bdda3::MAU::MAU>>`

		a, ok := AsTypeTagStruct(tag)
		assert.True(t, ok)

		assert.Equal(t, "0xe52923154e25c258d9befb0237a30b4001c63dc3bb73011c29cb3739befffcef", a.ModuleAddress.ToHex())
		assert.Equal(t, "swap_v2dot1", a.ModuleName)
		assert.Equal(t, "TokenPairMetadata", a.StructName)
		assert.Equal(t, 2, len(a.TypeParams))
		assert.Equal(t, "0x1::aptos_coin::AptosCoin", a.TypeParams[0].String())
		assert.Equal(t, "0x6ee5ff12d9af89de4cb9f127bc4c484d26acda56c03536b5e3792eac94da0a36::swap_v2::LPToken<0x1::aptos_coin::AptosCoin, 0x83b619e2d9e6e10d15ed4b714111a4cd9526c1c2ae0eec4b252a619d3e8bdda3::MAU::MAU>", a.TypeParams[1].String())
	})
}

func TestTypeTagStruct_String(t *testing.T) {
	t.Run("it should match correct tag", func(t *testing.T) {
		tag := "0xc7efb4076dbe143cbcd98cfaaa929ecfc8f299203dfff63b95ccb6bfe19850fa::swap::LPToken<0x159df6b7689437016108a019fd5bef736bac692b6d4a1f10c941f6fbb9a74ca6::oft::CakeOFT, 0x1::aptos_coin::AptosCoin>"

		typeTagStruct, ok := AsTypeTagStruct(tag)

		assert.True(t, ok)
		assert.Equal(t, tag, typeTagStruct.String())
	})
}
