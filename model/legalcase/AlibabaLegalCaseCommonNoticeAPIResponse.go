package legalcase

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasecommonnoticeAPIResponse 消息通知 API返回值
// alibaba.legal.case.common.notice
//
// 同步通知给诉讼系统
type AlibabalegalcasecommonnoticeAPIResponse struct {
	model.CommonResponse
	AlibabalegalcasecommonnoticeAPIResponseModel
}

// AlibabalegalcasecommonnoticeAPIResponseModel is 消息通知 成功返回结果
type AlibabalegalcasecommonnoticeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_legal_case_common_notice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// error
	Errcode string `json:"errcode,omitempty" xml:"errcode,omitempty"`
	// content
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// msg
	Errmsg string `json:"errmsg,omitempty" xml:"errmsg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
