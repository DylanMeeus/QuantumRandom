package pkg

import (
	"testing"
)

func TestQrngCache(t *testing.T) {
	next, err := NextUint8()
	if err != nil {
		t.Errorf("could not get initial value: %v", err)
	}
	if next != uint8(uint8cache.data[0]) {
		t.Errorf("initial value %v does not equal first cache value %v", next, uint8cache.data[0])
	}
	currentCache := uint8cache.data
	t.Logf("current cache %v\n", currentCache)
	for i := 1; i < len(currentCache); i++ {
		next, err = NextUint8()
		if err != nil {
			t.Errorf("could not get next value: %v", err)
		}
		if next != uint8(currentCache[i]) {
			t.Errorf("Cache mismatch, expected %v but got %v", currentCache[i], next)
		}
	}
	// cache should be exhausted by now
	if uint8cache.isExhausted() != true {
		t.Error("Expected cache to be exhausted")
	}
	_, err = NextUint8()
	if err != nil {
		t.Errorf("could not get next value: %v", err)
	}
	if uint8cache.isExhausted() == true {
		t.Error("Expected cache to no longer be exhausted")
	}
}
