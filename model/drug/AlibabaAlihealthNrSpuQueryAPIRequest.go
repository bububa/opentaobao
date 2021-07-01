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
type AlibabaAlihealthNrSpuQueryAPIRequest struct {
    model.Params
    // 标品查询条件
    _query   *TopAlihealthSpuQuery
    // 查询选择器
    _options   *TopAlihealthSpuQueryOptions
}

// 初始化AlibabaAlihealthNrSpuQueryAPIRequest对象
func NewAlibabaAlihealthNrSpuQueryRequest() *AlibabaAlihealthNrSpuQueryAPIRequest{
    return &AlibabaAlihealthNrSpuQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthNrSpuQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.nr.spu.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthNrSpuQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 标品查询条件
func (r *AlibabaAlihealthNrSpuQueryAPIRequest) SetQuery(_query *TopAlihealthSpuQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaAlihealthNrSpuQueryAPIRequest) GetQuery() *TopAlihealthSpuQuery {
    return r._query
}
// Options Setter
// 查询选择器
func (r *AlibabaAlihealthNrSpuQueryAPIRequest) SetOptions(_options *TopAlihealthSpuQueryOptions) error {
    r._options = _options
    r.Set("options", _options)
    return nil
}

// Options Getter
func (r AlibabaAlihealthNrSpuQueryAPIRequest) GetOptions() *TopAlihealthSpuQueryOptions {
    return r._options
}
