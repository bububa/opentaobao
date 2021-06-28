package trade

import (
    "github.com/bububa/opentaobao/model"
)

/* 
交易信息回流 APIResponse
alibaba.lst.vas.tradeflow.save

自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
*/
type AlibabaLstVasTradeflowSaveAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaLstVasTradeflowSaveResponse `json:"alibaba_lst_vas_tradeflow_save_response,omitempty"` 
    AlibabaLstVasTradeflowSaveResponse
}

/* model for simplify = false
type AlibabaLstVasTradeflowSaveResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaLstVasTradeflowSaveResult  *AlibabaLstVasTradeflowSaveResult `json:"alibaba_lst_vas_tradeflow_save_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaLstVasTradeflowSaveResponse struct {

    // 接口返回model
    Result   *AlibabaLstVasTradeflowSaveResult `json:"result,omitempty"`

}
