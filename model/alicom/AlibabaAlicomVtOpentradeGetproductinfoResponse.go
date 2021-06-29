package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询新虚拟产品配置信息 API返回值 
alibaba.alicom.vt.opentrade.getproductinfo

话费宝查询产品信息相关配置
*/
type AlibabaAlicomVtOpentradeGetproductinfoAPIResponse struct {
    model.CommonResponse
    AlibabaAlicomVtOpentradeGetproductinfoResponse
}

// 查询新虚拟产品配置信息 成功返回结果
type AlibabaAlicomVtOpentradeGetproductinfoResponse struct {
    XMLName xml.Name `xml:"alibaba_alicom_vt_opentrade_getproductinfo_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TopResultDTO `json:"result,omitempty" xml:"result,omitempty"`
}
