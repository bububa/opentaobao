package sungari

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商品商家处置结果查询 APIRequest
taobao.sungari.dispose.query

红盾云桥同政府合作，将线下寄函的商品商家处置转为线上处理
*/
type TaobaoSungariDisposeQueryRequest struct {
    model.Params

    // 查询的key列表
    paramList   []string 

}

func NewTaobaoSungariDisposeQueryRequest() *TaobaoSungariDisposeQueryRequest{
    return &TaobaoSungariDisposeQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoSungariDisposeQueryRequest) GetApiMethodName() string {
    return "taobao.sungari.dispose.query"
}

func (r TaobaoSungariDisposeQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoSungariDisposeQueryRequest) SetParamList(paramList []string) error {
    r.paramList = paramList
    r.Set("param_list", paramList)
    return nil
}

func (r TaobaoSungariDisposeQueryRequest) GetParamList() []string {
    return r.paramList
}

