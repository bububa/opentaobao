package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询屏蔽词 
alibaba.scbp.ad.campaign.find.forbidden.keyword

查询屏蔽词
*/
func AlibabaScbpAdCampaignFindForbiddenKeyword(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignFindForbiddenKeywordRequest, session string) (*scbp.AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse, error) {
    var resp scbp.AlibabaScbpAdCampaignFindForbiddenKeywordAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
