package util

import (
	"bytes"
	"net/url"
	"strings"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func GetBufferPool() *bytes.Buffer {
	return bufferPool.Get().(*bytes.Buffer)
}

func PutBufferPool(buf *bytes.Buffer) {
	buf.Reset()
	bufferPool.Put(buf)
}

var stringsBuilderPool = sync.Pool{
	New: func() any {
		return new(strings.Builder)
	},
}

func GetStringsBuilder() *strings.Builder {
	return stringsBuilderPool.Get().(*strings.Builder)
}

func PutStringsBuilder(b *strings.Builder) {
	b.Reset()
	stringsBuilderPool.Put(b)
}

var urlValuesPool = sync.Pool{
	New: func() any {
		return make(url.Values)
	},
}

func GetUrlValues() url.Values {
	vals := urlValuesPool.Get().(url.Values)
	for k := range vals {
		vals.Del(k)
	}
	return vals
}

func PutUrlValues(vals url.Values) {
	urlValuesPool.Put(vals)
}
