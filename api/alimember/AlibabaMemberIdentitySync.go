package alimember

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/alimember"
)

/* 
会员身份信息同步 
alibaba.member.identity.sync

会员身份信息同步
*/
func AlibabaMemberIdentitySync(clt *core.SDKClient, req *alimember.AlibabaMemberIdentitySyncAPIRequest, session string) (*alimember.AlibabaMemberIdentitySyncAPIResponse, error) {
    var resp alimember.AlibabaMemberIdentitySyncAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
