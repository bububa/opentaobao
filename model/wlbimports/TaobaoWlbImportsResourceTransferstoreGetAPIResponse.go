package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsResourceTransferstoreGetAPIResponse 根据指定的资源获取所有中转仓列表 API返回值
// taobao.wlb.imports.resource.transferstore.get
//
// 根据指定的资源获取所有中转仓列表
type TaobaoWlbImportsResourceTransferstoreGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbImportsResourceTransferstoreGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbImportsResourceTransferstoreGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbImportsResourceTransferstoreGetAPIResponseModel).Reset()
}

// TaobaoWlbImportsResourceTransferstoreGetAPIResponseModel is 根据指定的资源获取所有中转仓列表 成功返回结果
type TaobaoWlbImportsResourceTransferstoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_imports_resource_transferstore_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 符合条件的中转仓列表
	Stores []TranStoreResult `json:"stores,omitempty" xml:"stores>tran_store_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbImportsResourceTransferstoreGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Stores = m.Stores[:0]
}

var poolTaobaoWlbImportsResourceTransferstoreGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportsResourceTransferstoreGetAPIResponse)
	},
}

// GetTaobaoWlbImportsResourceTransferstoreGetAPIResponse 从 sync.Pool 获取 TaobaoWlbImportsResourceTransferstoreGetAPIResponse
func GetTaobaoWlbImportsResourceTransferstoreGetAPIResponse() *TaobaoWlbImportsResourceTransferstoreGetAPIResponse {
	return poolTaobaoWlbImportsResourceTransferstoreGetAPIResponse.Get().(*TaobaoWlbImportsResourceTransferstoreGetAPIResponse)
}

// ReleaseTaobaoWlbImportsResourceTransferstoreGetAPIResponse 将 TaobaoWlbImportsResourceTransferstoreGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbImportsResourceTransferstoreGetAPIResponse(v *TaobaoWlbImportsResourceTransferstoreGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbImportsResourceTransferstoreGetAPIResponse.Put(v)
}
