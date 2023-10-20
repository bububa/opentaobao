package product

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TmallItemSeriesItemseriesInsertorupdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallItemSeriesItemseriesInsertorupdateAPIResponseModel).Reset()
}

// TmallItemSeriesItemseriesInsertorupdateAPIResponseModel is 商品系列增删改接口 成功返回结果
type TmallItemSeriesItemseriesInsertorupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_item_series_itemseries_insertorupdate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *ResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallItemSeriesItemseriesInsertorupdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallItemSeriesItemseriesInsertorupdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallItemSeriesItemseriesInsertorupdateAPIResponse)
	},
}

// GetTmallItemSeriesItemseriesInsertorupdateAPIResponse 从 sync.Pool 获取 TmallItemSeriesItemseriesInsertorupdateAPIResponse
func GetTmallItemSeriesItemseriesInsertorupdateAPIResponse() *TmallItemSeriesItemseriesInsertorupdateAPIResponse {
	return poolTmallItemSeriesItemseriesInsertorupdateAPIResponse.Get().(*TmallItemSeriesItemseriesInsertorupdateAPIResponse)
}

// ReleaseTmallItemSeriesItemseriesInsertorupdateAPIResponse 将 TmallItemSeriesItemseriesInsertorupdateAPIResponse 保存到 sync.Pool
func ReleaseTmallItemSeriesItemseriesInsertorupdateAPIResponse(v *TmallItemSeriesItemseriesInsertorupdateAPIResponse) {
	v.Reset()
	poolTmallItemSeriesItemseriesInsertorupdateAPIResponse.Put(v)
}
