package tmallsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardRefuseAPIResponse 买家拒收 API返回值
// tmall.servicecenter.workcard.refuse
//
// 买家拒收通知接口
type TmallServicecenterWorkcardRefuseAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardRefuseAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardRefuseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardRefuseAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardRefuseAPIResponseModel is 买家拒收 成功返回结果
type TmallServicecenterWorkcardRefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 返回code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否成功
	MsgSuccess bool `json:"msg_success,omitempty" xml:"msg_success,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardRefuseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.MsgSuccess = false
}

var poolTmallServicecenterWorkcardRefuseAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardRefuseAPIResponse)
	},
}

// GetTmallServicecenterWorkcardRefuseAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardRefuseAPIResponse
func GetTmallServicecenterWorkcardRefuseAPIResponse() *TmallServicecenterWorkcardRefuseAPIResponse {
	return poolTmallServicecenterWorkcardRefuseAPIResponse.Get().(*TmallServicecenterWorkcardRefuseAPIResponse)
}

// ReleaseTmallServicecenterWorkcardRefuseAPIResponse 将 TmallServicecenterWorkcardRefuseAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardRefuseAPIResponse(v *TmallServicecenterWorkcardRefuseAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardRefuseAPIResponse.Put(v)
}
