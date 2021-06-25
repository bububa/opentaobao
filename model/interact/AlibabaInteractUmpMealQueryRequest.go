package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝卖家搭配套餐查询 APIRequest
alibaba.interact.ump.meal.query

查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接
*/
type AlibabaInteractUmpMealQueryRequest struct {
    model.Params

}

func NewAlibabaInteractUmpMealQueryRequest() *AlibabaInteractUmpMealQueryRequest{
    return &AlibabaInteractUmpMealQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractUmpMealQueryRequest) GetApiMethodName() string {
    return "alibaba.interact.ump.meal.query"
}

func (r AlibabaInteractUmpMealQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


