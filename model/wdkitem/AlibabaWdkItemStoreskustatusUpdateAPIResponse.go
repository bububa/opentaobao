package wdkitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemStoreskustatusUpdateAPIResponse 修改门店商品状态 API返回值
// alibaba.wdk.item.storeskustatus.update
//
// 五道口商品 修改门店商品状态
type AlibabaWdkItemStoreskustatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkItemStoreskustatusUpdateAPIResponseModel
}

// AlibabaWdkItemStoreskustatusUpdateAPIResponseModel is 修改门店商品状态 成功返回结果
type AlibabaWdkItemStoreskustatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_item_storeskustatus_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkItemStoreskustatusUpdateResult `json:"result,omitempty" xml:"result,omitempty"`
}
