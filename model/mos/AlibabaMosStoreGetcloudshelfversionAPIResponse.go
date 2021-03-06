package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosStoreGetcloudshelfversionAPIResponse 获取云货架版本信息 API返回值
// alibaba.mos.store.getcloudshelfversion
//
// 根据屏编号获取云货架版本信息
type AlibabaMosStoreGetcloudshelfversionAPIResponse struct {
	model.CommonResponse
	AlibabaMosStoreGetcloudshelfversionAPIResponseModel
}

// AlibabaMosStoreGetcloudshelfversionAPIResponseModel is 获取云货架版本信息 成功返回结果
type AlibabaMosStoreGetcloudshelfversionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_store_getcloudshelfversion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaMosStoreGetcloudshelfversionResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
