package exchange

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangegetAPIResponse 获取单笔换货详情 API返回值
// tmall.exchange.get
//
// 获取单笔换货详情
type TmallexchangegetAPIResponse struct {
	model.CommonResponse
	TmallexchangegetAPIResponseModel
}

// TmallexchangegetAPIResponseModel is 获取单笔换货详情 成功返回结果
type TmallexchangegetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
