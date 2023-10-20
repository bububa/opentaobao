package topoaid

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTopPackageAuthCheckAPIResponse 校验用户授权关系 API返回值
// taobao.top.package.auth.check
//
// 校验用户授权关系
type TaobaoTopPackageAuthCheckAPIResponse struct {
	model.CommonResponse
	TaobaoTopPackageAuthCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTopPackageAuthCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTopPackageAuthCheckAPIResponseModel).Reset()
}

// TaobaoTopPackageAuthCheckAPIResponseModel is 校验用户授权关系 成功返回结果
type TaobaoTopPackageAuthCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"top_package_auth_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 授权查询结果
	Result *AuthScopeCheckResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTopPackageAuthCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTopPackageAuthCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTopPackageAuthCheckAPIResponse)
	},
}

// GetTaobaoTopPackageAuthCheckAPIResponse 从 sync.Pool 获取 TaobaoTopPackageAuthCheckAPIResponse
func GetTaobaoTopPackageAuthCheckAPIResponse() *TaobaoTopPackageAuthCheckAPIResponse {
	return poolTaobaoTopPackageAuthCheckAPIResponse.Get().(*TaobaoTopPackageAuthCheckAPIResponse)
}

// ReleaseTaobaoTopPackageAuthCheckAPIResponse 将 TaobaoTopPackageAuthCheckAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTopPackageAuthCheckAPIResponse(v *TaobaoTopPackageAuthCheckAPIResponse) {
	v.Reset()
	poolTaobaoTopPackageAuthCheckAPIResponse.Put(v)
}
