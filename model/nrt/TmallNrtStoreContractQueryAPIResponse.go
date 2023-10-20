package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtStoreContractQueryAPIResponse 摊位合同查询接口 API返回值
// tmall.nrt.store.contract.query
//
// 摊位合同查询接口
type TmallNrtStoreContractQueryAPIResponse struct {
	model.CommonResponse
	TmallNrtStoreContractQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtStoreContractQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtStoreContractQueryAPIResponseModel).Reset()
}

// TmallNrtStoreContractQueryAPIResponseModel is 摊位合同查询接口 成功返回结果
type TmallNrtStoreContractQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_store_contract_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数对象
	Data []EaStoreContractDto `json:"data,omitempty" xml:"data>ea_store_contract_dto,omitempty"`
	// 错误码
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// 错误信息
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// 成功与否
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtStoreContractQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.Errcode = ""
	m.Errmsg = ""
	m.Succ = false
}

var poolTmallNrtStoreContractQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtStoreContractQueryAPIResponse)
	},
}

// GetTmallNrtStoreContractQueryAPIResponse 从 sync.Pool 获取 TmallNrtStoreContractQueryAPIResponse
func GetTmallNrtStoreContractQueryAPIResponse() *TmallNrtStoreContractQueryAPIResponse {
	return poolTmallNrtStoreContractQueryAPIResponse.Get().(*TmallNrtStoreContractQueryAPIResponse)
}

// ReleaseTmallNrtStoreContractQueryAPIResponse 将 TmallNrtStoreContractQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtStoreContractQueryAPIResponse(v *TmallNrtStoreContractQueryAPIResponse) {
	v.Reset()
	poolTmallNrtStoreContractQueryAPIResponse.Put(v)
}
