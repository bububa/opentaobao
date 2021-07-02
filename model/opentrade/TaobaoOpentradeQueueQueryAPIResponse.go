package opentrade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeQueueQueryAPIResponse 尖货交易排队信息查询 API返回值
// taobao.opentrade.queue.query
//
// 尖货交易排队信息查询
type TaobaoOpentradeQueueQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeQueueQueryAPIResponseModel
}

// TaobaoOpentradeQueueQueryAPIResponseModel is 尖货交易排队信息查询 成功返回结果
type TaobaoOpentradeQueueQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_queue_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 总记录数
	TotalCount string `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 返回的排队用户数据
	Results []McUserDto `json:"results,omitempty" xml:"results>mc_user_dto,omitempty"`
}
