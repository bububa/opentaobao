package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenWarehouseinfoSynchronizeAPIResponse 仓库同步接口 API返回值
// taobao.qimen.warehouseinfo.synchronize
//
// 仓库同步接口
type TaobaoQimenWarehouseinfoSynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenWarehouseinfoSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenWarehouseinfoSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenWarehouseinfoSynchronizeAPIResponseModel).Reset()
}

// TaobaoQimenWarehouseinfoSynchronizeAPIResponseModel is 仓库同步接口 成功返回结果
type TaobaoQimenWarehouseinfoSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_warehouseinfo_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应报文
	Response *TaobaoQimenWarehouseinfoSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenWarehouseinfoSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenWarehouseinfoSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenWarehouseinfoSynchronizeAPIResponse)
	},
}

// GetTaobaoQimenWarehouseinfoSynchronizeAPIResponse 从 sync.Pool 获取 TaobaoQimenWarehouseinfoSynchronizeAPIResponse
func GetTaobaoQimenWarehouseinfoSynchronizeAPIResponse() *TaobaoQimenWarehouseinfoSynchronizeAPIResponse {
	return poolTaobaoQimenWarehouseinfoSynchronizeAPIResponse.Get().(*TaobaoQimenWarehouseinfoSynchronizeAPIResponse)
}

// ReleaseTaobaoQimenWarehouseinfoSynchronizeAPIResponse 将 TaobaoQimenWarehouseinfoSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenWarehouseinfoSynchronizeAPIResponse(v *TaobaoQimenWarehouseinfoSynchronizeAPIResponse) {
	v.Reset()
	poolTaobaoQimenWarehouseinfoSynchronizeAPIResponse.Put(v)
}
