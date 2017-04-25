package kallax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUULIDIsEmpty(t *testing.T) {
	r := require.New(t)
	var id ULID
	r.True(id.IsEmpty())

	id = NewULID()
	r.False(id.IsEmpty())
}

func TestULID_Value(t *testing.T) {
	id := NewULID()
	v, _ := id.Value()
	require.Equal(t, id.String(), v)
}

func TestUULID_ThreeNewIDsAreDifferent(t *testing.T) {
	r := require.New(t)
	id1 := NewULID()
	id2 := NewULID()
	id3 := NewULID()

	r.NotEqual(id1, id2)
	r.NotEqual(id1, id3)
	r.NotEqual(id2, id3)

	r.True(id1 == id1)
	r.False(id1 == id2)
}

func TestULID_ScanValue(t *testing.T) {
	r := require.New(t)

	expected := NewULID()
	v, err := expected.Value()
	r.NoError(err)

	var id ULID
	r.NoError(id.Scan(v))

	r.Equal(expected, id)
	r.Equal(expected.String(), id.String())

	r.NoError(id.Scan([]byte("015af13d-2271-fb69-2dcd-fb24a1fd7dcc")))
}

func TestVirtualColumn(t *testing.T) {
	r := require.New(t)
	record := newModel("", "", 0)
	record.virtualColumns = nil
	r.Equal(nil, record.VirtualColumn("foo"))

	record.virtualColumns = nil
	s := VirtualColumn("foo", record, new(ULID))

	id := NewULID()
	v, err := id.Value()
	r.NoError(err)
	r.NoError(s.Scan(v))
	r.Len(record.virtualColumns, 1)
	r.Equal(&id, record.VirtualColumn("foo"))

	r.Error(s.Scan(nil))
}
