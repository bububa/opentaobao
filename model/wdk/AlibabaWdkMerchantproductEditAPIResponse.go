package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmerchantproducteditAPIResponse 商家产品服务-编辑产品 API返回值
// alibaba.wdk.merchantproduct.edit
//
// 商家产品服务-编辑产品
type AlibabawdkmerchantproducteditAPIResponse struct {
	model.CommonResponse
	AlibabawdkmerchantproducteditAPIResponseModel
}

// AlibabawdkmerchantproducteditAPIResponseModel is 商家产品服务-编辑产品 成功返回结果
type AlibabawdkmerchantproducteditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_merchantproduct_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品编辑返回结果
	Result *AlibabawdkmerchantproducteditApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
