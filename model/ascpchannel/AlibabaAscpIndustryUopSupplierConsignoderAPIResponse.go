package ascpchannel

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家推单 API返回值 
alibaba.ascp.industry.uop.supplier.consignoder

商家推单
*/
type AlibabaAscpIndustryUopSupplierConsignoderAPIResponse struct {
    model.CommonResponse
    AlibabaAscpIndustryUopSupplierConsignoderAPIResponseModel
}

// 商家推单 成功返回结果
type AlibabaAscpIndustryUopSupplierConsignoderAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_ascp_industry_uop_supplier_consignoder_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 商家推送天猫信息后，由天猫回传的字段
    Data   *AlibabaAscpIndustryUopSupplierConsignoderData `json:"data,omitempty" xml:"data,omitempty"`
}
