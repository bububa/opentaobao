package lsttrade

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家退款单商品状态同步 APIResponse
alibaba.lst.trade.fastrefund.goodsstatus.sync

卖家退款单商品状态同步
*/
type AlibabaLstTradeFastrefundGoodsstatusSyncAPIResponse struct {
    model.CommonResponse
    AlibabaLstTradeFastrefundGoodsstatusSyncResponse
}

type AlibabaLstTradeFastrefundGoodsstatusSyncResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_trade_fastrefund_goodsstatus_sync_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // true表示成功，false表示失败
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
}
