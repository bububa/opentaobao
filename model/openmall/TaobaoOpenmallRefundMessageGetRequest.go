package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
openmall获取退款单留言 API请求
taobao.openmall.refund.message.get

openmall获取退款单留言
*/
type TaobaoOpenmallRefundMessageGetRequest struct {
    model.Params
    // 分销者身份
    distributor   string
    // 翻页页码
    pageNo   int64
    // 翻页大小
    pageSize   int64
    // 退款单号
    refundId   int64
}

// 初始化TaobaoOpenmallRefundMessageGetRequest对象
func NewTaobaoOpenmallRefundMessageGetRequest() *TaobaoOpenmallRefundMessageGetRequest{
    return &TaobaoOpenmallRefundMessageGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundMessageGetRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.message.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundMessageGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 分销者身份
func (r *TaobaoOpenmallRefundMessageGetRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundMessageGetRequest) GetDistributor() string {
    return r.distributor
}
// PageNo Setter
// 翻页页码
func (r *TaobaoOpenmallRefundMessageGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoOpenmallRefundMessageGetRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 翻页大小
func (r *TaobaoOpenmallRefundMessageGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpenmallRefundMessageGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// RefundId Setter
// 退款单号
func (r *TaobaoOpenmallRefundMessageGetRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoOpenmallRefundMessageGetRequest) GetRefundId() int64 {
    return r.refundId
}
