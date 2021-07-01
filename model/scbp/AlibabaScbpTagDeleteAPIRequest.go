package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTagDeleteAPIRequest
删除关键词分组 API请求
alibaba.scbp.tag.delete

删除关键词分组 */
type AlibabaScbpTagDeleteAPIRequest struct {
	model.Params
	// 关键词分组名
	_tagName string
}

// New
