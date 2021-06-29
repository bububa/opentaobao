package aecreatives

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aecreatives"
)

/* 
联盟主题推广活动信息获取 
aliexpress.affiliate.featuredpromo.get

获取联盟主题推广活动信息
*/
func AliexpressAffiliateFeaturedpromoGet(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateFeaturedpromoGetRequest, session string) (*aecreatives.AliexpressAffiliateFeaturedpromoGetAPIResponse, error) {
    var resp aecreatives.AliexpressAffiliateFeaturedpromoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
