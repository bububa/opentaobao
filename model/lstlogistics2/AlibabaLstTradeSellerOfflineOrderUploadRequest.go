package lstlogistics2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-线下订单-导入接口 API请求
alibaba.lst.trade.seller.offline.order.upload

供应商线下订单数据上传、实现和零售通本地云仓订单的共配
*/
type AlibabaLstTradeSellerOfflineOrderUploadRequest struct {
    model.Params
    // 入参
    offlineOrderUploadParam   *LstOffLineOrderUploadParam
}

// 初始化AlibabaLstTradeSellerOfflineOrderUploadRequest对象
func NewAlibabaLstTradeSellerOfflineOrderUploadRequest() *AlibabaLstTradeSellerOfflineOrderUploadRequest{
    return &AlibabaLstTradeSellerOfflineOrderUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstTradeSellerOfflineOrderUploadRequest) GetApiMethodName() string {
    return "alibaba.lst.trade.seller.offline.order.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstTradeSellerOfflineOrderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineOrderUploadParam Setter
// 入参
func (r *AlibabaLstTradeSellerOfflineOrderUploadRequest) SetOfflineOrderUploadParam(offlineOrderUploadParam *LstOffLineOrderUploadParam) error {
    r.offlineOrderUploadParam = offlineOrderUploadParam
    r.Set("offline_order_upload_param", offlineOrderUploadParam)
    return nil
}

// OfflineOrderUploadParam Getter
func (r AlibabaLstTradeSellerOfflineOrderUploadRequest) GetOfflineOrderUploadParam() *LstOffLineOrderUploadParam {
    return r.offlineOrderUploadParam
}
