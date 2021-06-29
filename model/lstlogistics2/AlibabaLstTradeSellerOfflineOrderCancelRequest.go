package lstlogistics2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-线下订单-取消接口 APIRequest
alibaba.lst.trade.seller.offline.order.cancel

供应商线下订单数据上传之后取消
*/
type AlibabaLstTradeSellerOfflineOrderCancelRequest struct {
    model.Params

    // 入参
    offlineOrderCancalParam   *LstOfflineOrderCancalParam 

}

func NewAlibabaLstTradeSellerOfflineOrderCancelRequest() *AlibabaLstTradeSellerOfflineOrderCancelRequest{
    return &AlibabaLstTradeSellerOfflineOrderCancelRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeSellerOfflineOrderCancelRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.offline.order.cancel"
}

func (r AlibabaLstTradeSellerOfflineOrderCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeSellerOfflineOrderCancelRequest) SetOfflineOrderCancalParam(offlineOrderCancalParam *LstOfflineOrderCancalParam) error {
    r.offlineOrderCancalParam = offlineOrderCancalParam
    r.Set("offline_order_cancal_param", offlineOrderCancalParam)
    return nil
}

func (r AlibabaLstTradeSellerOfflineOrderCancelRequest) GetOfflineOrderCancalParam() *LstOfflineOrderCancalParam {
    return r.offlineOrderCancalParam
}

