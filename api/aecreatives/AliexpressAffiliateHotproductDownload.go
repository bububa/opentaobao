package aecreatives

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/aecreatives"
)

/* 
联盟营销爆品下载接口 
aliexpress.affiliate.hotproduct.download

查询联盟爆品API
*/
func AliexpressAffiliateHotproductDownload(clt *core.SDKClient, req *aecreatives.AliexpressAffiliateHotproductDownloadRequest, session string) (*aecreatives.AliexpressAffiliateHotproductDownloadAPIResponse, error) {
    var resp aecreatives.AliexpressAffiliateHotproductDownloadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
