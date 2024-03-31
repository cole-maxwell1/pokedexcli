package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Nanosecond * 5)

	if cache.cache == nil {
		t.Error("Expected a nil value returned from cache")
	}
}

func TestGetFromCache(t *testing.T) {
	cache := NewCache(time.Nanosecond * 5)

	cases := []struct {
		testKey   string
		testValue []byte
	}{
		{
			testKey:   "myKey",
			testValue: []byte("This is indeed a test!"),
		},
		{
			testKey:   "",
			testValue: []byte("ANOTHER VAL"),
		},
		{
			testKey:   "key2",
			testValue: []byte(""),
		},
		{
			testKey:   "",
			testValue: []byte(""),
		},
	}

	for _, c := range cases {
		cache.Add(c.testKey, c.testValue)

		actual, ok := cache.Get(c.testKey)
		if !ok {
			t.Error("Expected an OK value returned from cache")
			continue
		}

		if string(actual) != string(c.testValue) {
			t.Errorf("Bad value retrieved from cache: Expected %s, got %s",
				c.testValue,
				actual,
			)
		}

	}
}

func TestCacheReaperDestroy(t *testing.T) {
	timeToLive := time.Millisecond * 5
	testCache := NewCache(timeToLive)

	cases := []struct {
		testKey   string
		testValue []byte
	}{
		{
			testKey:   "myKey",
			testValue: []byte("This is indeed a test!"),
		},
		{
			testKey:   "",
			testValue: []byte("ANOTHER VAL"),
		},
		{
			testKey:   "key2",
			testValue: []byte(""),
		},
		{
			testKey:   "",
			testValue: []byte(""),
		},
	}

	for _, c := range cases {
		testCache.Add(c.testKey, c.testValue)
	}

	time.Sleep(timeToLive + time.Millisecond)

	if len(testCache.cache) != 0 {
		t.Errorf("Expected: 0 values returned from cache found %d", len(testCache.cache))
	}
}

func TestCacheReaperLifespan(t *testing.T) {
	timeToLive := time.Millisecond * 5
	testCache := NewCache(timeToLive)

	cases := []struct {
		testKey   string
		testValue []byte
	}{
		{
			testKey:   "myKey",
			testValue: []byte("This is indeed a test!"),
		},
		{
			testKey:   "",
			testValue: []byte("ANOTHER VAL"),
		},
		{
			testKey:   "key2",
			testValue: []byte(""),
		},
		{
			testKey:   "",
			testValue: []byte(""),
		},
	}

	for _, c := range cases {
		testCache.Add(c.testKey, c.testValue)
	}


	time.Sleep(timeToLive - 3 * time.Millisecond)

	_, ok := testCache.Get("myKey")

	if !ok {
		t.Error("Expected an OK value returned from cache for key: myKey")
	}

}
