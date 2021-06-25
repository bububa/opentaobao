package alink

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alink"
)

/* 
消息提醒开关 
alibaba.alink.message.config.set

阿里智能消息开关
*/
func AlibabaAlinkMessageConfigSet(clt *core.SDKClient, req *alink.AlibabaAlinkMessageConfigSetRequest, session string) (*alink.AlibabaAlinkMessageConfigSetResponse, error) {
    var resp alink.AlibabaAlinkMessageConfigSetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
