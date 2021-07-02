package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallTmjlappSapServiceorderCancelAPIRequest 取消售后服务单 API请求
// tmall.tmjlapp.sap.serviceorder.cancel
//
// SAP跟天猫精灵app接口对接，用户在app取消sap售后服务工单
type TmallTmjlappSapServiceorderCancelAPIRequest struct {
	model.Params
	// 取消服务单请求
	_cancelRequest *Dtcancelrequest
}

// NewTmallTmjlappSapServiceorderCancelRequest 初始化TmallTmjlappSapServiceorderCancelAPIRequest对象
func NewTmallTmjlappSapServiceorderCancelRequest() *TmallTmjlappSapServiceorderCancelAPIRequest {
	return &TmallTmjlappSapServiceorderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTmjlappSapServiceorderCancelAPIRequest) GetApiMethodName() string {
	return "tmall.tmjlapp.sap.serviceorder.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTmjlappSapServiceorderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CancelRequest Setter
// 取消服务单请求
func (r *TmallTmjlappSapServiceorderCancelAPIRequest) SetCancelRequest(_cancelRequest *Dtcancelrequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// Get CancelRequest Getter
func (r TmallTmjlappSapServiceorderCancelAPIRequest) GetCancelRequest() *Dtcancelrequest {
	return r._cancelRequest
}
