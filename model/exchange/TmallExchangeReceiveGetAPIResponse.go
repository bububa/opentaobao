package exchange

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeReceiveGetAPIResponse 卖家查询换货列表 API返回值
// tmall.exchange.receive.get
//
// 卖家查询换货列表
type TmallExchangeReceiveGetAPIResponse struct {
	model.CommonResponse
	TmallExchangeReceiveGetAPIResponseModel
}

// TmallExchangeReceiveGetAPIResponseModel is 卖家查询换货列表 成功返回结果
type TmallExchangeReceiveGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_receive_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Results []Exchange `json:"results,omitempty" xml:"results>exchange,omitempty"`
	// 所抛出异常
	Exception string `json:"exception,omitempty" xml:"exception,omitempty"`
	// 错误码
	ErrorCodes string `json:"error_codes,omitempty" xml:"error_codes,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 所有符合查询条件的换货单的数量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 当前页的换货单数量
	PageResults int64 `json:"page_results,omitempty" xml:"page_results,omitempty"`
	// 是否还有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}
