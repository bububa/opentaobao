package tmallgenie

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTmjlappSapServiceorderCancelAPIRequest) Reset() {
	r._cancelRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTmjlappSapServiceorderCancelAPIRequest) GetApiMethodName() string {
	return "tmall.tmjlapp.sap.serviceorder.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTmjlappSapServiceorderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTmjlappSapServiceorderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCancelRequest is CancelRequest Setter
// 取消服务单请求
func (r *TmallTmjlappSapServiceorderCancelAPIRequest) SetCancelRequest(_cancelRequest *Dtcancelrequest) error {
	r._cancelRequest = _cancelRequest
	r.Set("cancel_request", _cancelRequest)
	return nil
}

// GetCancelRequest CancelRequest Getter
func (r TmallTmjlappSapServiceorderCancelAPIRequest) GetCancelRequest() *Dtcancelrequest {
	return r._cancelRequest
}

var poolTmallTmjlappSapServiceorderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTmjlappSapServiceorderCancelRequest()
	},
}

// GetTmallTmjlappSapServiceorderCancelRequest 从 sync.Pool 获取 TmallTmjlappSapServiceorderCancelAPIRequest
func GetTmallTmjlappSapServiceorderCancelAPIRequest() *TmallTmjlappSapServiceorderCancelAPIRequest {
	return poolTmallTmjlappSapServiceorderCancelAPIRequest.Get().(*TmallTmjlappSapServiceorderCancelAPIRequest)
}

// ReleaseTmallTmjlappSapServiceorderCancelAPIRequest 将 TmallTmjlappSapServiceorderCancelAPIRequest 放入 sync.Pool
func ReleaseTmallTmjlappSapServiceorderCancelAPIRequest(v *TmallTmjlappSapServiceorderCancelAPIRequest) {
	v.Reset()
	poolTmallTmjlappSapServiceorderCancelAPIRequest.Put(v)
}
