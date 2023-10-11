package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSeriesItemseriesInsertseriesitemAPIResponse 向系列中添加系列商品 API返回值
// tmall.item.series.itemseries.insertseriesitem
//
// 向系列中添加系列商品
type TmallItemSeriesItemseriesInsertseriesitemAPIResponse struct {
	model.CommonResponse
	TmallItemSeriesItemseriesInsertseriesitemAPIResponseModel
}

// TmallItemSeriesItemseriesInsertseriesitemAPIResponseModel is 向系列中添加系列商品 成功返回结果
type TmallItemSeriesItemseriesInsertseriesitemAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_series_itemseries_insertseriesitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
