package uscesl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUsceslBizEslUnbindAPIResponse 电子价签解绑接口 API返回值
// taobao.uscesl.biz.esl.unbind
//
// 电子价签解绑接口
type TaobaoUsceslBizEslUnbindAPIResponse struct {
	model.CommonResponse
	TaobaoUsceslBizEslUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUsceslBizEslUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUsceslBizEslUnbindAPIResponseModel).Reset()
}

// TaobaoUsceslBizEslUnbindAPIResponseModel is 电子价签解绑接口 成功返回结果
type TaobaoUsceslBizEslUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"uscesl_biz_esl_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result.sucess表示本次调用是否成功，true或false
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUsceslBizEslUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolTaobaoUsceslBizEslUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUsceslBizEslUnbindAPIResponse)
	},
}

// GetTaobaoUsceslBizEslUnbindAPIResponse 从 sync.Pool 获取 TaobaoUsceslBizEslUnbindAPIResponse
func GetTaobaoUsceslBizEslUnbindAPIResponse() *TaobaoUsceslBizEslUnbindAPIResponse {
	return poolTaobaoUsceslBizEslUnbindAPIResponse.Get().(*TaobaoUsceslBizEslUnbindAPIResponse)
}

// ReleaseTaobaoUsceslBizEslUnbindAPIResponse 将 TaobaoUsceslBizEslUnbindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUsceslBizEslUnbindAPIResponse(v *TaobaoUsceslBizEslUnbindAPIResponse) {
	v.Reset()
	poolTaobaoUsceslBizEslUnbindAPIResponse.Put(v)
}
