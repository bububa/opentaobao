package lstlogistics2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-线下订单-导入接口 APIResponse
alibaba.lst.trade.seller.offline.order.upload

供应商线下订单数据上传、实现和零售通本地云仓订单的共配
*/
type AlibabaLstTradeSellerOfflineOrderUploadAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeSellerOfflineOrderUploadResponse
}

type AlibabaLstTradeSellerOfflineOrderUploadResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_seller_offline_order_upload_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaLstTradeSellerOfflineOrderUploadResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
