package kclub

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabakclubkcqagetAPIResponse 知识云-查询单个知识详情 API返回值
// alibaba.kclub.kc.qa.get
//
// 知识云-查询单个知识详情。通过租户id、问题id查询问题详情
type AlibabakclubkcqagetAPIResponse struct {
	model.CommonResponse
	AlibabakclubkcqagetAPIResponseModel
}

// AlibabakclubkcqagetAPIResponseModel is 知识云-查询单个知识详情 成功返回结果
type AlibabakclubkcqagetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_kclub_kc_qa_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabakclubkcqagetResult `json:"result,omitempty" xml:"result,omitempty"`
}
