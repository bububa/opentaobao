package interact

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝卖家搭配套餐查询 APIResponse
alibaba.interact.ump.meal.query

查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接
*/
type AlibabaInteractUmpMealQueryAPIResponse struct {
    model.CommonResponse
    AlibabaInteractUmpMealQueryResponse
}

type AlibabaInteractUmpMealQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_interact_ump_meal_query_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 优惠平台搭配套餐对象
    
    Meals   []OpenMealDo `json:"meals,omitempty" xml:"meals>open_meal_do,omitempty"`
    
    
}
