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
    startEffectiveDate   string
    // 采购退货单创建时间结束范围
    endEffectiveDate   string
    // 查询的页数
    pageNo   int64
    // 每页的数量
    pageSize   int64
    // 企业Id
    entId   int64
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
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetStartEffectiveDate(startEffectiveDate string) error {
    r.startEffectiveDate = startEffectiveDate
    r.Set("start_effective_date", startEffectiveDate)
    return nil
}

// StartEffectiveDate Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetStartEffectiveDate() string {
    return r.startEffectiveDate
}
// EndEffectiveDate Setter
// 采购退货单创建时间结束范围
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetEndEffectiveDate(endEffectiveDate string) error {
    r.endEffectiveDate = endEffectiveDate
    r.Set("end_effective_date", endEffectiveDate)
    return nil
}

// EndEffectiveDate Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetEndEffectiveDate() string {
    return r.endEffectiveDate
}
// PageNo Setter
// 查询的页数
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页的数量
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetPageSize() int64 {
    return r.pageSize
}
// EntId Setter
// 企业Id
func (r *AlibabaNlifeB2bTradeRefundListRequest) SetEntId(entId int64) error {
    r.entId = entId
    r.Set("ent_id", entId)
    return nil
}

// EntId Getter
func (r AlibabaNlifeB2bTradeRefundListRequest) GetEntId() int64 {
    return r.entId
}
