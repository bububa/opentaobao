package idleitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/idleitem"
)

/* 
图片上传 
alibaba.idle.item.media.add

上传图片
*/
func AlibabaIdleItemMediaAdd(clt *core.SDKClient, req *idleitem.AlibabaIdleItemMediaAddRequest, session string) (*idleitem.AlibabaIdleItemMediaAddAPIResponse, error) {
    var resp idleitem.AlibabaIdleItemMediaAddAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
