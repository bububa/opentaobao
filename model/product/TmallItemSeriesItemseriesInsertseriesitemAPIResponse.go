package product

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TmallItemSeriesItemseriesInsertseriesitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSeriesItemseriesInsertseriesitemAPIResponseModel).Reset()
}

// TmallItemSeriesItemseriesInsertseriesitemAPIResponseModel is 向系列中添加系列商品 成功返回结果
type TmallItemSeriesItemseriesInsertseriesitemAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_series_itemseries_insertseriesitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSeriesItemseriesInsertseriesitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallItemSeriesItemseriesInsertseriesitemAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSeriesItemseriesInsertseriesitemAPIResponse)
	},
}

// GetTmallItemSeriesItemseriesInsertseriesitemAPIResponse 从 sync.Pool 获取 TmallItemSeriesItemseriesInsertseriesitemAPIResponse
func GetTmallItemSeriesItemseriesInsertseriesitemAPIResponse() *TmallItemSeriesItemseriesInsertseriesitemAPIResponse {
	return poolTmallItemSeriesItemseriesInsertseriesitemAPIResponse.Get().(*TmallItemSeriesItemseriesInsertseriesitemAPIResponse)
}

// ReleaseTmallItemSeriesItemseriesInsertseriesitemAPIResponse 将 TmallItemSeriesItemseriesInsertseriesitemAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSeriesItemseriesInsertseriesitemAPIResponse(v *TmallItemSeriesItemseriesInsertseriesitemAPIResponse) {
	v.Reset()
	poolTmallItemSeriesItemseriesInsertseriesitemAPIResponse.Put(v)
}
