package foodscan

import (
	"sync"
)

// FilePackageRequest 结构体
type FilePackageRequest struct {
	// 城市名称
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 淘宝的nickName
	NickName string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// 角色ID
	RoleId string `json:"role_id,omitempty" xml:"role_id,omitempty"`
	// 区域编码
	DistrictAdcode string `json:"district_adcode,omitempty" xml:"district_adcode,omitempty"`
	// 用户唯一标识，可以是淘宝用户Id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 设备系统平台
	Platform string `json:"platform,omitempty" xml:"platform,omitempty"`
	// 设备品牌
	MobileBrand string `json:"mobile_brand,omitempty" xml:"mobile_brand,omitempty"`
	// 省份名称
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 前8位是日期，后10位是随机字符串
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 城市代码
	CityAdcode string `json:"city_adcode,omitempty" xml:"city_adcode,omitempty"`
	// 国家代码
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// 手机型号
	MobileModel string `json:"mobile_model,omitempty" xml:"mobile_model,omitempty"`
	// 区域名称
	District string `json:"district,omitempty" xml:"district,omitempty"`
	// 第一张图片的文件名
	FileName1 string `json:"file_name1,omitempty" xml:"file_name1,omitempty"`
	// 第二张图片的文件名
	FileName2 string `json:"file_name2,omitempty" xml:"file_name2,omitempty"`
	// 第三张图片的文件名
	FileName3 string `json:"file_name3,omitempty" xml:"file_name3,omitempty"`
	// 用户关系类型0/本人 1/爱人 2/父母 3/朋友 4/子女 99/他人
	RelationType int64 `json:"relation_type,omitempty" xml:"relation_type,omitempty"`
	// 1男2女
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// 1左脚 2右脚
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolFilePackageRequest = sync.Pool{
	New: func() any {
		return new(FilePackageRequest)
	},
}

// GetFilePackageRequest() 从对象池中获取FilePackageRequest
func GetFilePackageRequest() *FilePackageRequest {
	return poolFilePackageRequest.Get().(*FilePackageRequest)
}

// ReleaseFilePackageRequest 释放FilePackageRequest
func ReleaseFilePackageRequest(v *FilePackageRequest) {
	v.City = ""
	v.NickName = ""
	v.RoleId = ""
	v.DistrictAdcode = ""
	v.UserId = ""
	v.Platform = ""
	v.MobileBrand = ""
	v.Province = ""
	v.RequestId = ""
	v.CityAdcode = ""
	v.CountryCode = ""
	v.MobileModel = ""
	v.District = ""
	v.FileName1 = ""
	v.FileName2 = ""
	v.FileName3 = ""
	v.RelationType = 0
	v.Gender = 0
	v.Type = 0
	poolFilePackageRequest.Put(v)
}
