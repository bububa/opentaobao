package ott

import (
	"sync"
)

// RatingDto 结构体
type RatingDto struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// rating
	Rating string `json:"rating,omitempty" xml:"rating,omitempty"`
}

var poolRatingDto = sync.Pool{
	New: func() any {
		return new(RatingDto)
	},
}

// GetRatingDto() 从对象池中获取RatingDto
func GetRatingDto() *RatingDto {
	return poolRatingDto.Get().(*RatingDto)
}

// ReleaseRatingDto 释放RatingDto
func ReleaseRatingDto(v *RatingDto) {
	v.Code = ""
	v.Rating = ""
	poolRatingDto.Put(v)
}
