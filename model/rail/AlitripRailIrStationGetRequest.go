package rail

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国际火车票标准车站查询 API请求
alitrip.rail.ir.station.get

国际火车票提供给代理商用于查询标准车站信息，用于代理商对自己的车站与飞猪平台的车站做映射
*/
type AlitripRailIrStationGetRequest struct {
    model.Params
    // 商家id
    _agentId   int64
    // 页数 从1开始
    _pageIndex   int64
    // 每页条数
    _pageSize   int64
}

// 初始化AlitripRailIrStationGetRequest对象
func NewAlitripRailIrStationGetRequest() *AlitripRailIrStationGetRequest{
    return &AlitripRailIrStationGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripRailIrStationGetRequest) GetApiMethodName() string {
    return "alitrip.rail.ir.station.get"
}

// IRequest interface 方法, 获取API参数
func (r AlitripRailIrStationGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 商家id
func (r *AlitripRailIrStationGetRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r AlitripRailIrStationGetRequest) GetAgentId() int64 {
    return r._agentId
}
// PageIndex Setter
// 页数 从1开始
func (r *AlitripRailIrStationGetRequest) SetPageIndex(_pageIndex int64) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r AlitripRailIrStationGetRequest) GetPageIndex() int64 {
    return r._pageIndex
}
// PageSize Setter
// 每页条数
func (r *AlitripRailIrStationGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AlitripRailIrStationGetRequest) GetPageSize() int64 {
    return r._pageSize
}
