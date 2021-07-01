package icburfq

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/icburfq"
)

/* 
rfq推荐 
alibaba.icbu.rfq.recommend

rfq推荐
*/
func AlibabaIcbuRfqRecommend(clt *core.SDKClient, req *icburfq.AlibabaIcbuRfqRecommendAPIRequest, session string) (*icburfq.AlibabaIcbuRfqRecommendAPIResponse, error) {
    var resp icburfq.AlibabaIcbuRfqRecommendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
