package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商银行账号查询 APIRequest
alibaba.mos.finance.bankinfo.querybank

查询供应商对应的银行账号信息
*/
type AlibabaMosFinanceBankinfoQuerybankRequest struct {
    model.Params

    // 供应商id
    supplierId   string 

    // 门店号
    storeNo   string 

    // 签约主体id
    companyId   string 

}

func NewAlibabaMosFinanceBankinfoQuerybankRequest() *AlibabaMosFinanceBankinfoQuerybankRequest{
    return &AlibabaMosFinanceBankinfoQuerybankRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosFinanceBankinfoQuerybankRequest) GetApiMethodName() string {
    return "alibaba.mos.finance.bankinfo.querybank"
}

func (r AlibabaMosFinanceBankinfoQuerybankRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosFinanceBankinfoQuerybankRequest) SetSupplierId(supplierId string) error {
    r.supplierId = supplierId
    r.Set("supplier_id", supplierId)
    return nil
}

func (r AlibabaMosFinanceBankinfoQuerybankRequest) GetSupplierId() string {
    return r.supplierId
}

func (r *AlibabaMosFinanceBankinfoQuerybankRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

func (r AlibabaMosFinanceBankinfoQuerybankRequest) GetStoreNo() string {
    return r.storeNo
}

func (r *AlibabaMosFinanceBankinfoQuerybankRequest) SetCompanyId(companyId string) error {
    r.companyId = companyId
    r.Set("company_id", companyId)
    return nil
}

func (r AlibabaMosFinanceBankinfoQuerybankRequest) GetCompanyId() string {
    return r.companyId
}

