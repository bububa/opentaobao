package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentordercheckstatusuploadAPIRequest 上传验收结果 API请求
// alibaba.idle.rent.order.checkstatus.upload
//
// 上传验收结果
type AlibabaidlerentordercheckstatusuploadAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 校验结果
	_checkResult *CheckResultDto
}

// NewAlibabaidlerentordercheckstatusuploadRequest 初始化AlibabaidlerentordercheckstatusuploadAPIRequest对象
func NewAlibabaidlerentordercheckstatusuploadRequest() *AlibabaidlerentordercheckstatusuploadAPIRequest {
	return &AlibabaidlerentordercheckstatusuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentordercheckstatusuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.checkstatus.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentordercheckstatusuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentordercheckstatusuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaidlerentordercheckstatusuploadAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaidlerentordercheckstatusuploadAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetCheckResult is CheckResult Setter
// 校验结果
func (r *AlibabaidlerentordercheckstatusuploadAPIRequest) SetCheckResult(_checkResult *CheckResultDto) error {
	r._checkResult = _checkResult
	r.Set("check_result", _checkResult)
	return nil
}

// GetCheckResult CheckResult Getter
func (r AlibabaidlerentordercheckstatusuploadAPIRequest) GetCheckResult() *CheckResultDto {
	return r._checkResult
}
