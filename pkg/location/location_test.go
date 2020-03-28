package location

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"setzna/pkg/location"
)

func TestDistance(t *testing.T) {
	assert.Equal(t, 3.523778708998686, location.Distance(35.658482, 139.701441, 35.690224, 139.700089))
}
