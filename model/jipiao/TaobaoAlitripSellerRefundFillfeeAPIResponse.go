package jipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellerrefundfillfeeAPIResponse 机票代理商】回填手续费 API返回值
// taobao.alitrip.seller.refund.fillfee
//
// 回填手续费
type TaobaoalitripsellerrefundfillfeeAPIResponse struct {
	model.CommonResponse
	TaobaoalitripsellerrefundfillfeeAPIResponseModel
}

// TaobaoalitripsellerrefundfillfeeAPIResponseModel is 机票代理商】回填手续费 成功返回结果
type TaobaoalitripsellerrefundfillfeeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refund_fillfee_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
