package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamosstoregetdefautitemsAPIResponse 获取默认状态下商品列表 API返回值
// alibaba.mos.store.getdefautitems
//
// 获取默认状态下商品列表
type AlibabamosstoregetdefautitemsAPIResponse struct {
	model.CommonResponse
	AlibabamosstoregetdefautitemsAPIResponseModel
}

// AlibabamosstoregetdefautitemsAPIResponseModel is 获取默认状态下商品列表 成功返回结果
type AlibabamosstoregetdefautitemsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_store_getdefautitems_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabamosstoregetdefautitemsResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
