package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailIrStationGetAPIRequest 国际火车票标准车站查询 API请求
// alitrip.rail.ir.station.get
//
// 国际火车票提供给代理商用于查询标准车站信息，用于代理商对自己的车站与飞猪平台的车站做映射
type AlitripRailIrStationGetAPIRequest struct {
	model.Params
	// 商家id
	_agentId int64
	// 页数 从1开始
	_pageIndex int64
	// 每页条数
	_pageSize int64
}

// NewAlitripRailIrStationGetRequest 初始化AlitripRailIrStationGetAPIRequest对象
func NewAlitripRailIrStationGetRequest() *AlitripRailIrStationGetAPIRequest {
	return &AlitripRailIrStationGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRailIrStationGetAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.ir.station.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRailIrStationGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAgentId is AgentId Setter
// 商家id
func (r *AlitripRailIrStationGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitripRailIrStationGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetPageIndex is PageIndex Setter
// 页数 从1开始
func (r *AlitripRailIrStationGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r AlitripRailIrStationGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AlitripRailIrStationGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlitripRailIrStationGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
