package eleenterpriserestaurant

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询餐厅信息 APIResponse
alibaba.ele.enterprise.restaurant.get

查询餐厅信息
*/
type AlibabaEleEnterpriseRestaurantGetAPIResponse struct {
    model.CommonResponse
    AlibabaEleEnterpriseRestaurantGetResponse
}

type AlibabaEleEnterpriseRestaurantGetResponse struct {
    XMLName xml.Name `xml:"alibaba_ele_enterprise_restaurant_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应code
    
    EnterpriseCode   string `json:"enterprise_code,omitempty" xml:"enterprise_code,omitempty"`

    
    // 响应信息
    
    EnterpriseMsg   string `json:"enterprise_msg,omitempty" xml:"enterprise_msg,omitempty"`

    
    // 返回餐厅信息
    
    EnterpriseData   *EnterpriseData `json:"enterprise_data,omitempty" xml:"enterprise_data,omitempty"`

    
    // 请求id
    
    EnterpriseRequestid   string `json:"enterprise_requestid,omitempty" xml:"enterprise_requestid,omitempty"`

    
}
