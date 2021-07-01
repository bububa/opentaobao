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

// NewAlibabaWdkChannelCommentCreateRequest 初始化AlibabaWdkChannelCommentCreateAPIRequest对象
func NewAlibabaWdkChannelCommentCreateRequest() *AlibabaWdkChannelCommentCreateAPIRequest {
	return &AlibabaWdkChannelCommentCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelCommentCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.comment.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelCommentCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CommentCreateInfo Setter
// 差评信息
func (r *AlibabaWdkChannelCommentCreateAPIRequest) SetCommentCreateInfo(_commentCreateInfo *CommentCreateInfo) error {
	r._commentCreateInfo = _commentCreateInfo
	r.Set("comment_create_info", _commentCreateInfo)
	return nil
}

// Get CommentCreateInfo Getter
func (r AlibabaWdkChannelCommentCreateAPIRequest) GetCommentCreateInfo() *CommentCreateInfo {
	return r._commentCreateInfo
}
