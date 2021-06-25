package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询存送产品信息 APIResponse
alibaba.alicom.wtt.opentrade.getproductinfo

话费宝查询产品信息相关配置
*/
type AlibabaAlicomWttOpentradeGetproductinfoAPIResponse struct {
    model.CommonResponse
    Response *AlibabaAlicomWttOpentradeGetproductinfoResponse `json:"alibaba_alicom_wtt_opentrade_getproductinfo_response,omitempty"`
}

type AlibabaAlicomWttOpentradeGetproductinfoResponse struct {

    // result
    Result   *TopResultDto `json:"result,omitempty"`

}
