package crypto_test

import (
	"fmt"
	"github.com/skupperproject/skupper/internal/crypto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCrypto(t *testing.T) {

	privateKey, err := crypto.NewKey()
	require.NoError(t, err)
	publicKey, err := privateKey.PublicKey()
	require.NoError(t, err)

	message := []byte("hello world")
	originalSealed, err := crypto.SealV1(publicKey[:], message)
	require.NoError(t, err)

	data := originalSealed.String()
	fmt.Println(data)
	parsedSealed, err := crypto.ParseSealed(data)
	require.NoError(t, err)

	require.Equal(t, originalSealed, parsedSealed)

	actual, err := parsedSealed.Open(privateKey[:])
	require.NoError(t, err)

	require.Equal(t, message, actual)
}
