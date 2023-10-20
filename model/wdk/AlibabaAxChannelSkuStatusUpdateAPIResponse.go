package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaxchannelskustatusupdateAPIResponse 翱象渠道商品上下架接口 API返回值
// alibaba.ax.channel.sku.status.update
//
// 翱象渠道商品上下架接口
type AlibabaaxchannelskustatusupdateAPIResponse struct {
	model.CommonResponse
	AlibabaaxchannelskustatusupdateAPIResponseModel
}

// AlibabaaxchannelskustatusupdateAPIResponseModel is 翱象渠道商品上下架接口 成功返回结果
type AlibabaaxchannelskustatusupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ax_channel_sku_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用接口返回结果
	ApiResult *AlibabaaxchannelskustatusupdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}
