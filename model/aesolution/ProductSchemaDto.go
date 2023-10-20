package aesolution

import (
	"sync"
)

// ProductSchemaDto 结构体
type ProductSchemaDto struct {
	// error code
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// error message
	ErrorMessage string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// schema
	Schema string `json:"schema,omitempty" xml:"schema,omitempty"`
	// success flag
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolProductSchemaDto = sync.Pool{
	New: func() any {
		return new(ProductSchemaDto)
	},
}

// GetProductSchemaDto() 从对象池中获取ProductSchemaDto
func GetProductSchemaDto() *ProductSchemaDto {
	return poolProductSchemaDto.Get().(*ProductSchemaDto)
}

// ReleaseProductSchemaDto 释放ProductSchemaDto
func ReleaseProductSchemaDto(v *ProductSchemaDto) {
	v.ErrorCode = ""
	v.ErrorMessage = ""
	v.Schema = ""
	v.Success = false
	poolProductSchemaDto.Put(v)
}
