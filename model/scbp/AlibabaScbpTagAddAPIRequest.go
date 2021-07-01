package scbp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpTagAddAPIRequest
创建关键词分组 API请求
alibaba.scbp.tag.add

创建关键词分组 */
type AlibabaScbpTagAddAPIRequest struct {
	model.Params
	// 分组名称，最多允许创建100个
	_tagName string
}

// New
