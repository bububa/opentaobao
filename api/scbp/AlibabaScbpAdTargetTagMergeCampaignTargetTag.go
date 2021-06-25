package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
标签增删改 
alibaba.scbp.ad.target.tag.merge.campaign.target.tag

标签增删改
*/
func AlibabaScbpAdTargetTagMergeCampaignTargetTag(clt *core.SDKClient, req *scbp.AlibabaScbpAdTargetTagMergeCampaignTargetTagRequest, session string) (*scbp.AlibabaScbpAdTargetTagMergeCampaignTargetTagResponse, error) {
    var resp scbp.AlibabaScbpAdTargetTagMergeCampaignTargetTagAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
