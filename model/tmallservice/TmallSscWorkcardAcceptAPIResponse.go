package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSscWorkcardAcceptAPIResponse 服务商接单完成 API返回值
// tmall.ssc.workcard.accept
//
// 工单完结
type TmallSscWorkcardAcceptAPIResponse struct {
	model.CommonResponse
	TmallSscWorkcardAcceptAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSscWorkcardAcceptAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSscWorkcardAcceptAPIResponseModel).Reset()
}

// TmallSscWorkcardAcceptAPIResponseModel is 服务商接单完成 成功返回结果
type TmallSscWorkcardAcceptAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_ssc_workcard_accept_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接单结果
	Result *TmallSscWorkcardAcceptResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSscWorkcardAcceptAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSscWorkcardAcceptAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSscWorkcardAcceptAPIResponse)
	},
}

// GetTmallSscWorkcardAcceptAPIResponse 从 sync.Pool 获取 TmallSscWorkcardAcceptAPIResponse
func GetTmallSscWorkcardAcceptAPIResponse() *TmallSscWorkcardAcceptAPIResponse {
	return poolTmallSscWorkcardAcceptAPIResponse.Get().(*TmallSscWorkcardAcceptAPIResponse)
}

// ReleaseTmallSscWorkcardAcceptAPIResponse 将 TmallSscWorkcardAcceptAPIResponse 保存到 sync.Pool
func ReleaseTmallSscWorkcardAcceptAPIResponse(v *TmallSscWorkcardAcceptAPIResponse) {
	v.Reset()
	poolTmallSscWorkcardAcceptAPIResponse.Put(v)
}
