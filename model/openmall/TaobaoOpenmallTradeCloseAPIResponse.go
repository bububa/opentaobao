package openmall

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenmalltradecloseAPIResponse 关闭订单 API返回值
// taobao.openmall.trade.close
//
// 关闭订单
type TaobaoopenmalltradecloseAPIResponse struct {
	model.CommonResponse
	TaobaoopenmalltradecloseAPIResponseModel
}

// TaobaoopenmalltradecloseAPIResponseModel is 关闭订单 成功返回结果
type TaobaoopenmalltradecloseAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *TopTradeResultVo `json:"result,omitempty" xml:"result,omitempty"`
}
