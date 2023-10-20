package exchange

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangereturngoodsagreeAPIResponse 卖家确认收货 API返回值
// tmall.exchange.returngoods.agree
//
// 卖家确认收货
type TmallexchangereturngoodsagreeAPIResponse struct {
	model.CommonResponse
	TmallexchangereturngoodsagreeAPIResponseModel
}

// TmallexchangereturngoodsagreeAPIResponseModel is 卖家确认收货 成功返回结果
type TmallexchangereturngoodsagreeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_returngoods_agree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
