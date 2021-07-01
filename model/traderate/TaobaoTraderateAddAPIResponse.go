package traderate

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTraderateAddAPIResponse
新增单个评价 API返回值
taobao.traderate.add

新增单个评价(<font color="red">注：在评价之前需要对订单成功的时间进行判定（end_time）,如果超过15天，不能再通过该接口进行评价</font>) */
type TaobaoTraderateAddAPIResponse struct {
	model.CommonResponse
	TaobaoTraderateAddAPIResponseModel
}

// TaobaoTraderateAddAPIResponseModel is 新增单个评价 成功返回结果
type TaobaoTraderateAddAPIResponseModel struct {
	XMLName xml.Name `xml:"traderate_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回tid、oid、create
	TradeRate *TradeRateRequest `json:"trade_rate,omitempty" xml:"trade_rate,omitempty"`
}
