package topoaid

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageUnauthQueryAPIResponse 查询某手机号下的包裹 API返回值
// taobao.top.package.unauth.query
//
// 查询某手机号下的包裹
type TaobaoTopPackageUnauthQueryAPIResponse struct {
	model.CommonResponse
	TaobaoTopPackageUnauthQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopPackageUnauthQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopPackageUnauthQueryAPIResponseModel).Reset()
}

// TaobaoTopPackageUnauthQueryAPIResponseModel is 查询某手机号下的包裹 成功返回结果
type TaobaoTopPackageUnauthQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"top_package_unauth_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 包裹信息
	Data []PackageInfo `json:"data,omitempty" xml:"data>package_info,omitempty"`
	// 总数量
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopPackageUnauthQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = m.Data[:0]
	m.TotalNum = 0
}

var poolTaobaoTopPackageUnauthQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopPackageUnauthQueryAPIResponse)
	},
}

// GetTaobaoTopPackageUnauthQueryAPIResponse 从 sync.Pool 获取 TaobaoTopPackageUnauthQueryAPIResponse
func GetTaobaoTopPackageUnauthQueryAPIResponse() *TaobaoTopPackageUnauthQueryAPIResponse {
	return poolTaobaoTopPackageUnauthQueryAPIResponse.Get().(*TaobaoTopPackageUnauthQueryAPIResponse)
}

// ReleaseTaobaoTopPackageUnauthQueryAPIResponse 将 TaobaoTopPackageUnauthQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopPackageUnauthQueryAPIResponse(v *TaobaoTopPackageUnauthQueryAPIResponse) {
	v.Reset()
	poolTaobaoTopPackageUnauthQueryAPIResponse.Put(v)
}
