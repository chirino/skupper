package crypto_test

import (
	"github.com/skupperproject/skupper/internal/crypto"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCrypto(t *testing.T) {

	privateKey, err := crypto.NewKey()
	require.NoError(t, err)

	{
		privateTxt := privateKey.String()
		actual, err := crypto.ParseKey(privateTxt)
		require.NoError(t, err)
		require.Equal(t, privateKey, actual)
	}

	publicKey, err := privateKey.PublicKey()
	require.NoError(t, err)

	{
		publicTxt := publicKey.String()
		actual, err := crypto.ParseKey(publicTxt)
		require.NoError(t, err)
		require.Equal(t, publicKey, actual)
	}

	message := []byte("hello world")
	originalSealed, err := crypto.SealV1(publicKey[:], message)
	require.NoError(t, err)

	data := originalSealed.String()
	parsedSealed, err := crypto.ParseSealed(data)
	require.NoError(t, err)

	require.Equal(t, originalSealed, parsedSealed)

	actual, err := parsedSealed.Open(privateKey[:])
	require.NoError(t, err)

	require.Equal(t, message, actual)
}
