package lstlogistics2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-线下订单-查询接口 APIRequest
alibaba.lst.trade.seller.offline.order.query

供应商线下订单数据上传后查询物流状态
*/
type AlibabaLstTradeSellerOfflineOrderQueryRequest struct {
    model.Params

    // 入参
    offlineOrderQueryParam   *LstOfflineOrderQueryParam 

}

func NewAlibabaLstTradeSellerOfflineOrderQueryRequest() *AlibabaLstTradeSellerOfflineOrderQueryRequest{
    return &AlibabaLstTradeSellerOfflineOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeSellerOfflineOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.offline.order.query"
}

func (r AlibabaLstTradeSellerOfflineOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeSellerOfflineOrderQueryRequest) SetOfflineOrderQueryParam(offlineOrderQueryParam *LstOfflineOrderQueryParam) error {
    r.offlineOrderQueryParam = offlineOrderQueryParam
    r.Set("offline_order_query_param", offlineOrderQueryParam)
    return nil
}

func (r AlibabaLstTradeSellerOfflineOrderQueryRequest) GetOfflineOrderQueryParam() *LstOfflineOrderQueryParam {
    return r.offlineOrderQueryParam
}

