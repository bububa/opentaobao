package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptaglistAPIResponse 查询所有分组 API返回值
// alibaba.scbp.tag.list
//
// 查询所有分组
type AlibabascbptaglistAPIResponse struct {
	model.CommonResponse
	AlibabascbptaglistAPIResponseModel
}

// AlibabascbptaglistAPIResponseModel is 查询所有分组 成功返回结果
type AlibabascbptaglistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_tag_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 所有分组
	TagList []TagGroup `json:"tag_list,omitempty" xml:"tag_list>tag_group,omitempty"`
}
