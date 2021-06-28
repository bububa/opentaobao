package mei

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/mei"
)

/* 
支付码获取会员信息 
tmall.mei.crm.member.getbypaycode

通过支付码获取会员信息
*/
func TmallMeiCrmMemberGetbypaycode(clt *core.SDKClient, req *mei.TmallMeiCrmMemberGetbypaycodeRequest, session string) (*mei.TmallMeiCrmMemberGetbypaycodeAPIResponse, error) {
    var resp mei.TmallMeiCrmMemberGetbypaycodeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
