package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
搭配套餐查询 APIRequest
taobao.promotion.meal.get

搭配套餐查询。每个卖家最多创建50个搭配套餐，所以查询不会分页，会将所有的满足状态的搭配套餐全部查出。该接口不会校验商品的下架或库存为0，查询结果的状态表明搭配套餐在数据库中的状态，商品的状态请isv自己验证。在卖家后台页面点击查看会触发数据库状态的修改。
*/
type TaobaoPromotionMealGetRequest struct {
    model.Params

    // 搭配套餐id
    mealId   int64 

    // 套餐状态。有效：VALID;失效：INVALID(有效套餐为可使用的套餐,无效套餐为套餐中有商品下架或库存为0时)。默认时两种情况都会查询。
    status   string 

}

func NewTaobaoPromotionMealGetRequest() *TaobaoPromotionMealGetRequest{
    return &TaobaoPromotionMealGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPromotionMealGetRequest) GetApiMethodName() string {
    return "taobao.promotion.meal.get"
}

func (r TaobaoPromotionMealGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPromotionMealGetRequest) SetMealId(mealId int64) error {
    r.mealId = mealId
    r.Set("meal_id", mealId)
    return nil
}

func (r TaobaoPromotionMealGetRequest) GetMealId() int64 {
    return r.mealId
}

func (r *TaobaoPromotionMealGetRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r TaobaoPromotionMealGetRequest) GetStatus() string {
    return r.status
}

