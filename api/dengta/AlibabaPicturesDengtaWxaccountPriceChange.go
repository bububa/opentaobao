package dengta

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/dengta"
)

/* 
微信公众号价格变化通知 
alibaba.pictures.dengta.wxaccount.price.change

微信公众号推广价格变更通知接口
*/
func AlibabaPicturesDengtaWxaccountPriceChange(clt *core.SDKClient, req *dengta.AlibabaPicturesDengtaWxaccountPriceChangeRequest, session string) (*dengta.AlibabaPicturesDengtaWxaccountPriceChangeAPIResponse, error) {
    var resp dengta.AlibabaPicturesDengtaWxaccountPriceChangeAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
