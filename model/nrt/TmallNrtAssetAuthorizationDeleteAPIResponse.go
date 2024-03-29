package nrt

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtAssetAuthorizationDeleteAPIResponse 移除资产数据权限授权关系 API返回值
// tmall.nrt.asset.authorization.delete
//
// 移除资产数据权限授权关系
type TmallNrtAssetAuthorizationDeleteAPIResponse struct {
	model.CommonResponse
	TmallNrtAssetAuthorizationDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *TmallNrtAssetAuthorizationDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallNrtAssetAuthorizationDeleteAPIResponseModel).Reset()
}

// TmallNrtAssetAuthorizationDeleteAPIResponseModel is 移除资产数据权限授权关系 成功返回结果
type TmallNrtAssetAuthorizationDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_nrt_asset_authorization_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求结果
	Data *TopAssetDataAuthResultDto `json:"data,omitempty" xml:"data,omitempty"`
	// 请求是否成功
	Succ bool `json:"succ,omitempty" xml:"succ,omitempty"`
}

// Reset 清空结构体
func (m *TmallNrtAssetAuthorizationDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
	m.Succ = false
}

var poolTmallNrtAssetAuthorizationDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallNrtAssetAuthorizationDeleteAPIResponse)
	},
}

// GetTmallNrtAssetAuthorizationDeleteAPIResponse 从 sync.Pool 获取 TmallNrtAssetAuthorizationDeleteAPIResponse
func GetTmallNrtAssetAuthorizationDeleteAPIResponse() *TmallNrtAssetAuthorizationDeleteAPIResponse {
	return poolTmallNrtAssetAuthorizationDeleteAPIResponse.Get().(*TmallNrtAssetAuthorizationDeleteAPIResponse)
}

// ReleaseTmallNrtAssetAuthorizationDeleteAPIResponse 将 TmallNrtAssetAuthorizationDeleteAPIResponse 保存到 sync.Pool
func ReleaseTmallNrtAssetAuthorizationDeleteAPIResponse(v *TmallNrtAssetAuthorizationDeleteAPIResponse) {
	v.Reset()
	poolTmallNrtAssetAuthorizationDeleteAPIResponse.Put(v)
}
