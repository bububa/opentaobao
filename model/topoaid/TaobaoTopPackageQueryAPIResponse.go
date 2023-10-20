package topoaid

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageQueryAPIResponse 淘系包裹查询 API返回值
// taobao.top.package.query
//
// 淘系包裹查询
type TaobaoTopPackageQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTopPackageQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopPackageQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopPackageQueryAPIResponseModel).Reset()
}

// TaobaoTopPackageQueryAPIResponseModel is 淘系包裹查询 成功返回结果
type TaobaoTopPackageQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"top_package_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 包裹信息列表
	Data []PackageInfo `json:"data,omitempty" xml:"data>package_info,omitempty"`
	// 面单总数量
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopPackageQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.TotalNum = 0
}

var poolTaobaoTopPackageQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopPackageQueryAPIResponse)
	},
}

// GetTaobaoTopPackageQueryAPIResponse 从 sync.Pool 获取 TaobaoTopPackageQueryAPIResponse
func GetTaobaoTopPackageQueryAPIResponse() *TaobaoTopPackageQueryAPIResponse {
	return poolTaobaoTopPackageQueryAPIResponse.Get().(*TaobaoTopPackageQueryAPIResponse)
}

// ReleaseTaobaoTopPackageQueryAPIResponse 将 TaobaoTopPackageQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopPackageQueryAPIResponse(v *TaobaoTopPackageQueryAPIResponse) {
	v.Reset()
	poolTaobaoTopPackageQueryAPIResponse.Put(v)
}
