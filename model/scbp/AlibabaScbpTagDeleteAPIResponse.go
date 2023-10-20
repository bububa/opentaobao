package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTagDeleteAPIResponse 删除关键词分组 API返回值
// alibaba.scbp.tag.delete
//
// 删除关键词分组
type AlibabaScbpTagDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTagDeleteAPIResponseModel
}

// AlibabaScbpTagDeleteAPIResponseModel is 删除关键词分组 成功返回结果
type AlibabaScbpTagDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_tag_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除关键词分组成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
