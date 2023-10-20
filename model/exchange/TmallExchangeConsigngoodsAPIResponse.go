package exchange

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangeconsigngoodsAPIResponse 卖家发货 API返回值
// tmall.exchange.consigngoods
//
// 卖家发货
type TmallexchangeconsigngoodsAPIResponse struct {
	model.CommonResponse
	TmallexchangeconsigngoodsAPIResponseModel
}

// TmallexchangeconsigngoodsAPIResponseModel is 卖家发货 成功返回结果
type TmallexchangeconsigngoodsAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_consigngoods_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
