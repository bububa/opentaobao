package tvpay

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tvpay"
)

/* 
tv支付查询消费抽奖配置 
taobao.tvpay.promotion.info.get

查询消费抽奖配置
*/
func TaobaoTvpayPromotionInfoGet(clt *core.SDKClient, req *tvpay.TaobaoTvpayPromotionInfoGetRequest, session string) (*tvpay.TaobaoTvpayPromotionInfoGetAPIResponse, error) {
    var resp tvpay.TaobaoTvpayPromotionInfoGetAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
