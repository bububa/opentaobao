package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询官方的活动名称接口 APIRequest
taobao.singletreasure.activity.name.query

查询官方的活动名称列表接口
*/
type TaobaoSingletreasureActivityNameQueryRequest struct {
    model.Params

}

func NewTaobaoSingletreasureActivityNameQueryRequest() *TaobaoSingletreasureActivityNameQueryRequest{
    return &TaobaoSingletreasureActivityNameQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityNameQueryRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.name.query"
}

func (r TaobaoSingletreasureActivityNameQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


