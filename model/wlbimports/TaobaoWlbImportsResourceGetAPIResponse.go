package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbImportsResourceGetAPIResponse 获取所有服务列表 API返回值
// taobao.wlb.imports.resource.get
//
// 一般进口TOP接口，获取所有服务列表。
type TaobaoWlbImportsResourceGetAPIResponse struct {
	model.CommonResponse
	TaobaoWlbImportsResourceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoWlbImportsResourceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoWlbImportsResourceGetAPIResponseModel).Reset()
}

// TaobaoWlbImportsResourceGetAPIResponseModel is 获取所有服务列表 成功返回结果
type TaobaoWlbImportsResourceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"wlb_imports_resource_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 一般进口所有服务商列表
	Resources []ResourceResult `json:"resources,omitempty" xml:"resources>resource_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoWlbImportsResourceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Resources = m.Resources[:0]
}

var poolTaobaoWlbImportsResourceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoWlbImportsResourceGetAPIResponse)
	},
}

// GetTaobaoWlbImportsResourceGetAPIResponse 从 sync.Pool 获取 TaobaoWlbImportsResourceGetAPIResponse
func GetTaobaoWlbImportsResourceGetAPIResponse() *TaobaoWlbImportsResourceGetAPIResponse {
	return poolTaobaoWlbImportsResourceGetAPIResponse.Get().(*TaobaoWlbImportsResourceGetAPIResponse)
}

// ReleaseTaobaoWlbImportsResourceGetAPIResponse 将 TaobaoWlbImportsResourceGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoWlbImportsResourceGetAPIResponse(v *TaobaoWlbImportsResourceGetAPIResponse) {
	v.Reset()
	poolTaobaoWlbImportsResourceGetAPIResponse.Put(v)
}
