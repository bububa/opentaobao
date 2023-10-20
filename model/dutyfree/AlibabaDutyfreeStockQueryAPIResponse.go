package dutyfree

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDutyfreeStockQueryAPIResponse 对外库存查询接口 API返回值
// alibaba.dutyfree.stock.query
//
// 对外部服务提供库存查询接口
type AlibabaDutyfreeStockQueryAPIResponse struct {
	model.CommonResponse
	AlibabaDutyfreeStockQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDutyfreeStockQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDutyfreeStockQueryAPIResponseModel).Reset()
}

// AlibabaDutyfreeStockQueryAPIResponseModel is 对外库存查询接口 成功返回结果
type AlibabaDutyfreeStockQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dutyfree_stock_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaDutyfreeStockQueryResultDo `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDutyfreeStockQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaDutyfreeStockQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDutyfreeStockQueryAPIResponse)
	},
}

// GetAlibabaDutyfreeStockQueryAPIResponse 从 sync.Pool 获取 AlibabaDutyfreeStockQueryAPIResponse
func GetAlibabaDutyfreeStockQueryAPIResponse() *AlibabaDutyfreeStockQueryAPIResponse {
	return poolAlibabaDutyfreeStockQueryAPIResponse.Get().(*AlibabaDutyfreeStockQueryAPIResponse)
}

// ReleaseAlibabaDutyfreeStockQueryAPIResponse 将 AlibabaDutyfreeStockQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDutyfreeStockQueryAPIResponse(v *AlibabaDutyfreeStockQueryAPIResponse) {
	v.Reset()
	poolAlibabaDutyfreeStockQueryAPIResponse.Put(v)
}
