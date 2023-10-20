package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangConsignorderGiftReverseAPIResponse 赠品绑赠回滚 API返回值
// alibaba.dchain.aoxiang.consignorder.gift.reverse
//
// 赠品绑赠回滚
type AlibabaDchainAoxiangConsignorderGiftReverseAPIResponse struct {
	model.CommonResponse
	AlibabaDchainAoxiangConsignorderGiftReverseAPIResponseModel
}

// AlibabaDchainAoxiangConsignorderGiftReverseAPIResponseModel is 赠品绑赠回滚 成功返回结果
type AlibabaDchainAoxiangConsignorderGiftReverseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dchain_aoxiang_consignorder_gift_reverse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结构体
	ReverseConsignorderGiftResponse *ReverseConsignOrderGiftRequest `json:"reverse_consignorder_gift_response,omitempty" xml:"reverse_consignorder_gift_response,omitempty"`
}
