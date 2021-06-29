package lstlogistics2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-线下订单-导入接口 APIRequest
alibaba.lst.trade.seller.offline.order.upload

供应商线下订单数据上传、实现和零售通本地云仓订单的共配
*/
type AlibabaLstTradeSellerOfflineOrderUploadRequest struct {
    model.Params

    // 入参
    offlineOrderUploadParam   *LstOffLineOrderUploadParam 

}

func NewAlibabaLstTradeSellerOfflineOrderUploadRequest() *AlibabaLstTradeSellerOfflineOrderUploadRequest{
    return &AlibabaLstTradeSellerOfflineOrderUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstTradeSellerOfflineOrderUploadRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.offline.order.upload"
}

func (r AlibabaLstTradeSellerOfflineOrderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstTradeSellerOfflineOrderUploadRequest) SetOfflineOrderUploadParam(offlineOrderUploadParam *LstOffLineOrderUploadParam) error {
    r.offlineOrderUploadParam = offlineOrderUploadParam
    r.Set("offline_order_upload_param", offlineOrderUploadParam)
    return nil
}

func (r AlibabaLstTradeSellerOfflineOrderUploadRequest) GetOfflineOrderUploadParam() *LstOffLineOrderUploadParam {
    return r.offlineOrderUploadParam
}

