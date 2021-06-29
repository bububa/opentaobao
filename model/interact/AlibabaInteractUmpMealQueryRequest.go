package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝卖家搭配套餐查询 API请求
alibaba.interact.ump.meal.query

查询卖家在优惠平台设置的搭配套餐列表，每个套餐包括名称、套餐价格、手淘套餐购买链接
*/
type AlibabaInteractUmpMealQueryRequest struct {
    model.Params
}

// 初始化AlibabaInteractUmpMealQueryRequest对象
func NewAlibabaInteractUmpMealQueryRequest() *AlibabaInteractUmpMealQueryRequest{
    return &AlibabaInteractUmpMealQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractUmpMealQueryRequest) GetApiMethodName() string {
    return "alibaba.interact.ump.meal.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractUmpMealQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
