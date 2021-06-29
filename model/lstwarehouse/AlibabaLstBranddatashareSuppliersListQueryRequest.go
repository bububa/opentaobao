package lstwarehouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌数据授权的供应商列表 API请求
alibaba.lst.branddatashare.suppliers.list.query

品牌商查询品牌数据授权的供应商列表
*/
type AlibabaLstBranddatashareSuppliersListQueryRequest struct {
    model.Params
    // 入参
    _query   *LstBmSupplierQuery
}

// 初始化AlibabaLstBranddatashareSuppliersListQueryRequest对象
func NewAlibabaLstBranddatashareSuppliersListQueryRequest() *AlibabaLstBranddatashareSuppliersListQueryRequest{
    return &AlibabaLstBranddatashareSuppliersListQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstBranddatashareSuppliersListQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.branddatashare.suppliers.list.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstBranddatashareSuppliersListQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参
func (r *AlibabaLstBranddatashareSuppliersListQueryRequest) SetQuery(_query *LstBmSupplierQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaLstBranddatashareSuppliersListQueryRequest) GetQuery() *LstBmSupplierQuery {
    return r._query
}
