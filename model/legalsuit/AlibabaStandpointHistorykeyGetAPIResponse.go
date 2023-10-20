package legalsuit

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabastandpointhistorykeygetAPIResponse 查询历史数据 API返回值
// alibaba.standpoint.historykey.get
//
// 查询历史数据
type AlibabastandpointhistorykeygetAPIResponse struct {
	model.CommonResponse
	AlibabastandpointhistorykeygetAPIResponseModel
}

// AlibabastandpointhistorykeygetAPIResponseModel is 查询历史数据 成功返回结果
type AlibabastandpointhistorykeygetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_standpoint_historykey_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 数据列表
	Content []string `json:"content,omitempty" xml:"content>string,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 状态吗
	ErrorCodeRes int64 `json:"error_code_res,omitempty" xml:"error_code_res,omitempty"`
	// 是否成功
	SuccessRes bool `json:"success_res,omitempty" xml:"success_res,omitempty"`
}
