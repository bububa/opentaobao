package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取openmall订单 API请求
taobao.openmall.trade.batch.get

批量获取openmall订单
注意：该接口数据存在延迟，实时数据请通过taobao.openmall.trade.get获取
*/
type TaobaoOpenmallTradeBatchGetRequest struct {
    model.Params
    // 查询范围结束时间，闭区间
    _endCreated   string
    // 查询页码，从1开始
    _pageIndex   int64
    // 页面大小，不超过100
    _pageSize   int64
    // 渠道商Nick
    _distributor   string
    // 查询范围开始时间，闭区间
    _startCreated   string
}

// 初始化TaobaoOpenmallTradeBatchGetRequest对象
func NewTaobaoOpenmallTradeBatchGetRequest() *TaobaoOpenmallTradeBatchGetRequest{
    return &TaobaoOpenmallTradeBatchGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeBatchGetRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.batch.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeBatchGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EndCreated Setter
// 查询范围结束时间，闭区间
func (r *TaobaoOpenmallTradeBatchGetRequest) SetEndCreated(_endCreated string) error {
    r._endCreated = _endCreated
    r.Set("end_created", _endCreated)
    return nil
}

// EndCreated Getter
func (r TaobaoOpenmallTradeBatchGetRequest) GetEndCreated() string {
    return r._endCreated
}
// PageIndex Setter
// 查询页码，从1开始
func (r *TaobaoOpenmallTradeBatchGetRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r TaobaoOpenmallTradeBatchGetRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 页面大小，不超过100
func (r *TaobaoOpenmallTradeBatchGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoOpenmallTradeBatchGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// Distributor Setter
// 渠道商Nick
func (r *TaobaoOpenmallTradeBatchGetRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeBatchGetRequest) GetDistributor() string {
    return r._distributor
}
// StartCreated Setter
// 查询范围开始时间，闭区间
func (r *TaobaoOpenmallTradeBatchGetRequest) SetStartCreated(_startCreated string) error {
    r._startCreated = _startCreated
    r.Set("start_created", _startCreated)
    return nil
}

// StartCreated Getter
func (r TaobaoOpenmallTradeBatchGetRequest) GetStartCreated() string {
    return r._startCreated
}
