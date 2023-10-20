package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangReceiverinfoQueryAPIResponse 供应链优仓即时配隐私小号查询 API返回值
// alibaba.dchain.aoxiang.receiverinfo.query
//
// 供应链优仓即时配隐私小号查询
type AlibabaDchainAoxiangReceiverinfoQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangReceiverinfoQueryAPIResponseModel
}

// AlibabaDchainAoxiangReceiverinfoQueryAPIResponseModel is 供应链优仓即时配隐私小号查询 成功返回结果
type AlibabaDchainAoxiangReceiverinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_receiverinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 用户信息
	OrderReceiverPrivacyResponse *OrderReceiverPrivacyResponse `json:"order_receiver_privacy_response,omitempty" xml:"order_receiver_privacy_response,omitempty"`
}
