package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询新虚拟产品配置信息 APIResponse
alibaba.alicom.vt.opentrade.getproductinfo

话费宝查询产品信息相关配置
*/
type AlibabaAlicomVtOpentradeGetproductinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlicomVtOpentradeGetproductinfoResponse
}

type AlibabaAlicomVtOpentradeGetproductinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alicom_vt_opentrade_getproductinfo_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *TopResultDto `json:"result,omitempty" xml:"result,omitempty"`

    
}
