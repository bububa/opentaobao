package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardSigninAPIResponse 服务商确认工人签到成功 API返回值
// tmall.servicecenter.workcard.signin
//
// 服务商确认工人签到成功。需要服务商自己保证工人是在现场服务中。否则虚假回传签到而引起的后续问题全部由服务商自己承担
type TmallServicecenterWorkcardSigninAPIResponse struct {
	model.CommonResponse
	TmallServicecenterWorkcardSigninAPIResponseModel
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardSigninAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallServicecenterWorkcardSigninAPIResponseModel).Reset()
}

// TmallServicecenterWorkcardSigninAPIResponseModel is 服务商确认工人签到成功 成功返回结果
type TmallServicecenterWorkcardSigninAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_signin_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// .
	Result *TmallServicecenterWorkcardSigninResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallServicecenterWorkcardSigninAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallServicecenterWorkcardSigninAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallServicecenterWorkcardSigninAPIResponse)
	},
}

// GetTmallServicecenterWorkcardSigninAPIResponse 从 sync.Pool 获取 TmallServicecenterWorkcardSigninAPIResponse
func GetTmallServicecenterWorkcardSigninAPIResponse() *TmallServicecenterWorkcardSigninAPIResponse {
	return poolTmallServicecenterWorkcardSigninAPIResponse.Get().(*TmallServicecenterWorkcardSigninAPIResponse)
}

// ReleaseTmallServicecenterWorkcardSigninAPIResponse 将 TmallServicecenterWorkcardSigninAPIResponse 保存到 sync.Pool
func ReleaseTmallServicecenterWorkcardSigninAPIResponse(v *TmallServicecenterWorkcardSigninAPIResponse) {
	v.Reset()
	poolTmallServicecenterWorkcardSigninAPIResponse.Put(v)
}
