package jipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripsellerrefundsearchAPIResponse 【机票代理商】退票申请单列表 API返回值
// taobao.alitrip.seller.refund.search
//
// 查询退票申请单列表
type TaobaoalitripsellerrefundsearchAPIResponse struct {
	model.CommonResponse
	TaobaoalitripsellerrefundsearchAPIResponseModel
}

// TaobaoalitripsellerrefundsearchAPIResponseModel is 【机票代理商】退票申请单列表 成功返回结果
type TaobaoalitripsellerrefundsearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_seller_refund_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoalitripsellerrefundsearchResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
