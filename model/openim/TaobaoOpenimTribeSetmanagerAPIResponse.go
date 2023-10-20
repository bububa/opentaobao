package openim

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenimTribeSetmanagerAPIResponse OPENIM群设置管理员 API返回值
// taobao.openim.tribe.setmanager
//
// OPENIM群设置管理员
type TaobaoOpenimTribeSetmanagerAPIResponse struct {
	model.CommonResponse
	TaobaoOpenimTribeSetmanagerAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeSetmanagerAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenimTribeSetmanagerAPIResponseModel).Reset()
}

// TaobaoOpenimTribeSetmanagerAPIResponseModel is OPENIM群设置管理员 成功返回结果
type TaobaoOpenimTribeSetmanagerAPIResponseModel struct {
	XMLName xml.Name `xml:"openim_tribe_setmanager_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 群服务code
	TribeCode int64 `json:"tribe_code,omitempty" xml:"tribe_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenimTribeSetmanagerAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TribeCode = 0
}

var poolTaobaoOpenimTribeSetmanagerAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenimTribeSetmanagerAPIResponse)
	},
}

// GetTaobaoOpenimTribeSetmanagerAPIResponse 从 sync.Pool 获取 TaobaoOpenimTribeSetmanagerAPIResponse
func GetTaobaoOpenimTribeSetmanagerAPIResponse() *TaobaoOpenimTribeSetmanagerAPIResponse {
	return poolTaobaoOpenimTribeSetmanagerAPIResponse.Get().(*TaobaoOpenimTribeSetmanagerAPIResponse)
}

// ReleaseTaobaoOpenimTribeSetmanagerAPIResponse 将 TaobaoOpenimTribeSetmanagerAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenimTribeSetmanagerAPIResponse(v *TaobaoOpenimTribeSetmanagerAPIResponse) {
	v.Reset()
	poolTaobaoOpenimTribeSetmanagerAPIResponse.Put(v)
}
