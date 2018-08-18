package logstruct_test

import (
	"testing"

	logstruct "github.com/m-mizutani/logstruct/lib"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateModel(t *testing.T) {
	m := logstruct.NewModel()

	f1, isNew1 := m.InputLog("a b c d e")
	assert.True(t, isNew1)
	require.NotNil(t, f1)

	f2, isNew2 := m.InputLog("a b x d e")
	assert.False(t, isNew2)
	require.NotNil(t, f2)

	assert.Equal(t, f1, f2)
}

func TestExportAndImport(t *testing.T) {
	m := logstruct.NewModel()
	m.InputLog("a b c d e")
	m.InputLog("a x b y z")

	data, err := m.Export()
	require.NoError(t, err)

	m2 := logstruct.NewModel()
	err = m2.Import(data)
	require.NoError(t, err)

	f, n := m2.InputLog("a X b Y z")
	assert.NotNil(t, f)
	assert.False(t, n)
}
