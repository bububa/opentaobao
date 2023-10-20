package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabacharitycharitytimelistAPIResponse 授权用户的公益时记录查询 API返回值
// alibaba.charity.charitytime.list
//
// 查询授权用户的公益时记录
type AlibabacharitycharitytimelistAPIResponse struct {
	model.CommonResponse
	AlibabacharitycharitytimelistAPIResponseModel
}

// AlibabacharitycharitytimelistAPIResponseModel is 授权用户的公益时记录查询 成功返回结果
type AlibabacharitycharitytimelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 公益时记录列表
	Data []CharityTimeDto `json:"data,omitempty" xml:"data>charity_time_dto,omitempty"`
	// 错误内容
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回码 200成功 403无权限
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
