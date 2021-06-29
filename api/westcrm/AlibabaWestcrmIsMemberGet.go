package westcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/westcrm"
)

/* 
查询是否是亲橙里会员 
alibaba.westcrm.is.member.get

根据淘宝Id查询是否是亲橙里会员
*/
func AlibabaWestcrmIsMemberGet(clt *core.SDKClient, req *westcrm.AlibabaWestcrmIsMemberGetRequest, session string) (*westcrm.AlibabaWestcrmIsMemberGetAPIResponse, error) {
    var resp westcrm.AlibabaWestcrmIsMemberGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
