package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellTicketingListAPIRequest 销售出票列表 API请求
// alitrip.agent.flight.sell.ticketing.list
//
// 销售出票列表
type AlitripAgentFlightSellTicketingListAPIRequest struct {
	model.Params
	// 入参对象
	_param *TicketingListRequestDto
}

// NewAlitripAgentFlightSellTicketingListRequest 初始化AlitripAgentFlightSellTicketingListAPIRequest对象
func NewAlitripAgentFlightSellTicketingListRequest() *AlitripAgentFlightSellTicketingListAPIRequest {
	return &AlitripAgentFlightSellTicketingListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightSellTicketingListAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellTicketingListAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.ticketing.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellTicketingListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightSellTicketingListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlitripAgentFlightSellTicketingListAPIRequest) SetParam(_param *TicketingListRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripAgentFlightSellTicketingListAPIRequest) GetParam() *TicketingListRequestDto {
	return r._param
}

var poolAlitripAgentFlightSellTicketingListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightSellTicketingListRequest()
	},
}

// GetAlitripAgentFlightSellTicketingListRequest 从 sync.Pool 获取 AlitripAgentFlightSellTicketingListAPIRequest
func GetAlitripAgentFlightSellTicketingListAPIRequest() *AlitripAgentFlightSellTicketingListAPIRequest {
	return poolAlitripAgentFlightSellTicketingListAPIRequest.Get().(*AlitripAgentFlightSellTicketingListAPIRequest)
}

// ReleaseAlitripAgentFlightSellTicketingListAPIRequest 将 AlitripAgentFlightSellTicketingListAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightSellTicketingListAPIRequest(v *AlitripAgentFlightSellTicketingListAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightSellTicketingListAPIRequest.Put(v)
}
