package tmallsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecentercallrecordqueryAPIResponse 天猫服务平台服务商查询通话记录接口 API返回值
// tmall.servicecenter.callrecord.query
//
// 天猫服务平台服务商查询通话记录接口
type TmallservicecentercallrecordqueryAPIResponse struct {
	model.CommonResponse
	TmallservicecentercallrecordqueryAPIResponseModel
}

// TmallservicecentercallrecordqueryAPIResponseModel is 天猫服务平台服务商查询通话记录接口 成功返回结果
type TmallservicecentercallrecordqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_servicecenter_callrecord_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallservicecentercallrecordqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
