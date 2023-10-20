package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailIrDivisionGetAPIRequest 国际火车票标准城市查询 API请求
// alitrip.rail.ir.division.get
//
// 国际火车票提供给代理商用于查询标准城市信息，全部城市数据量209530条，含除中国大陆以外的全部海外区域。
// 代理商通过分页查询的方式，拉取飞猪平台方全部境外标准城市，用于自身城市与飞猪平台城市的映射。
type AlitripRailIrDivisionGetAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
	// 层级，1洲，2是国家，3是省，4是市，5是区，6是街道/镇，7是村，8是逻辑行政区，境外火车票业务只需要市级别，传4就可以
	_level int64
	// 每页条数
	_pageSize int64
	// 页数，从1开始
	_pageIndex int64
}

// NewAlitripRailIrDivisionGetRequest 初始化AlitripRailIrDivisionGetAPIRequest对象
func NewAlitripRailIrDivisionGetRequest() *AlitripRailIrDivisionGetAPIRequest {
	return &AlitripRailIrDivisionGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRailIrDivisionGetAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.ir.division.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRailIrDivisionGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripRailIrDivisionGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *AlitripRailIrDivisionGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitripRailIrDivisionGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetLevel is Level Setter
// 层级，1洲，2是国家，3是省，4是市，5是区，6是街道/镇，7是村，8是逻辑行政区，境外火车票业务只需要市级别，传4就可以
func (r *AlitripRailIrDivisionGetAPIRequest) SetLevel(_level int64) error {
	r._level = _level
	r.Set("level", _level)
	return nil
}

// GetLevel Level Getter
func (r AlitripRailIrDivisionGetAPIRequest) GetLevel() int64 {
	return r._level
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AlitripRailIrDivisionGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlitripRailIrDivisionGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// 页数，从1开始
func (r *AlitripRailIrDivisionGetAPIRequest) SetPageIndex(_pageIndex int64) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r AlitripRailIrDivisionGetAPIRequest) GetPageIndex() int64 {
	return r._pageIndex
}
