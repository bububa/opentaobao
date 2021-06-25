package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
差评导入 APIRequest
alibaba.wdk.channel.comment.create

差评导入
*/
type AlibabaWdkChannelCommentCreateRequest struct {
    model.Params

    // 差评信息
    commentCreateInfo   *CommentCreateInfo 

}

func NewAlibabaWdkChannelCommentCreateRequest() *AlibabaWdkChannelCommentCreateRequest{
    return &AlibabaWdkChannelCommentCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkChannelCommentCreateRequest) GetApiMethodName() string {
    return "alibaba.wdk.channel.comment.create"
}

func (r AlibabaWdkChannelCommentCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkChannelCommentCreateRequest) SetCommentCreateInfo(commentCreateInfo *CommentCreateInfo) error {
    r.commentCreateInfo = commentCreateInfo
    r.Set("comment_create_info", commentCreateInfo)
    return nil
}

func (r AlibabaWdkChannelCommentCreateRequest) GetCommentCreateInfo() *CommentCreateInfo {
    return r.commentCreateInfo
}

