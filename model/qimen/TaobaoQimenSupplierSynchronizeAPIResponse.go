package qimen

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenSupplierSynchronizeAPIResponse 供应商同步接口 API返回值
// taobao.qimen.supplier.synchronize
//
// 这个接口用来同步供应商信息
type TaobaoQimenSupplierSynchronizeAPIResponse struct {
	model.CommonResponse
	TaobaoQimenSupplierSynchronizeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoQimenSupplierSynchronizeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoQimenSupplierSynchronizeAPIResponseModel).Reset()
}

// TaobaoQimenSupplierSynchronizeAPIResponseModel is 供应商同步接口 成功返回结果
type TaobaoQimenSupplierSynchronizeAPIResponseModel struct {
	XMLName xml.Name `xml:"qimen_supplier_synchronize_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	//
	Response *TaobaoQimenSupplierSynchronizeResponse `json:"response,omitempty" xml:"response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoQimenSupplierSynchronizeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Response = nil
}

var poolTaobaoQimenSupplierSynchronizeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenSupplierSynchronizeAPIResponse)
	},
}

// GetTaobaoQimenSupplierSynchronizeAPIResponse 从 sync.Pool 获取 TaobaoQimenSupplierSynchronizeAPIResponse
func GetTaobaoQimenSupplierSynchronizeAPIResponse() *TaobaoQimenSupplierSynchronizeAPIResponse {
	return poolTaobaoQimenSupplierSynchronizeAPIResponse.Get().(*TaobaoQimenSupplierSynchronizeAPIResponse)
}

// ReleaseTaobaoQimenSupplierSynchronizeAPIResponse 将 TaobaoQimenSupplierSynchronizeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoQimenSupplierSynchronizeAPIResponse(v *TaobaoQimenSupplierSynchronizeAPIResponse) {
	v.Reset()
	poolTaobaoQimenSupplierSynchronizeAPIResponse.Put(v)
}
