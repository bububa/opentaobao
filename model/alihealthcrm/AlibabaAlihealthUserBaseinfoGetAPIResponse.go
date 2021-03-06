package alihealthcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthUserBaseinfoGetAPIResponse 获取用户基础信息 API返回值
// alibaba.alihealth.user.baseinfo.get
//
// 获取用户基础信息
type AlibabaAlihealthUserBaseinfoGetAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthUserBaseinfoGetAPIResponseModel
}

// AlibabaAlihealthUserBaseinfoGetAPIResponseModel is 获取用户基础信息 成功返回结果
type AlibabaAlihealthUserBaseinfoGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_user_baseinfo_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
