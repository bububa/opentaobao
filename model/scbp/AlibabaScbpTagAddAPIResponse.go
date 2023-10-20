package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptagaddAPIResponse 创建关键词分组 API返回值
// alibaba.scbp.tag.add
//
// 创建关键词分组
type AlibabascbptagaddAPIResponse struct {
	model.CommonResponse
	AlibabascbptagaddAPIResponseModel
}

// AlibabascbptagaddAPIResponseModel is 创建关键词分组 成功返回结果
type AlibabascbptagaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_tag_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分组名称
	TagName string `json:"tag_name,omitempty" xml:"tag_name,omitempty"`
}
