package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSeriesItemseriesInsertorupdateAPIResponse 商品系列增删改接口 API返回值
// tmall.item.series.itemseries.insertorupdate
//
// 商品系列增删改接口
type TmallItemSeriesItemseriesInsertorupdateAPIResponse struct {
	model.CommonResponse
	TmallItemSeriesItemseriesInsertorupdateAPIResponseModel
}

// TmallItemSeriesItemseriesInsertorupdateAPIResponseModel is 商品系列增删改接口 成功返回结果
type TmallItemSeriesItemseriesInsertorupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_series_itemseries_insertorupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
