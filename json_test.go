package ptr

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeJson(t *testing.T) {
	as := assert.New(t)
	now := time.Now()
	tPtr := NewTime(now)
	j, err := json.Marshal(tPtr)
	as.NoError(err)
	as.Equal([]byte(fmt.Sprintf(`"%s"`, tPtr.String())), j)

	nowJ, err := json.Marshal(&now)
	as.NoError(err)
	as.NoError(json.Unmarshal(nowJ, tPtr))
	as.True(tPtr.Value().Equal(now))

	as.NoError(json.Unmarshal(j, tPtr))
	as.True(tPtr.Value().Equal(now))

}
