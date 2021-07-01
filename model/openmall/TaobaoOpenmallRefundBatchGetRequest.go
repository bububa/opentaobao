package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取openmall退款单 API请求
taobao.openmall.refund.batch.get

批量获取openmall退款单
注意：该接口信息存在延迟，如需实时详情请访问taobao.openmall.refund.get
*/
type TaobaoOpenmallRefundBatchGetAPIRequest struct {
    model.Params
    // 查询范围结束时间，闭区间
    _endCreated   string
    // 翻页页码，从1开始
    _pageIndex   int64
    // 页面大小，不超过100
    _pageSize   int64
    // 查询的渠道商Nick
    _distributor   string
    // 查询范围开始时间，闭区间
    _startCreated   string
}

// 初始化TaobaoOpenmallRefundBatchGetAPIRequest对象
func NewTaobaoOpenmallRefundBatchGetRequest() *TaobaoOpenmallRefundBatchGetAPIRequest{
    return &TaobaoOpenmallRefundBatchGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallRefundBatchGetAPIRequest) GetApiMethodName() string {
    return "taobao.openmall.refund.batch.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallRefundBatchGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EndCreated Setter
// 查询范围结束时间，闭区间
func (r *TaobaoOpenmallRefundBatchGetAPIRequest) SetEndCreated(_endCreated string) error {
    r._endCreated = _endCreated
    r.Set("end_created", _endCreated)
    return nil
}

// EndCreated Getter
func (r TaobaoOpenmallRefundBatchGetAPIRequest) GetEndCreated() string {
    return r._endCreated
}
// PageIndex Setter
// 翻页页码，从1开始
func (r *TaobaoOpenmallRefundBatchGetAPIRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpenmallRefundBatchGetAPIRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 页面大小，不超过100
func (r *TaobaoOpenmallRefundBatchGetAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpenmallRefundBatchGetAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
// Distributor Setter
// 查询的渠道商Nick
func (r *TaobaoOpenmallRefundBatchGetAPIRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallRefundBatchGetAPIRequest) GetDistributor() string {
    return r._distributor
}
// StartCreated Setter
// 查询范围开始时间，闭区间
func (r *TaobaoOpenmallRefundBatchGetAPIRequest) SetStartCreated(_startCreated string) error {
    r._startCreated = _startCreated
    r.Set("start_created", _startCreated)
    return nil
}

// StartCreated Getter
func (r TaobaoOpenmallRefundBatchGetAPIRequest) GetStartCreated() string {
    return r._startCreated
}
