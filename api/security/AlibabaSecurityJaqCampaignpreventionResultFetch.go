package security

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/security"
)

/* 
获取活动保护结果 
alibaba.security.jaq.campaignprevention.result.fetch

获取活动保护结果
*/
func AlibabaSecurityJaqCampaignpreventionResultFetch(clt *core.SDKClient, req *security.AlibabaSecurityJaqCampaignpreventionResultFetchRequest, session string) (*security.AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse, error) {
    var resp security.AlibabaSecurityJaqCampaignpreventionResultFetchAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
