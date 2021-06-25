package scbp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/scbp"
)

/* 
查询标签数据 
alibaba.scbp.ad.target.tag.find.campaign.target.tag

查询标签数据
*/
func AlibabaScbpAdTargetTagFindCampaignTargetTag(clt *core.SDKClient, req *scbp.AlibabaScbpAdTargetTagFindCampaignTargetTagRequest, session string) (*scbp.AlibabaScbpAdTargetTagFindCampaignTargetTagResponse, error) {
    var resp scbp.AlibabaScbpAdTargetTagFindCampaignTargetTagAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
