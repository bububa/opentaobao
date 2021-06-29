package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自送发货 APIRequest
taobao.omniorder.dtd.consign

该接口触发门店自送发货，推进淘系订单状态为发货，为消费者发送核销码短信，并将物流信息写入订单
*/
type TaobaoOmniorderDtdConsignRequest struct {
    model.Params

    // 淘宝订单主订单号
    mainOrderId   int64 

    // 发货对应的商户中心门店ID
    storeId   int64 

}

func NewTaobaoOmniorderDtdConsignRequest() *TaobaoOmniorderDtdConsignRequest{
    return &TaobaoOmniorderDtdConsignRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderDtdConsignRequest) GetApiMethodName() string {
    return "taobao.omniorder.dtd.consign"
}

func (r TaobaoOmniorderDtdConsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderDtdConsignRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoOmniorderDtdConsignRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoOmniorderDtdConsignRequest) SetStoreId(storeId int64) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

func (r TaobaoOmniorderDtdConsignRequest) GetStoreId() int64 {
    return r.storeId
}

