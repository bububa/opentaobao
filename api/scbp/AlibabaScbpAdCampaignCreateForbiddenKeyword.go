package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
创建屏蔽词 
alibaba.scbp.ad.campaign.create.forbidden.keyword

创建屏蔽词
*/
func AlibabaScbpAdCampaignCreateForbiddenKeyword(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignCreateForbiddenKeywordRequest, session string) (*scbp.AlibabaScbpAdCampaignCreateForbiddenKeywordResponse, error) {
    var resp scbp.AlibabaScbpAdCampaignCreateForbiddenKeywordAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
