package westcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/westcrm"
)

/* 
处理Usc会员消息接口 
alibaba.westcrm.member.message.handle

处理Usc会员消息接口
*/
func AlibabaWestcrmMemberMessageHandle(clt *core.SDKClient, req *westcrm.AlibabaWestcrmMemberMessageHandleAPIRequest, session string) (*westcrm.AlibabaWestcrmMemberMessageHandleAPIResponse, error) {
    var resp westcrm.AlibabaWestcrmMemberMessageHandleAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
