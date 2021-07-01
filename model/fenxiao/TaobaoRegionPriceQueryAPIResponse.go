package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoRegionPriceQueryAPIResponse
区域价格查询 API返回值
taobao.region.price.query

区域价格查询 */
type TaobaoRegionPriceQueryAPIResponse struct {
	model.CommonResponse
	TaobaoRegionPriceQueryAPIResponseModel
}

// TaobaoRegionPriceQueryAPIResponseModel is 区域价格查询 成功返回结果
type TaobaoRegionPriceQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"region_price_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
