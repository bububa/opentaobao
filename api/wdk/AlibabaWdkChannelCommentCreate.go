package wdk

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/wdk"
)

/* 
差评导入 
alibaba.wdk.channel.comment.create

差评导入
*/
func AlibabaWdkChannelCommentCreate(clt *core.SDKClient, req *wdk.AlibabaWdkChannelCommentCreateAPIRequest, session string) (*wdk.AlibabaWdkChannelCommentCreateAPIResponse, error) {
    var resp wdk.AlibabaWdkChannelCommentCreateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
