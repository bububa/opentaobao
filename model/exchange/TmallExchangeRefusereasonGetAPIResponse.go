package exchange

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeRefusereasonGetAPIResponse 获取拒绝换货原因列表 API返回值
// tmall.exchange.refusereason.get
//
// 获取拒绝换货原因列表
type TmallExchangeRefusereasonGetAPIResponse struct {
	model.CommonResponse
	TmallExchangeRefusereasonGetAPIResponseModel
}

// TmallExchangeRefusereasonGetAPIResponseModel is 获取拒绝换货原因列表 成功返回结果
type TmallExchangeRefusereasonGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_refusereason_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TmallExchangeRefusereasonGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
