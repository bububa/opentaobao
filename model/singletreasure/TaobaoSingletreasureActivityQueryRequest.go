package singletreasure

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询活动列表接口 APIRequest
taobao.singletreasure.activity.query

查询活动列表接口
*/
type TaobaoSingletreasureActivityQueryRequest struct {
    model.Params

    // 查询对象
    query   *PageQueryDto 

}

func NewTaobaoSingletreasureActivityQueryRequest() *TaobaoSingletreasureActivityQueryRequest{
    return &TaobaoSingletreasureActivityQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSingletreasureActivityQueryRequest) GetApiMethodName() string {
    return "taobao.singletreasure.activity.query"
}

func (r TaobaoSingletreasureActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSingletreasureActivityQueryRequest) SetQuery(query *PageQueryDto) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r TaobaoSingletreasureActivityQueryRequest) GetQuery() *PageQueryDto {
    return r.query
}

