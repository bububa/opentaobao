package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
五道口餐饮取餐号获取接口 APIResponse
wdk.meal.pos.getfetchmealcode

pos机创建订单前获取餐饮取餐号
*/
type WdkMealPosGetfetchmealcodeAPIResponse struct {
    model.CommonResponse
    Response *WdkMealPosGetfetchmealcodeResponse `json:"wdk_meal_pos_getfetchmealcode_response,omitempty"`
}

type WdkMealPosGetfetchmealcodeResponse struct {

    // 取餐号
    FetchMealCode   string `json:"fetch_meal_code,omitempty"`

}
