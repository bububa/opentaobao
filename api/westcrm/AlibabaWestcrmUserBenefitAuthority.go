package westcrm

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/westcrm"
)

/* 
获取指定用户是否含有会员权益配置菜单权限 
alibaba.westcrm.user.benefit.authority

获取指定用户是否含有会员权益配置菜单权限
*/
func AlibabaWestcrmUserBenefitAuthority(clt *core.SDKClient, req *westcrm.AlibabaWestcrmUserBenefitAuthorityRequest, session string) (*westcrm.AlibabaWestcrmUserBenefitAuthorityAPIResponse, error) {
    var resp westcrm.AlibabaWestcrmUserBenefitAuthorityAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
