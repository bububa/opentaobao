package tmallservice

import (
	"encoding/xml"

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

// TmallServicecenterWorkcardSigninAPIResponseModel is 服务商确认工人签到成功 成功返回结果
type TmallServicecenterWorkcardSigninAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_workcard_signin_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// .
	Result *TmallServicecenterWorkcardSigninResult `json:"result,omitempty" xml:"result,omitempty"`
}
