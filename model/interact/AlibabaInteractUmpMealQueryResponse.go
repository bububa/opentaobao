package interact

import (
    "github.com/bububa/opentaobao/model"
)

/* 
淘宝卖家搭配套餐查询 APIResponse
alibaba.interact.ump.meal.query

查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接
*/
type AlibabaInteractUmpMealQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaInteractUmpMealQueryResponse `json:"alibaba_interact_ump_meal_query_response,omitempty"` 
    AlibabaInteractUmpMealQueryResponse
}

/* model for simplify = false
type AlibabaInteractUmpMealQueryResponse struct {

    // 优惠平台搭配套餐对象
    
    Meals  struct {
        OpenMealDo  []OpenMealDo `json:"open_meal_do,omitempty"`
    } `json:"meals,omitempty"`
    

}
*/

type AlibabaInteractUmpMealQueryResponse struct {

    // 优惠平台搭配套餐对象
    Meals   []OpenMealDo `json:"meals,omitempty"`

}
