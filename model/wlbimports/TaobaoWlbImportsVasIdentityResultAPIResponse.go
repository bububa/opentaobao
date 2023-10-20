package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsVasIdentityResultAPIResponse 集货鉴定结果 API返回值
// taobao.wlb.imports.vas.identity.result
//
// 集货鉴定结果查询
type TaobaoWlbImportsVasIdentityResultAPIResponse struct {
	model.CommonResponse
	TaobaoWlbImportsVasIdentityResultAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbImportsVasIdentityResultAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbImportsVasIdentityResultAPIResponseModel).Reset()
}

// TaobaoWlbImportsVasIdentityResultAPIResponseModel is 集货鉴定结果 成功返回结果
type TaobaoWlbImportsVasIdentityResultAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_imports_vas_identity_result_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回出参数结果
	Result *TaobaoWlbImportsVasIdentityResultTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbImportsVasIdentityResultAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoWlbImportsVasIdentityResultAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportsVasIdentityResultAPIResponse)
	},
}

// GetTaobaoWlbImportsVasIdentityResultAPIResponse 从 sync.Pool 获取 TaobaoWlbImportsVasIdentityResultAPIResponse
func GetTaobaoWlbImportsVasIdentityResultAPIResponse() *TaobaoWlbImportsVasIdentityResultAPIResponse {
	return poolTaobaoWlbImportsVasIdentityResultAPIResponse.Get().(*TaobaoWlbImportsVasIdentityResultAPIResponse)
}

// ReleaseTaobaoWlbImportsVasIdentityResultAPIResponse 将 TaobaoWlbImportsVasIdentityResultAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbImportsVasIdentityResultAPIResponse(v *TaobaoWlbImportsVasIdentityResultAPIResponse) {
	v.Reset()
	poolTaobaoWlbImportsVasIdentityResultAPIResponse.Put(v)
}
