package ptr

import (
	"encoding/json"
	"fmt"
	"time"
)

func (p *Time) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, p.String())), nil
}

func (p *Time) UnmarshalJSON(data []byte) error {
	var (
		t   time.Time
		err error
	)
	t, err = time.Parse(time.RFC3339, string(data))
	if err != nil {
		err = json.Unmarshal(data, &t)
		if err != nil {
			return err
		}
	}
	if p != nil {
		*p = Time(t)
	}
	return nil
}
