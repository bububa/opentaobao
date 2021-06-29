package lstlogistics2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-线下订单-取消接口 APIResponse
alibaba.lst.trade.seller.offline.order.cancel

供应商线下订单数据上传之后取消
*/
type AlibabaLstTradeSellerOfflineOrderCancelAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeSellerOfflineOrderCancelResponse
}

type AlibabaLstTradeSellerOfflineOrderCancelResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_seller_offline_order_cancel_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaLstTradeSellerOfflineOrderCancelResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
