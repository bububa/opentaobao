package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
五道口餐饮取餐号获取接口 APIResponse
wdk.meal.pos.getfetchmealcode

pos机创建订单前获取餐饮取餐号
*/
type WdkMealPosGetfetchmealcodeAPIResponse struct {
    model.CommonResponse
    WdkMealPosGetfetchmealcodeResponse
}

type WdkMealPosGetfetchmealcodeResponse struct {
    XMLName xml.Name `xml:"wdk_meal_pos_getfetchmealcode_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 取餐号
    
    FetchMealCode   string `json:"fetch_meal_code,omitempty" xml:"fetch_meal_code,omitempty"`

    
}
