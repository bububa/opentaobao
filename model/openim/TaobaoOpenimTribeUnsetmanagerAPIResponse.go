package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeUnsetmanagerAPIResponse OPENIM群取消管理员 API返回值
// taobao.openim.tribe.unsetmanager
//
// OPENIM群取消管理员
type TaobaoOpenimTribeUnsetmanagerAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeUnsetmanagerAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeUnsetmanagerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeUnsetmanagerAPIResponseModel).Reset()
}

// TaobaoOpenimTribeUnsetmanagerAPIResponseModel is OPENIM群取消管理员 成功返回结果
type TaobaoOpenimTribeUnsetmanagerAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_unsetmanager_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeUnsetmanagerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeCode = 0
}

var poolTaobaoOpenimTribeUnsetmanagerAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeUnsetmanagerAPIResponse)
	},
}

// GetTaobaoOpenimTribeUnsetmanagerAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeUnsetmanagerAPIResponse
func GetTaobaoOpenimTribeUnsetmanagerAPIResponse() *TaobaoOpenimTribeUnsetmanagerAPIResponse {
	return poolTaobaoOpenimTribeUnsetmanagerAPIResponse.Get().(*TaobaoOpenimTribeUnsetmanagerAPIResponse)
}

// ReleaseTaobaoOpenimTribeUnsetmanagerAPIResponse 将 TaobaoOpenimTribeUnsetmanagerAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeUnsetmanagerAPIResponse(v *TaobaoOpenimTribeUnsetmanagerAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeUnsetmanagerAPIResponse.Put(v)
}
