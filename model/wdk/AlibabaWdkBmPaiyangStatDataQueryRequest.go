package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
派样统计数据查询 APIRequest
alibaba.wdk.bm.paiyang.stat.data.query

派样统计数据查询
*/
type AlibabaWdkBmPaiyangStatDataQueryRequest struct {
    model.Params

    // 入参对象
    param   *PaiyangStatDataParam 

}

func NewAlibabaWdkBmPaiyangStatDataQueryRequest() *AlibabaWdkBmPaiyangStatDataQueryRequest{
    return &AlibabaWdkBmPaiyangStatDataQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkBmPaiyangStatDataQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.paiyang.stat.data.query"
}

func (r AlibabaWdkBmPaiyangStatDataQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkBmPaiyangStatDataQueryRequest) SetParam(param *PaiyangStatDataParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r AlibabaWdkBmPaiyangStatDataQueryRequest) GetParam() *PaiyangStatDataParam {
    return r.param
}

