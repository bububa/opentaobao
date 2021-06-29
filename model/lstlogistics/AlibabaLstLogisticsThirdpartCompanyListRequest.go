package lstlogistics

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商-异云-第三方物流公司列表 APIRequest
alibaba.lst.logistics.thirdpart.company.list

异地云仓发货时，需填写的第三方物流公司列表
*/
type AlibabaLstLogisticsThirdpartCompanyListRequest struct {
    model.Params

    // 入参
    query   *LstLogisticsCompanyQuery 

}

func NewAlibabaLstLogisticsThirdpartCompanyListRequest() *AlibabaLstLogisticsThirdpartCompanyListRequest{
    return &AlibabaLstLogisticsThirdpartCompanyListRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaLstLogisticsThirdpartCompanyListRequest) GetApiMethodName() string {
    return "alibaba.lst.logistics.thirdpart.company.list"
}

func (r AlibabaLstLogisticsThirdpartCompanyListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaLstLogisticsThirdpartCompanyListRequest) SetQuery(query *LstLogisticsCompanyQuery) error {
    r.query = query
    r.Set("query", query)
    return nil
}

func (r AlibabaLstLogisticsThirdpartCompanyListRequest) GetQuery() *LstLogisticsCompanyQuery {
    return r.query
}

