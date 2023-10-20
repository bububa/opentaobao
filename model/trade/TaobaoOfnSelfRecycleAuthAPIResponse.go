package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOfnSelfRecycleAuthAPIResponse 自助回收鉴权 API返回值
// taobao.ofn.self.recycle.auth
//
// 自助回收鉴权
type TaobaoOfnSelfRecycleAuthAPIResponse struct {
	model.CommonResponse
	TaobaoOfnSelfRecycleAuthAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOfnSelfRecycleAuthAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOfnSelfRecycleAuthAPIResponseModel).Reset()
}

// TaobaoOfnSelfRecycleAuthAPIResponseModel is 自助回收鉴权 成功返回结果
type TaobaoOfnSelfRecycleAuthAPIResponseModel struct {
	XMLName xml.Name `xml:"ofn_self_recycle_auth_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 鉴权结果
	Data *OfnSelfRecycleAuthDto `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOfnSelfRecycleAuthAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = nil
}

var poolTaobaoOfnSelfRecycleAuthAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOfnSelfRecycleAuthAPIResponse)
	},
}

// GetTaobaoOfnSelfRecycleAuthAPIResponse 从 sync.Pool 获取 TaobaoOfnSelfRecycleAuthAPIResponse
func GetTaobaoOfnSelfRecycleAuthAPIResponse() *TaobaoOfnSelfRecycleAuthAPIResponse {
	return poolTaobaoOfnSelfRecycleAuthAPIResponse.Get().(*TaobaoOfnSelfRecycleAuthAPIResponse)
}

// ReleaseTaobaoOfnSelfRecycleAuthAPIResponse 将 TaobaoOfnSelfRecycleAuthAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOfnSelfRecycleAuthAPIResponse(v *TaobaoOfnSelfRecycleAuthAPIResponse) {
	v.Reset()
	poolTaobaoOfnSelfRecycleAuthAPIResponse.Put(v)
}
