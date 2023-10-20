package alink

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliyunalinkopendataurlqueryAPIResponse 开放数据授权访问URL查询 API返回值
// aliyun.alink.opendata.url.query
//
// 厂商数据授权访问URL查询
type AliyunalinkopendataurlqueryAPIResponse struct {
	model.CommonResponse
	AliyunalinkopendataurlqueryAPIResponseModel
}

// AliyunalinkopendataurlqueryAPIResponseModel is 开放数据授权访问URL查询 成功返回结果
type AliyunalinkopendataurlqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_alink_opendata_url_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 授权url
	Module []string `json:"module,omitempty" xml:"module>string,omitempty"`
	// 接口描述
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
