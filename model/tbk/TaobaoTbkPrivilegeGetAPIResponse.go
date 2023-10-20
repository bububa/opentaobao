package tbk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkPrivilegeGetAPIResponse 淘宝客-服务商-单品券高效转链 API返回值
// taobao.tbk.privilege.get
//
// 单品券高效转链API
type TaobaoTbkPrivilegeGetAPIResponse struct {
	model.CommonResponse
	TaobaoTbkPrivilegeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTbkPrivilegeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTbkPrivilegeGetAPIResponseModel).Reset()
}

// TaobaoTbkPrivilegeGetAPIResponseModel is 淘宝客-服务商-单品券高效转链 成功返回结果
type TaobaoTbkPrivilegeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tbk_privilege_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoTbkPrivilegeGetRpcResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTbkPrivilegeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoTbkPrivilegeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTbkPrivilegeGetAPIResponse)
	},
}

// GetTaobaoTbkPrivilegeGetAPIResponse 从 sync.Pool 获取 TaobaoTbkPrivilegeGetAPIResponse
func GetTaobaoTbkPrivilegeGetAPIResponse() *TaobaoTbkPrivilegeGetAPIResponse {
	return poolTaobaoTbkPrivilegeGetAPIResponse.Get().(*TaobaoTbkPrivilegeGetAPIResponse)
}

// ReleaseTaobaoTbkPrivilegeGetAPIResponse 将 TaobaoTbkPrivilegeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTbkPrivilegeGetAPIResponse(v *TaobaoTbkPrivilegeGetAPIResponse) {
	v.Reset()
	poolTaobaoTbkPrivilegeGetAPIResponse.Put(v)
}
