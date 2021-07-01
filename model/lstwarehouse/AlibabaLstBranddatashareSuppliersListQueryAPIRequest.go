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
type AlibabaLstBranddatashareSuppliersListQueryAPIRequest struct {
    model.Params
    // 入参
    _query   *LstBmSupplierQuery
}

// 初始化AlibabaLstBranddatashareSuppliersListQueryAPIRequest对象
func NewAlibabaLstBranddatashareSuppliersListQueryRequest() *AlibabaLstBranddatashareSuppliersListQueryAPIRequest{
    return &AlibabaLstBranddatashareSuppliersListQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstBranddatashareSuppliersListQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.lst.branddatashare.suppliers.list.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstBranddatashareSuppliersListQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参
func (r *AlibabaLstBranddatashareSuppliersListQueryAPIRequest) SetQuery(_query *LstBmSupplierQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaLstBranddatashareSuppliersListQueryAPIRequest) GetQuery() *LstBmSupplierQuery {
    return r._query
}
