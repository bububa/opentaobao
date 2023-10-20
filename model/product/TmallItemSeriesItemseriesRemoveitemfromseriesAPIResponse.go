package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallitemseriesitemseriesremoveitemfromseriesAPIResponse 从商品系列中移除商品 API返回值
// tmall.item.series.itemseries.removeitemfromseries
//
// 从商品系列中移除商品
type TmallitemseriesitemseriesremoveitemfromseriesAPIResponse struct {
	model.CommonResponse
	TmallitemseriesitemseriesremoveitemfromseriesAPIResponseModel
}

// TmallitemseriesitemseriesremoveitemfromseriesAPIResponseModel is 从商品系列中移除商品 成功返回结果
type TmallitemseriesitemseriesremoveitemfromseriesAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_series_itemseries_removeitemfromseries_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求返回值
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
