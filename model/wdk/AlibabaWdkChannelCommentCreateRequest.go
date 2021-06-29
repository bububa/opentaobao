package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
差评导入 API请求
alibaba.wdk.channel.comment.create

差评导入
*/
type AlibabaWdkChannelCommentCreateRequest struct {
    model.Params
    // 差评信息
    commentCreateInfo   *CommentCreateInfo
}

// 初始化AlibabaWdkChannelCommentCreateRequest对象
func NewAlibabaWdkChannelCommentCreateRequest() *AlibabaWdkChannelCommentCreateRequest{
    return &AlibabaWdkChannelCommentCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelCommentCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.comment.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelCommentCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CommentCreateInfo Setter
// 差评信息
func (r *AlibabaWdkChannelCommentCreateRequest) SetCommentCreateInfo(commentCreateInfo *CommentCreateInfo) error {
    r.commentCreateInfo = commentCreateInfo
    r.Set("comment_create_info", commentCreateInfo)
    return nil
}

// CommentCreateInfo Getter
func (r AlibabaWdkChannelCommentCreateRequest) GetCommentCreateInfo() *CommentCreateInfo {
    return r.commentCreateInfo
}
