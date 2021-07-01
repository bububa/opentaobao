package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkChannelCommentCreateAPIRequest
差评导入 API请求
alibaba.wdk.channel.comment.create

差评导入 */
type AlibabaWdkChannelCommentCreateAPIRequest struct {
	model.Params
	// 差评信息
	_commentCreateInfo *CommentCreateInfo
}

// New
