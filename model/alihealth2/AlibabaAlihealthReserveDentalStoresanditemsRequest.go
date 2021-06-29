package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商户门店，商品列表 APIRequest
alibaba.alihealth.reserve.dental.storesanditems

查询商户门店，商品列表
*/
type AlibabaAlihealthReserveDentalStoresanditemsRequest struct {
    model.Params

    // 页码，每页100个门店，超过100个门店分页请求
    pageNo   int64 

}

func NewAlibabaAlihealthReserveDentalStoresanditemsRequest() *AlibabaAlihealthReserveDentalStoresanditemsRequest{
    return &AlibabaAlihealthReserveDentalStoresanditemsRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthReserveDentalStoresanditemsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reserve.dental.storesanditems"
}

func (r AlibabaAlihealthReserveDentalStoresanditemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthReserveDentalStoresanditemsRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AlibabaAlihealthReserveDentalStoresanditemsRequest) GetPageNo() int64 {
    return r.pageNo
}

