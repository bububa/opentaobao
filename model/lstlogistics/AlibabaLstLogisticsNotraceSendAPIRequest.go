package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstlogisticsnotracesendAPIRequest 供应商-异云-无需物流发货 API请求
// alibaba.lst.logistics.notrace.send
//
// 异地云仓的订单，使用无需物流的发货方式来变更订单发货状态
type AlibabalstlogisticsnotracesendAPIRequest struct {
	model.Params
	// 入参
	_param *SendDummyOrderParam
}

// NewAlibabalstlogisticsnotracesendRequest 初始化AlibabalstlogisticsnotracesendAPIRequest对象
func NewAlibabalstlogisticsnotracesendRequest() *AlibabalstlogisticsnotracesendAPIRequest {
	return &AlibabalstlogisticsnotracesendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstlogisticsnotracesendAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.notrace.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstlogisticsnotracesendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstlogisticsnotracesendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabalstlogisticsnotracesendAPIRequest) SetParam(_param *SendDummyOrderParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabalstlogisticsnotracesendAPIRequest) GetParam() *SendDummyOrderParam {
	return r._param
}
