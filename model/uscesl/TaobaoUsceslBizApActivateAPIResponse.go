package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizApActivateAPIResponse 激活AP价签通讯模块 API返回值
// taobao.uscesl.biz.ap.activate
//
// 激活AP价签通讯模块
type TaobaoUsceslBizApActivateAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizApActivateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizApActivateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizApActivateAPIResponseModel).Reset()
}

// TaobaoUsceslBizApActivateAPIResponseModel is 激活AP价签通讯模块 成功返回结果
type TaobaoUsceslBizApActivateAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_ap_activate_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功与否看result.success，返回true或者false
	Result *TaobaoUsceslBizApActivateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizApActivateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUsceslBizApActivateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizApActivateAPIResponse)
	},
}

// GetTaobaoUsceslBizApActivateAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizApActivateAPIResponse
func GetTaobaoUsceslBizApActivateAPIResponse() *TaobaoUsceslBizApActivateAPIResponse {
	return poolTaobaoUsceslBizApActivateAPIResponse.Get().(*TaobaoUsceslBizApActivateAPIResponse)
}

// ReleaseTaobaoUsceslBizApActivateAPIResponse 将 TaobaoUsceslBizApActivateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizApActivateAPIResponse(v *TaobaoUsceslBizApActivateAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizApActivateAPIResponse.Put(v)
}
