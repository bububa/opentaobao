package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenWarehouseinfoQueryAPIResponse 货主仓库资源查询接口 API返回值
// taobao.qimen.warehouseinfo.query
//
// 货主仓库资源查询
type TaobaoQimenWarehouseinfoQueryAPIResponse struct {
	model.CommonResponse
	TaobaoQimenWarehouseinfoQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenWarehouseinfoQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenWarehouseinfoQueryAPIResponseModel).Reset()
}

// TaobaoQimenWarehouseinfoQueryAPIResponseModel is 货主仓库资源查询接口 成功返回结果
type TaobaoQimenWarehouseinfoQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_warehouseinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenWarehouseinfoQueryResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenWarehouseinfoQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenWarehouseinfoQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWarehouseinfoQueryAPIResponse)
	},
}

// GetTaobaoQimenWarehouseinfoQueryAPIResponse 从 sync.Pool 获取 TaobaoQimenWarehouseinfoQueryAPIResponse
func GetTaobaoQimenWarehouseinfoQueryAPIResponse() *TaobaoQimenWarehouseinfoQueryAPIResponse {
	return poolTaobaoQimenWarehouseinfoQueryAPIResponse.Get().(*TaobaoQimenWarehouseinfoQueryAPIResponse)
}

// ReleaseTaobaoQimenWarehouseinfoQueryAPIResponse 将 TaobaoQimenWarehouseinfoQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenWarehouseinfoQueryAPIResponse(v *TaobaoQimenWarehouseinfoQueryAPIResponse) {
	v.Reset()
	poolTaobaoQimenWarehouseinfoQueryAPIResponse.Put(v)
}
