package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos新选单退 APIRequest
alibaba.mos.orderqs.misbigpos.order.query

大pos新选单退
*/
type AlibabaMosOrderqsMisbigposOrderQueryRequest struct {
    model.Params

    // 外部门店号
    storeNo   string 

    // 基本信息获取参数
    queryBaseData   bool 

    // 小票号
    receiptNo   string 

    // 券扩展数据获取
    queryCouponExtern   bool 

}

func NewAlibabaMosOrderqsMisbigposOrderQueryRequest() *AlibabaMosOrderqsMisbigposOrderQueryRequest{
    return &AlibabaMosOrderqsMisbigposOrderQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.mos.orderqs.misbigpos.order.query"
}

func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetStoreNo() string {
    return r.storeNo
}

func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetQueryBaseData(queryBaseData bool) error {
    r.queryBaseData = queryBaseData
    r.Set("query_base_data", queryBaseData)
    return nil
}

func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetQueryBaseData() bool {
    return r.queryBaseData
}

func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetReceiptNo(receiptNo string) error {
    r.receiptNo = receiptNo
    r.Set("receipt_no", receiptNo)
    return nil
}

func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetReceiptNo() string {
    return r.receiptNo
}

func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetQueryCouponExtern(queryCouponExtern bool) error {
    r.queryCouponExtern = queryCouponExtern
    r.Set("query_coupon_extern", queryCouponExtern)
    return nil
}

func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetQueryCouponExtern() bool {
    return r.queryCouponExtern
}

