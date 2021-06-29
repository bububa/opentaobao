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
    _storeNo   string
    // 基本信息获取参数
    _queryBaseData   bool
    // 小票号
    _receiptNo   string
    // 券扩展数据获取
    _queryCouponExtern   bool
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
func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetStoreNo(_storeNo string) error {
    r._storeNo = _storeNo
    r.Set("store_no", _storeNo)
    return nil
}

// StoreNo Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetStoreNo() string {
    return r._storeNo
}
// QueryBaseData Setter
// 基本信息获取参数
func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetQueryBaseData(_queryBaseData bool) error {
    r._queryBaseData = _queryBaseData
    r.Set("query_base_data", _queryBaseData)
    return nil
}

// QueryBaseData Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetQueryBaseData() bool {
    return r._queryBaseData
}
// ReceiptNo Setter
// 小票号
func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetReceiptNo(_receiptNo string) error {
    r._receiptNo = _receiptNo
    r.Set("receipt_no", _receiptNo)
    return nil
}

// ReceiptNo Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetReceiptNo() string {
    return r._receiptNo
}
// QueryCouponExtern Setter
// 券扩展数据获取
func (r *AlibabaMosOrderqsMisbigposOrderQueryRequest) SetQueryCouponExtern(_queryCouponExtern bool) error {
    r._queryCouponExtern = _queryCouponExtern
    r.Set("query_coupon_extern", _queryCouponExtern)
    return nil
}

// QueryCouponExtern Getter
func (r AlibabaMosOrderqsMisbigposOrderQueryRequest) GetQueryCouponExtern() bool {
    return r._queryCouponExtern
}
