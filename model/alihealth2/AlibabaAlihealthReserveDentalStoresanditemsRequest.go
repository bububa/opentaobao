package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商户门店，商品列表 API请求
alibaba.alihealth.reserve.dental.storesanditems

查询商户门店，商品列表
*/
type AlibabaAlihealthReserveDentalStoresanditemsRequest struct {
    model.Params
    // 页码，每页100个门店，超过100个门店分页请求
    _pageNo   int64
}

// 初始化AlibabaAlihealthReserveDentalStoresanditemsRequest对象
func NewAlibabaAlihealthReserveDentalStoresanditemsRequest() *AlibabaAlihealthReserveDentalStoresanditemsRequest{
    return &AlibabaAlihealthReserveDentalStoresanditemsRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthReserveDentalStoresanditemsRequest) GetApiMethodName() string {
    return "alibaba.alihealth.reserve.dental.storesanditems"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthReserveDentalStoresanditemsRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PageNo Setter
// 页码，每页100个门店，超过100个门店分页请求
func (r *AlibabaAlihealthReserveDentalStoresanditemsRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaAlihealthReserveDentalStoresanditemsRequest) GetPageNo() int64 {
    return r._pageNo
}
