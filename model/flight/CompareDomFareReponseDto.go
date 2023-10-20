package flight

import (
	"sync"
)

// CompareDomFareReponseDto 结构体
type CompareDomFareReponseDto struct {
	// 返回政策信息
	PriceComparisonList []PriceComparisonDto `json:"price_comparison_list,omitempty" xml:"price_comparison_list>price_comparison_dto,omitempty"`
	// 返回信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolCompareDomFareReponseDto = sync.Pool{
	New: func() any {
		return new(CompareDomFareReponseDto)
	},
}

// GetCompareDomFareReponseDto() 从对象池中获取CompareDomFareReponseDto
func GetCompareDomFareReponseDto() *CompareDomFareReponseDto {
	return poolCompareDomFareReponseDto.Get().(*CompareDomFareReponseDto)
}

// ReleaseCompareDomFareReponseDto 释放CompareDomFareReponseDto
func ReleaseCompareDomFareReponseDto(v *CompareDomFareReponseDto) {
	v.PriceComparisonList = v.PriceComparisonList[:0]
	v.Message = ""
	poolCompareDomFareReponseDto.Put(v)
}
