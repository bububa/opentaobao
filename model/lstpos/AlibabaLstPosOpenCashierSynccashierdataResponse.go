package lstpos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
收银快照同步接口(最多10条订单信息) APIResponse
alibaba.lst.pos.open.cashier.synccashierdata

收银快照同步接口(最多10条订单信息)
*/
type AlibabaLstPosOpenCashierSynccashierdataAPIResponse struct {
    model.CommonResponse
    AlibabaLstPosOpenCashierSynccashierdataResponse
}

type AlibabaLstPosOpenCashierSynccashierdataResponse struct {
    XMLName xml.Name `xml:"alibaba_lst_pos_open_cashier_synccashierdata_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应结果
    
    Result   *AlibabaLstPosOpenCashierSynccashierdataResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
