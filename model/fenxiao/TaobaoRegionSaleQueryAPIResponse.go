package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoregionsalequeryAPIResponse 查询商品销售区域 API返回值
// taobao.region.sale.query
//
// 查询商品销售区域
type TaobaoregionsalequeryAPIResponse struct {
	model.CommonResponse
	TaobaoregionsalequeryAPIResponseModel
}

// TaobaoregionsalequeryAPIResponseModel is 查询商品销售区域 成功返回结果
type TaobaoregionsalequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"region_sale_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
