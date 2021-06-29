package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-第三方物流公司列表 API请求
alibaba.lst.logistics.thirdpart.company.list

异地云仓发货时，需填写的第三方物流公司列表
*/
type AlibabaLstLogisticsThirdpartCompanyListRequest struct {
    model.Params
    // 入参
    _query   *LstLogisticsCompanyQuery
}

// 初始化AlibabaLstLogisticsThirdpartCompanyListRequest对象
func NewAlibabaLstLogisticsThirdpartCompanyListRequest() *AlibabaLstLogisticsThirdpartCompanyListRequest{
    return &AlibabaLstLogisticsThirdpartCompanyListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstLogisticsThirdpartCompanyListRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.thirdpart.company.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstLogisticsThirdpartCompanyListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Query Setter
// 入参
func (r *AlibabaLstLogisticsThirdpartCompanyListRequest) SetQuery(_query *LstLogisticsCompanyQuery) error {
    r._query = _query
    r.Set("query", _query)
    return nil
}

// Query Getter
func (r AlibabaLstLogisticsThirdpartCompanyListRequest) GetQuery() *LstLogisticsCompanyQuery {
    return r._query
}
