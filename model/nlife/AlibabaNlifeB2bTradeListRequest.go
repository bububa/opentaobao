package nlife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取企业下的采购单列表 API请求
alibaba.nlife.b2b.trade.list

获取指定门店下的采购单列表
*/
type AlibabaNlifeB2bTradeListRequest struct {
    model.Params
    // 采购单生效时间开始范围
    _startEffectiveDate   string
    // 采购单生效时间结束范围
    _endEffectiveDate   string
    // 查询的页码
    _pageNo   int64
    // 每页的数量
    _pageSize   int64
    // 企业ID
    _entId   int64
}

// 初始化AlibabaNlifeB2bTradeListRequest对象
func NewAlibabaNlifeB2bTradeListRequest() *AlibabaNlifeB2bTradeListRequest{
    return &AlibabaNlifeB2bTradeListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaNlifeB2bTradeListRequest) GetApiMethodName() string {
    return "alibaba.nlife.b2b.trade.list"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaNlifeB2bTradeListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartEffectiveDate Setter
// 采购单生效时间开始范围
func (r *AlibabaNlifeB2bTradeListRequest) SetStartEffectiveDate(_startEffectiveDate string) error {
    r._startEffectiveDate = _startEffectiveDate
    r.Set("start_effective_date", _startEffectiveDate)
    return nil
}

// StartEffectiveDate Getter
func (r AlibabaNlifeB2bTradeListRequest) GetStartEffectiveDate() string {
    return r._startEffectiveDate
}
// EndEffectiveDate Setter
// 采购单生效时间结束范围
func (r *AlibabaNlifeB2bTradeListRequest) SetEndEffectiveDate(_endEffectiveDate string) error {
    r._endEffectiveDate = _endEffectiveDate
    r.Set("end_effective_date", _endEffectiveDate)
    return nil
}

// EndEffectiveDate Getter
func (r AlibabaNlifeB2bTradeListRequest) GetEndEffectiveDate() string {
    return r._endEffectiveDate
}
// PageNo Setter
// 查询的页码
func (r *AlibabaNlifeB2bTradeListRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AlibabaNlifeB2bTradeListRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页的数量
func (r *AlibabaNlifeB2bTradeListRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlibabaNlifeB2bTradeListRequest) GetPageSize() int64 {
    return r._pageSize
}
// EntId Setter
// 企业ID
func (r *AlibabaNlifeB2bTradeListRequest) SetEntId(_entId int64) error {
    r._entId = _entId
    r.Set("ent_id", _entId)
    return nil
}

// EntId Getter
func (r AlibabaNlifeB2bTradeListRequest) GetEntId() int64 {
    return r._entId
}
