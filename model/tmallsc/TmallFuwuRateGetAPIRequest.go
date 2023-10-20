package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallfuwurategetAPIRequest 服务商需获取到单条服务单评价信息 API请求
// tmall.fuwu.rate.get
//
// 服务商需获取到单条服务单评价信息
type TmallfuwurategetAPIRequest struct {
	model.Params
	// 请求
	_serviceRateMessageQueryCmd *ServiceRateMessageQueryCmd
}

// NewTmallfuwurategetRequest 初始化TmallfuwurategetAPIRequest对象
func NewTmallfuwurategetRequest() *TmallfuwurategetAPIRequest {
	return &TmallfuwurategetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallfuwurategetAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.rate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallfuwurategetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallfuwurategetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceRateMessageQueryCmd is ServiceRateMessageQueryCmd Setter
// 请求
func (r *TmallfuwurategetAPIRequest) SetServiceRateMessageQueryCmd(_serviceRateMessageQueryCmd *ServiceRateMessageQueryCmd) error {
	r._serviceRateMessageQueryCmd = _serviceRateMessageQueryCmd
	r.Set("service_rate_message_query_cmd", _serviceRateMessageQueryCmd)
	return nil
}

// GetServiceRateMessageQueryCmd ServiceRateMessageQueryCmd Getter
func (r TmallfuwurategetAPIRequest) GetServiceRateMessageQueryCmd() *ServiceRateMessageQueryCmd {
	return r._serviceRateMessageQueryCmd
}
