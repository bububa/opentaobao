package icbulogistics

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
四级地址库-市 APIResponse
alibaba.onetouch.logistics.express.address.city.list

四级地址库-市
*/
type AlibabaOnetouchLogisticsExpressAddressCityListAPIResponse struct {
    model.CommonResponse
    AlibabaOnetouchLogisticsExpressAddressCityListResponse
}

type AlibabaOnetouchLogisticsExpressAddressCityListResponse struct {
    XMLName xml.Name `xml:"alibaba_onetouch_logistics_express_address_city_list_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *AlibabaOnetouchLogisticsExpressAddressCityListResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
