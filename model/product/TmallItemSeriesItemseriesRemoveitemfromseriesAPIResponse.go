package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse 从商品系列中移除商品 API返回值
// tmall.item.series.itemseries.removeitemfromseries
//
// 从商品系列中移除商品
type TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse struct {
	model.CommonResponse
	TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponseModel
}

// Reset 清空结构体
func (m *TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponseModel).Reset()
}

// TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponseModel is 从商品系列中移除商品 成功返回结果
type TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_series_itemseries_removeitemfromseries_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求返回值
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse)
	},
}

// GetTmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse 从 sync.Pool 获取 TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse
func GetTmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse() *TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse {
	return poolTmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse.Get().(*TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse)
}

// ReleaseTmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse 将 TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse(v *TmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse) {
	v.Reset()
	poolTmallItemSeriesItemseriesRemoveitemfromseriesAPIResponse.Put(v)
}
