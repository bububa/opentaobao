package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaascpcnskumodifyAPIResponse 供应链中台货品修改接口 API返回值
// alibaba.ascp.cnsku.modify
//
// 供应链中台货品修改接口
type AlibabaascpcnskumodifyAPIResponse struct {
	model.CommonResponse
	AlibabaascpcnskumodifyAPIResponseModel
}

// AlibabaascpcnskumodifyAPIResponseModel is 供应链中台货品修改接口 成功返回结果
type AlibabaascpcnskumodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_cnsku_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异常信息
	ErrorMessages []string `json:"error_messages,omitempty" xml:"error_messages>string,omitempty"`
	// Null
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 异常信息Code
	SysErrorCode string `json:"sys_error_code,omitempty" xml:"sys_error_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 是否系统异常
	IsSystemFailed bool `json:"is_system_failed,omitempty" xml:"is_system_failed,omitempty"`
}
