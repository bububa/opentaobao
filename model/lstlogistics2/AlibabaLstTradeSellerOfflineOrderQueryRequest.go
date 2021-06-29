package lstlogistics2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-线下订单-查询接口 API请求
alibaba.lst.trade.seller.offline.order.query

供应商线下订单数据上传后查询物流状态
*/
type AlibabaLstTradeSellerOfflineOrderQueryRequest struct {
    model.Params
    // 入参
    _offlineOrderQueryParam   *LstOfflineOrderQueryParam
}

// 初始化AlibabaLstTradeSellerOfflineOrderQueryRequest对象
func NewAlibabaLstTradeSellerOfflineOrderQueryRequest() *AlibabaLstTradeSellerOfflineOrderQueryRequest{
    return &AlibabaLstTradeSellerOfflineOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOfflineOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.offline.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOfflineOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineOrderQueryParam Setter
// 入参
func (r *AlibabaLstTradeSellerOfflineOrderQueryRequest) SetOfflineOrderQueryParam(_offlineOrderQueryParam *LstOfflineOrderQueryParam) error {
    r._offlineOrderQueryParam = _offlineOrderQueryParam
    r.Set("offline_order_query_param", _offlineOrderQueryParam)
    return nil
}

// OfflineOrderQueryParam Getter
func (r AlibabaLstTradeSellerOfflineOrderQueryRequest) GetOfflineOrderQueryParam() *LstOfflineOrderQueryParam {
    return r._offlineOrderQueryParam
}
