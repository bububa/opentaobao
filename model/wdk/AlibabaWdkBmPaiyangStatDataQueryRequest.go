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
type AlibabaWdkBmPaiyangStatDataQueryAPIRequest struct {
    model.Params
    // 入参对象
    _param   *PaiyangStatDataParam
}

// 初始化AlibabaWdkBmPaiyangStatDataQueryAPIRequest对象
func NewAlibabaWdkBmPaiyangStatDataQueryRequest() *AlibabaWdkBmPaiyangStatDataQueryAPIRequest{
    return &AlibabaWdkBmPaiyangStatDataQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmPaiyangStatDataQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.wdk.bm.paiyang.stat.data.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmPaiyangStatDataQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 入参对象
func (r *AlibabaWdkBmPaiyangStatDataQueryAPIRequest) SetParam(_param *PaiyangStatDataParam) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r AlibabaWdkBmPaiyangStatDataQueryAPIRequest) GetParam() *PaiyangStatDataParam {
    return r._param
}
