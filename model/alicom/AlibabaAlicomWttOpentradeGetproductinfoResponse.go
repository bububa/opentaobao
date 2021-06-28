package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询存送产品信息 APIResponse
alibaba.alicom.wtt.opentrade.getproductinfo

话费宝查询产品信息相关配置
*/
type AlibabaAlicomWttOpentradeGetproductinfoAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alicom_wtt_opentrade_getproductinfo_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TopResultDto `json:"result,omitempty" xml:"