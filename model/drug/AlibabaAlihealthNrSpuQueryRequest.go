package drug

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取标品库标品信息 API请求
alibaba.alihealth.nr.spu.query

提供给ERP使用的，获取健康标品库信息
*/
type AlibabaAlihealthNrSpuQueryRequest struct {
    model.Params
    // 标品查询条件
    query   *TopAlihealthSpuQuery
    // 查询选择器
    options   *TopAlihealthSpuQueryOptions
}

// 初始化AlibabaAlihealthNrSpuQueryRequest对象
func NewAlibabaAlihealthNrSpuQueryRequest() *AlibabaAlihealthNrSpuQueryRequest{
    return &AlibabaAlihealthNrSpuQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrSpuQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.spu.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrSpuQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 标品查询条件
func (r *AlibabaAlihealthNrSpuQueryRequest) SetQuery(query *TopAlihealthSpuQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

// Query Getter
func (r AlibabaAlihealthNrSpuQueryRequest) GetQuery() *TopAlihealthSpuQuery {
    return r.query
}
// Options Setter
// 查询选择器
func (r *AlibabaAlihealthNrSpuQueryRequest) SetOptions(options *TopAlihealthSpuQueryOptions) error {
    r.options = options
    r.Set("options", options)
    return nil
}

// Options Getter
func (r AlibabaAlihealthNrSpuQueryRequest) GetOptions() *TopAlihealthSpuQueryOptions {
    return r.options
}
