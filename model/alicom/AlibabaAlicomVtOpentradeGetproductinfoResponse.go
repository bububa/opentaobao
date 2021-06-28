package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询新虚拟产品配置信息 APIResponse
alibaba.alicom.vt.opentrade.getproductinfo

话费宝查询产品信息相关配置
*/
type AlibabaAlicomVtOpentradeGetproductinfoAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlicomVtOpentradeGetproductinfoResponse `json:"alibaba_alicom_vt_opentrade_getproductinfo_response,omitempty"` 
    AlibabaAlicomVtOpentradeGetproductinfoResponse
}

/* model for simplify = false
type AlibabaAlicomVtOpentradeGetproductinfoResponse struct {

    // result
    
    Result  *struct {
        TopResultDto  *TopResultDto `json:"top_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlicomVtOpentradeGetproductinfoResponse struct {

    // result
    Result   *TopResultDto `json:"result,omitempty"`

}
