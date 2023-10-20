package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkskuchannelskuupdateAPIResponse 更新渠道商品 API返回值
// alibaba.wdk.sku.channelsku.update
//
// 批量更新渠道商品，商家通过Top接入
type AlibabawdkskuchannelskuupdateAPIResponse struct {
	model.CommonResponse
	AlibabawdkskuchannelskuupdateAPIResponseModel
}

// AlibabawdkskuchannelskuupdateAPIResponseModel is 更新渠道商品 成功返回结果
type AlibabawdkskuchannelskuupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_sku_channelsku_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用结果
	Result *AlibabawdkskuchannelskuupdateApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
