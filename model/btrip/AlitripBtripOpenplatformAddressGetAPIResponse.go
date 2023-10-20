package btrip

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripbtripopenplatformaddressgetAPIResponse 【商旅】开放平台对外页面跳转 API返回值
// alitrip.btrip.openplatform.address.get
//
// 获取类目预定页跳转地址
type AlitripbtripopenplatformaddressgetAPIResponse struct {
	model.CommonResponse
	AlitripbtripopenplatformaddressgetAPIResponseModel
}

// AlitripbtripopenplatformaddressgetAPIResponseModel is 【商旅】开放平台对外页面跳转 成功返回结果
type AlitripbtripopenplatformaddressgetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_btrip_openplatform_address_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *HisvResult `json:"result,omitempty" xml:"result,omitempty"`
}
