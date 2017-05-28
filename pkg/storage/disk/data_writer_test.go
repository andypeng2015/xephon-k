package disk

import (
	"os"
	"testing"

	asst "github.com/stretchr/testify/assert"
	"github.com/xephonhq/xephon-k/pkg/common"
	"github.com/xephonhq/xephon-k/pkg/util"
)

func TestNewLocalFileIndexWriter(t *testing.T) {
	assert := asst.New(t)
	f := util.TempFile(t, "xephon")
	defer os.Remove(f.Name())

	w := NewLocalFileWriter(f, -1)
	assert.NotNil(w.Close())
}

func TestLocalFileWriter_WriteSeries(t *testing.T) {
	assert := asst.New(t)
	f := util.TempFile(t, "xephon")
	//defer os.Remove(f.Name())

	w := NewLocalFileWriter(f, -1)
	s := common.NewIntSeries("s")
	s.Tags = map[string]string{"os": "ubuntu", "machine": "machine-01"}
	s.Points = []common.IntPoint{{T: 1359788400000, V: 1}, {T: 1359788500000, V: 2}}

	assert.Nil(w.WriteSeries(s))
	assert.Equal(uint64(9+2+16*2), w.n)
	assert.Equal(ErrNotFinalized, w.Close())
	assert.Nil(w.WriteIndex())
	assert.Nil(w.Close())

	// TODO: need to read stuff back
}