package bus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusAgentCityChangeAPIResponse 城市变更 API返回值
// taobao.bus.agent.city.change
//
// 代理商通知城市变更，比如可售变为不可售等
type TaobaoBusAgentCityChangeAPIResponse struct {
	model.CommonResponse
	TaobaoBusAgentCityChangeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBusAgentCityChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBusAgentCityChangeAPIResponseModel).Reset()
}

// TaobaoBusAgentCityChangeAPIResponseModel is 城市变更 成功返回结果
type TaobaoBusAgentCityChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_agent_city_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统异常错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 异常描述
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBusAgentCityChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ErrCode = ""
	m.ErrMsg = ""
	m.IsSuccess = false
}

var poolTaobaoBusAgentCityChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBusAgentCityChangeAPIResponse)
	},
}

// GetTaobaoBusAgentCityChangeAPIResponse 从 sync.Pool 获取 TaobaoBusAgentCityChangeAPIResponse
func GetTaobaoBusAgentCityChangeAPIResponse() *TaobaoBusAgentCityChangeAPIResponse {
	return poolTaobaoBusAgentCityChangeAPIResponse.Get().(*TaobaoBusAgentCityChangeAPIResponse)
}

// ReleaseTaobaoBusAgentCityChangeAPIResponse 将 TaobaoBusAgentCityChangeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBusAgentCityChangeAPIResponse(v *TaobaoBusAgentCityChangeAPIResponse) {
	v.Reset()
	poolTaobaoBusAgentCityChangeAPIResponse.Put(v)
}
