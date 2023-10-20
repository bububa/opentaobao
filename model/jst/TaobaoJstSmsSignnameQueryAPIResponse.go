package jst

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaojstsmssignnamequeryAPIResponse 淘宝短信签名查询 API返回值
// taobao.jst.sms.signname.query
//
// 淘宝短信签名查询
type TaobaojstsmssignnamequeryAPIResponse struct {
	model.CommonResponse
	TaobaojstsmssignnamequeryAPIResponseModel
}

// TaobaojstsmssignnamequeryAPIResponseModel is 淘宝短信签名查询 成功返回结果
type TaobaojstsmssignnamequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"jst_sms_signname_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	RCode string `json:"r_code,omitempty" xml:"r_code,omitempty"`
	// 失败原因
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 返回结果（signStatus： 0--待审核  1--通过 2--拒绝）
	Module *QuerySmsSignDto `json:"module,omitempty" xml:"module,omitempty"`
	// 请求成功
	RSuccess bool `json:"r_success,omitempty" xml:"r_success,omitempty"`
}
