package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
运营商获得用户身份信息 APIRequest
alibaba.aliqin.tcc.trade.identity.get

天猫网厅运营商官方旗舰店获取用户身份信息
*/
type AlibabaAliqinTccTradeIdentityGetRequest struct {
    model.Params

    // 订单编号
    bizOrderId   int64 

    // 店铺名称
    sellerNick   string 

}

func NewAlibabaAliqinTccTradeIdentityGetRequest() *AlibabaAliqinTccTradeIdentityGetRequest{
    return &AlibabaAliqinTccTradeIdentityGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinTccTradeIdentityGetRequest) GetApiMethodName() string {
    return "alibaba.aliqin.tcc.trade.identity.get"
}

func (r AlibabaAliqinTccTradeIdentityGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinTccTradeIdentityGetRequest) SetBizOrderId(bizOrderId int64) error {
    r.bizOrderId = bizOrderId
    r.Set("biz_order_id", bizOrderId)
    return nil
}

func (r AlibabaAliqinTccTradeIdentityGetRequest) GetBizOrderId() int64 {
    return r.bizOrderId
}

func (r *AlibabaAliqinTccTradeIdentityGetRequest) SetSellerNick(sellerNick string) error {
    r.sellerNick = sellerNick
    r.Set("seller_nick", sellerNick)
    return nil
}

func (r AlibabaAliqinTccTradeIdentityGetRequest) GetSellerNick() string {
    return r.sellerNick
}

