package eleenterpriserestaurant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
检查地址是否在餐厅配送范围内 APIResponse
alibaba.ele.enterprise.restaurant.checkaddress

检查地址是否在餐厅配送范围内
*/
type AlibabaEleEnterpriseRestaurantCheckaddressAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseRestaurantCheckaddressResponse
}

type AlibabaEleEnterpriseRestaurantCheckaddressResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_restaurant_checkaddress_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回值信息
    
    EnterpriseDatas   []EnterpriseData `json:"enterprise_datas,omitempty" xml:"enterprise_datas>enterprise_data,omitempty"`
    
    
    // 响应code
    
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`

    
    // 响应信息
    
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`

    
    // 请求id
    
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`

    
}
