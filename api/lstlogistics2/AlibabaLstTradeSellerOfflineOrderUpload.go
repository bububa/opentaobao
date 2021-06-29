package lstlogistics2

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/lstlogistics2"
)

/* 
供应商-线下订单-导入接口 
alibaba.lst.trade.seller.offline.order.upload

供应商线下订单数据上传、实现和零售通本地云仓订单的共配
*/
func AlibabaLstTradeSellerOfflineOrderUpload(clt *core.SDKClient, req *lstlogistics2.AlibabaLstTradeSellerOfflineOrderUploadRequest, session string) (*lstlogistics2.AlibabaLstTradeSellerOfflineOrderUploadAPIResponse, error) {
    var resp lstlogistics2.AlibabaLstTradeSellerOfflineOrderUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
