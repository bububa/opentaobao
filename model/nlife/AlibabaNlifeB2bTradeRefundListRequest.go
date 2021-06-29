package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取采购退货单列表 API请求
alibaba.nlife.b2b.trade.refund.list

获取采购退货单列表
*/
type AlibabaNlifeB2bTradeRefundListRequest struct {
    model.Params
    // 采购退货单创建时间开始范围
    _startEffectiveDate   string
    // 采购退货单创建时间结束范围
    _endEffectiveDate   string
    // 查询的页数
    _pageNo   int64
    // 每页的数量
    _pageSize   int64
    // 企业Id
    _entId   int64
}

// 初始化AlibabaNlifeB2bTradeRefundListRequest对象
func NewAlibabaNlifeB2bTradeRefundListRequest() *AlibabaNlifeB2bTradeRefundListRequest{
    return &AlibabaNlifeB2bTradeRefundListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2bTradeRefundListRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2b.trade.refund.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2bTradeRefundListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartEffectiveDate Setter
// 采购退货单创建时间开始范围
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetStartEffectiveDate(_startEffectiveDate string) error {
    r._startEffectiveDate = _startEffectiveDate
    r.Set("start_effective_date", _startEffectiveDate)
    return nil
}

// StartEffectiveDate Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetStartEffectiveDate() string {
    return r._startEffectiveDate
}
// EndEffectiveDate Setter
// 采购退货单创建时间结束范围
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetEndEffectiveDate(_endEffectiveDate string) error {
    r._endEffectiveDate = _endEffectiveDate
    r.Set("end_effective_date", _endEffectiveDate)
    return nil
}

// EndEffectiveDate Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetEndEffectiveDate() string {
    return r._endEffectiveDate
}
// PageNo Setter
// 查询的页数
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页的数量
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetPageSize() int64 {
    return r._pageSize
}
// EntId Setter
// 企业Id
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetEntId(_entId int64) error {
    r._entId = _entId
    r.Set("ent_id", _entId)
    return nil
}

// EntId Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetEntId() int64 {
    return r._entId
}
