package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询官方的活动名称接口 API请求
taobao.singletreasure.activity.name.query

查询官方的活动名称列表接口
*/
type TaobaoSingletreasureActivityNameQueryRequest struct {
    model.Params
}

// 初始化TaobaoSingletreasureActivityNameQueryRequest对象
func NewTaobaoSingletreasureActivityNameQueryRequest() *TaobaoSingletreasureActivityNameQueryRequest{
    return &TaobaoSingletreasureActivityNameQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoSingletreasureActivityNameQueryRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.name.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoSingletreasureActivityNameQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
