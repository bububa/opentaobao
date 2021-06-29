package mos

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
大pos新选单退 API请求
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

// 初始化AlibabaMosOrderqsMisbigposOrderQueryRequest对象
func NewAlibabaMosOrderqsMisbigposOrderQueryRequest() *AlibabaMosOrderqsMisbigposOrderQueryRequest{
    return &AlibabaMosOrderqsMisbigposOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetApiMethodName() string {
    return "alibaba.mos.orderqs.misbigpos.order.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StoreNo Setter
// 外部门店号
func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetStoreNo(storeNo string) error {
    r.storeNo = storeNo
    r.Set("store_no", storeNo)
    return nil
}

// StoreNo Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetStoreNo() string {
    return r.storeNo
}
// QueryBaseData Setter
// 基本信息获取参数
func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetQueryBaseData(queryBaseData bool) error {
    r.queryBaseData = queryBaseData
    r.Set("query_base_data", queryBaseData)
    return nil
}

// QueryBaseData Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetQueryBaseData() bool {
    return r.queryBaseData
}
// ReceiptNo Setter
// 小票号
func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetReceiptNo(receiptNo string) error {
    r.receiptNo = receiptNo
    r.Set("receipt_no", receiptNo)
    return nil
}

// ReceiptNo Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetReceiptNo() string {
    return r.receiptNo
}
// QueryCouponExtern Setter
// 券扩展数据获取
func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetQueryCouponExtern(queryCouponExtern bool) error {
    r.queryCouponExtern = queryCouponExtern
    r.Set("query_coupon_extern", queryCouponExtern)
    return nil
}

// QueryCouponExtern Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetQueryCouponExtern() bool {
    return r.queryCouponExtern
}
