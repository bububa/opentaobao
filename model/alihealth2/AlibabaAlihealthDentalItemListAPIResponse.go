package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalItemListAPIResponse ISV获取口腔标品列表 API返回值
// alibaba.alihealth.dental.item.list
//
// ISV获取口腔标品列表
type AlibabaAlihealthDentalItemListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthDentalItemListAPIResponseModel
}

// AlibabaAlihealthDentalItemListAPIResponseModel is ISV获取口腔标品列表 成功返回结果
type AlibabaAlihealthDentalItemListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_dental_item_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAlihealthDentalItemListMtopResult `json:"result,omitempty" xml:"result,omitempty"`
}
