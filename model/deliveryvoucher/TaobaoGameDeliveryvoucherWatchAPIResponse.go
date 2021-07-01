package deliveryvoucher

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoGameDeliveryvoucherWatchAPIResponse
监控预约数据 API返回值
taobao.game.deliveryvoucher.watch

监控预约数据 */
type TaobaoGameDeliveryvoucherWatchAPIResponse struct {
	model.CommonResponse
	TaobaoGameDeliveryvoucherWatchAPIResponseModel
}

// TaobaoGameDeliveryvoucherWatchAPIResponseModel is 监控预约数据 成功返回结果
type TaobaoGameDeliveryvoucherWatchAPIResponseModel struct {
	XMLName xml.Name `xml:"game_deliveryvoucher_watch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// isSuccess
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
