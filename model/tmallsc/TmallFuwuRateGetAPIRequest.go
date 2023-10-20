package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallFuwuRateGetAPIRequest 服务商需获取到单条服务单评价信息 API请求
// tmall.fuwu.rate.get
//
// 服务商需获取到单条服务单评价信息
type TmallFuwuRateGetAPIRequest struct {
	model.Params
	// 请求
	_serviceRateMessageQueryCmd *ServiceRateMessageQueryCmd
}

// NewTmallFuwuRateGetRequest 初始化TmallFuwuRateGetAPIRequest对象
func NewTmallFuwuRateGetRequest() *TmallFuwuRateGetAPIRequest {
	return &TmallFuwuRateGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallFuwuRateGetAPIRequest) GetApiMethodName() string {
	return "tmall.fuwu.rate.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallFuwuRateGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallFuwuRateGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetServiceRateMessageQueryCmd is ServiceRateMessageQueryCmd Setter
// 请求
func (r *TmallFuwuRateGetAPIRequest) SetServiceRateMessageQueryCmd(_serviceRateMessageQueryCmd *ServiceRateMessageQueryCmd) error {
	r._serviceRateMessageQueryCmd = _serviceRateMessageQueryCmd
	r.Set("service_rate_message_query_cmd", _serviceRateMessageQueryCmd)
	return nil
}

// GetServiceRateMessageQueryCmd ServiceRateMessageQueryCmd Getter
func (r TmallFuwuRateGetAPIRequest) GetServiceRateMessageQueryCmd() *ServiceRateMessageQueryCmd {
	return r._serviceRateMessageQueryCmd
}
