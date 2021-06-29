package lstwarehouse

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
品牌数据授权的供应商列表 APIRequest
alibaba.lst.branddatashare.suppliers.list.query

品牌商查询品牌数据授权的供应商列表
*/
type AlibabaLstBranddatashareSuppliersListQueryRequest struct {
    model.Params

    // 入参
    query   *LstBmSupplierQuery 

}

func NewAlibabaLstBranddatashareSuppliersListQueryRequest() *AlibabaLstBranddatashareSuppliersListQueryRequest{
    return &AlibabaLstBranddatashareSuppliersListQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstBranddatashareSuppliersListQueryRequest) GetApiMethodName() string {
    return "alibaba.lst.branddatashare.suppliers.list.query"
}

func (r AlibabaLstBranddatashareSuppliersListQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstBranddatashareSuppliersListQueryRequest) SetQuery(query *LstBmSupplierQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaLstBranddatashareSuppliersListQueryRequest) GetQuery() *LstBmSupplierQuery {
    return r.query
}

