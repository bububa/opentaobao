package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinAxbVendorHeartBeatAPIResponse
供应商心跳上报接口 API返回值
alibaba.aliqin.axb.vendor.heart.beat

供应商上报自己的心跳信息 */
type AlibabaAliqinAxbVendorHeartBeatAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinAxbVendorHeartBeatAPIResponseModel
}

// AlibabaAliqinAxbVendorHeartBeatAPIResponseModel is 供应商心跳上报接口 成功返回结果
type AlibabaAliqinAxbVendorHeartBeatAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_heart_beat_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaAliqinAxbVendorHeartBeatResponse `json:"result,omitempty" xml:"result,omitempty"`
}
