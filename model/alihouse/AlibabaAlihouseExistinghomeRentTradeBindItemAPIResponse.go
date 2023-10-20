package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihouseexistinghomerenttradebinditemAPIResponse 交易绑定商品 API返回值
// alibaba.alihouse.existinghome.rent.trade.bind.item
//
// 交易绑定商品
type AlibabaalihouseexistinghomerenttradebinditemAPIResponse struct {
	model.CommonResponse
	AlibabaalihouseexistinghomerenttradebinditemAPIResponseModel
}

// AlibabaalihouseexistinghomerenttradebinditemAPIResponseModel is 交易绑定商品 成功返回结果
type AlibabaalihouseexistinghomerenttradebinditemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_rent_trade_bind_item_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值对象
	Result *AlibabaalihouseexistinghomerenttradebinditemResult `json:"result,omitempty" xml:"result,omitempty"`
}
