package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallInventoryQueryForstoreAPIResponse 查询后端商品仓库库存 API返回值
// tmall.inventory.query.forstore
//
// 商家查询后端商品仓库库存
type TmallInventoryQueryForstoreAPIResponse struct {
	model.CommonResponse
	TmallInventoryQueryForstoreAPIResponseModel
}

// Reset 清空结构体
func (m *TmallInventoryQueryForstoreAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallInventoryQueryForstoreAPIResponseModel).Reset()
}

// TmallInventoryQueryForstoreAPIResponseModel is 查询后端商品仓库库存 成功返回结果
type TmallInventoryQueryForstoreAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_inventory_query_forstore_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误code
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// 错误信息
	Errormessage string `json:"errormessage,omitempty" xml:"errormessage,omitempty"`
	// 查询结果
	Result *InventoryQueryResult `json:"result,omitempty" xml:"result,omitempty"`
	// 整体成功或失败
	Issuccess bool `json:"issuccess,omitempty" xml:"issuccess,omitempty"`
}

// Reset 清空结构体
func (m *TmallInventoryQueryForstoreAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Errorcode = ""
	m.Errormessage = ""
	m.Result = nil
	m.Issuccess = false
}

var poolTmallInventoryQueryForstoreAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallInventoryQueryForstoreAPIResponse)
	},
}

// GetTmallInventoryQueryForstoreAPIResponse 从 sync.Pool 获取 TmallInventoryQueryForstoreAPIResponse
func GetTmallInventoryQueryForstoreAPIResponse() *TmallInventoryQueryForstoreAPIResponse {
	return poolTmallInventoryQueryForstoreAPIResponse.Get().(*TmallInventoryQueryForstoreAPIResponse)
}

// ReleaseTmallInventoryQueryForstoreAPIResponse 将 TmallInventoryQueryForstoreAPIResponse 保存到 sync.Pool
func ReleaseTmallInventoryQueryForstoreAPIResponse(v *TmallInventoryQueryForstoreAPIResponse) {
	v.Reset()
	poolTmallInventoryQueryForstoreAPIResponse.Put(v)
}
