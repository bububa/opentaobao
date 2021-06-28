package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
删除屏蔽词 
alibaba.scbp.ad.campaign.delete.forbidden.keyword

删除屏蔽词
*/
func AlibabaScbpAdCampaignDeleteForbiddenKeyword(clt *core.SDKClient, req *scbp.AlibabaScbpAdCampaignDeleteForbiddenKeywordRequest, session string) (*scbp.AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse, error) {
    var resp scbp.AlibabaScbpAdCampaignDeleteForbiddenKeywordAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
