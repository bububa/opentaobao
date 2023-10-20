package mei

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMeiCrmCallbackPointChangeAPIResponse 品牌积分变更回调API API返回值
// tmall.mei.crm.callback.point.change
//
// 线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。
type TmallMeiCrmCallbackPointChangeAPIResponse struct {
	model.CommonResponse
	TmallMeiCrmCallbackPointChangeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallMeiCrmCallbackPointChangeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallMeiCrmCallbackPointChangeAPIResponseModel).Reset()
}

// TmallMeiCrmCallbackPointChangeAPIResponseModel is 品牌积分变更回调API 成功返回结果
type TmallMeiCrmCallbackPointChangeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mei_crm_callback_point_change_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallMeiCrmCallbackPointChangeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultSuccess = false
}

var poolTmallMeiCrmCallbackPointChangeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallMeiCrmCallbackPointChangeAPIResponse)
	},
}

// GetTmallMeiCrmCallbackPointChangeAPIResponse 从 sync.Pool 获取 TmallMeiCrmCallbackPointChangeAPIResponse
func GetTmallMeiCrmCallbackPointChangeAPIResponse() *TmallMeiCrmCallbackPointChangeAPIResponse {
	return poolTmallMeiCrmCallbackPointChangeAPIResponse.Get().(*TmallMeiCrmCallbackPointChangeAPIResponse)
}

// ReleaseTmallMeiCrmCallbackPointChangeAPIResponse 将 TmallMeiCrmCallbackPointChangeAPIResponse 保存到 sync.Pool
func ReleaseTmallMeiCrmCallbackPointChangeAPIResponse(v *TmallMeiCrmCallbackPointChangeAPIResponse) {
	v.Reset()
	poolTmallMeiCrmCallbackPointChangeAPIResponse.Put(v)
}
