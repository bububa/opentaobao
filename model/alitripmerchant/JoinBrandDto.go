package alitripmerchant

import (
	"sync"
)

// JoinBrandDto 结构体
type JoinBrandDto struct {
	// 品牌logo
	BrandLogo string `json:"brand_logo,omitempty" xml:"brand_logo,omitempty"`
	// 品牌跳转url
	LogoRedirectUrl string `json:"logo_redirect_url,omitempty" xml:"logo_redirect_url,omitempty"`
}

var poolJoinBrandDto = sync.Pool{
	New: func() any {
		return new(JoinBrandDto)
	},
}

// GetJoinBrandDto() 从对象池中获取JoinBrandDto
func GetJoinBrandDto() *JoinBrandDto {
	return poolJoinBrandDto.Get().(*JoinBrandDto)
}

// ReleaseJoinBrandDto 释放JoinBrandDto
func ReleaseJoinBrandDto(v *JoinBrandDto) {
	v.BrandLogo = ""
	v.LogoRedirectUrl = ""
	poolJoinBrandDto.Put(v)
}
