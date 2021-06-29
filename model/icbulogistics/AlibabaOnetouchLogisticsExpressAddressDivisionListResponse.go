package icbulogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-区域 API返回值 
alibaba.onetouch.logistics.express.address.division.list

四级地址库-区
*/
type AlibabaOnetouchLogisticsExpressAddressDivisionListAPIResponse struct {
    model.CommonResponse
    AlibabaOnetouchLogisticsExpressAddressDivisionListResponse
}

// 四级地址库-区域 成功返回结果
type AlibabaOnetouchLogisticsExpressAddressDivisionListResponse struct {
    XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_address_division_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaOnetouchLogisticsExpressAddressDivisionListResult `json:"result,omitempty" xml:"result,omitempty"`
}
