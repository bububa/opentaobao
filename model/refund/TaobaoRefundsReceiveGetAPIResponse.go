package refund

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundsreceivegetAPIResponse 查询卖家收到的退款列表 API返回值
// taobao.refunds.receive.get
//
// 查询卖家收到的退款列表
type TaobaorefundsreceivegetAPIResponse struct {
	model.CommonResponse
	TaobaorefundsreceivegetAPIResponseModel
}

// TaobaorefundsreceivegetAPIResponseModel is 查询卖家收到的退款列表 成功返回结果
type TaobaorefundsreceivegetAPIResponseModel struct {
	XMLName xml.Name `xml:"refunds_receive_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 搜索到的退款信息列表
	Refunds []Refund `json:"refunds,omitempty" xml:"refunds>refund,omitempty"`
	// 搜索到的退款信息总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
