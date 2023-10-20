package user

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappeleuserinfogetAPIResponse 获取饿了么用户信息 API返回值
// taobao.miniapp.eleuserinfo.get
//
// 获取饿了么用户信息
type TaobaominiappeleuserinfogetAPIResponse struct {
	model.CommonResponse
	TaobaominiappeleuserinfogetAPIResponseModel
}

// TaobaominiappeleuserinfogetAPIResponseModel is 获取饿了么用户信息 成功返回结果
type TaobaominiappeleuserinfogetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_eleuserinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// traceId
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 1
	Result *EleUicInfo `json:"result,omitempty" xml:"result,omitempty"`
}
