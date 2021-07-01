package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTagRenameAPIRequest
重命名关键词分组 API请求
alibaba.scbp.tag.rename

重命名关键词分组 */
type AlibabaScbpTagRenameAPIRequest struct {
	model.Params
	// 需要重命名的关键词分组名
	_tagName string
	// 新分组名
	_newTagName string
}

// New
