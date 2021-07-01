package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTagRenameAPIResponse
重命名关键词分组 API返回值
alibaba.scbp.tag.rename

重命名关键词分组 */
type AlibabaScbpTagRenameAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTagRenameAPIResponseModel
}

// AlibabaScbpTagRenameAPIResponseModel is 重命名关键词分组 成功返回结果
type AlibabaScbpTagRenameAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_tag_rename_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 重命名分组成功或者失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
