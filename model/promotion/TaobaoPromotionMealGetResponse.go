package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
搭配套餐查询 APIResponse
taobao.promotion.meal.get

搭配套餐查询。每个卖家最多创建50个搭配套餐，所以查询不会分页，会将所有的满足状态的搭配套餐全部查出。该接口不会校验商品的下架或库存为0，查询结果的状态表明搭配套餐在数据库中的状态，商品的状态请isv自己验证。在卖家后台页面点击查看会触发数据库状态的修改。
*/
type TaobaoPromotionMealGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoPromotionMealGetResponse `json:"taobao_promotion_meal_get_response,omitempty"`
}

type TaobaoPromotionMealGetResponse struct {

    // 搭配套餐列表。
    MealList   []Meal `json:"meal_list,omitempty"`

}
