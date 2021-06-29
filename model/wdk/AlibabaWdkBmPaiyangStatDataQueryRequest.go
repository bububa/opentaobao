package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
派样统计数据查询 API请求
alibaba.wdk.bm.paiyang.stat.data.query

派样统计数据查询
*/
type AlibabaWdkBmPaiyangStatDataQueryRequest struct {
    model.Params
    // 入参对象
    param   *PaiyangStatDataParam
}

// 初始化AlibabaWdkBmPaiyangStatDataQueryRequest对象
func NewAlibabaWdkBmPaiyangStatDataQueryRequest() *AlibabaWdkBmPaiyangStatDataQueryRequest{
    return &AlibabaWdkBmPaiyangStatDataQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmPaiyangStatDataQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.paiyang.stat.data.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmPaiyangStatDataQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AlibabaWdkBmPaiyangStatDataQueryRequest) SetParam(param *PaiyangStatDataParam) error {
    r.param = param
    r.Set("param", param)
    return nil
}

// Param Getter
func (r AlibabaWdkBmPaiyangStatDataQueryRequest) GetParam() *PaiyangStatDataParam {
    return r.param
}
