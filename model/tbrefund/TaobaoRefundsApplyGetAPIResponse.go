package tbrefund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundsapplygetAPIResponse 查询买家申请的退款列表 API返回值
// taobao.refunds.apply.get
//
// 查询买家申请的退款列表，且查询外店的退款列表时需要指定交易类型
type TaobaorefundsapplygetAPIResponse struct {
	model.CommonResponse
	TaobaorefundsapplygetAPIResponseModel
}

// TaobaorefundsapplygetAPIResponseModel is 查询买家申请的退款列表 成功返回结果
type TaobaorefundsapplygetAPIResponseModel struct {
	XMLName xml.Name `xml:"refunds_apply_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的退款信息列表
	Refunds []Refund `json:"refunds,omitempty" xml:"refunds>refund,omitempty"`
	// 搜索到的交易信息总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}
