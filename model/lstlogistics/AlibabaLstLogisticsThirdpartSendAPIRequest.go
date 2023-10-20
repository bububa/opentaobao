package lstlogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstlogisticsthirdpartsendAPIRequest 供应商-异云-使用第三方物流发货 API请求
// alibaba.lst.logistics.thirdpart.send
//
// 异地云仓的订单，使用第三方物流的发货方式来变更订单发货状态
type AlibabalstlogisticsthirdpartsendAPIRequest struct {
	model.Params
	// 入参
	_param *SendOfflineOrderParam
}

// NewAlibabalstlogisticsthirdpartsendRequest 初始化AlibabalstlogisticsthirdpartsendAPIRequest对象
func NewAlibabalstlogisticsthirdpartsendRequest() *AlibabalstlogisticsthirdpartsendAPIRequest {
	return &AlibabalstlogisticsthirdpartsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstlogisticsthirdpartsendAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.logistics.thirdpart.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstlogisticsthirdpartsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstlogisticsthirdpartsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlibabalstlogisticsthirdpartsendAPIRequest) SetParam(_param *SendOfflineOrderParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabalstlogisticsthirdpartsendAPIRequest) GetParam() *SendOfflineOrderParam {
	return r._param
}
