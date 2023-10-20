package iotticket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoIotTicketSpCommentAPIRequest IoT售后服务商工单备注 API请求
// cainiao.iot.ticket.sp.comment
//
// IoT售后服务商工单备注
type CainiaoIotTicketSpCommentAPIRequest struct {
	model.Params
	// 请求参数
	_param *CommentTicketTopRequest
}

// NewCainiaoIotTicketSpCommentRequest 初始化CainiaoIotTicketSpCommentAPIRequest对象
func NewCainiaoIotTicketSpCommentRequest() *CainiaoIotTicketSpCommentAPIRequest {
	return &CainiaoIotTicketSpCommentAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *CainiaoIotTicketSpCommentAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoIotTicketSpCommentAPIRequest) GetApiMethodName() string {
	return "cainiao.iot.ticket.sp.comment"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoIotTicketSpCommentAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaoIotTicketSpCommentAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求参数
func (r *CainiaoIotTicketSpCommentAPIRequest) SetParam(_param *CommentTicketTopRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r CainiaoIotTicketSpCommentAPIRequest) GetParam() *CommentTicketTopRequest {
	return r._param
}

var poolCainiaoIotTicketSpCommentAPIRequest = sync.Pool{
	New: func() any {
		return NewCainiaoIotTicketSpCommentRequest()
	},
}

// GetCainiaoIotTicketSpCommentRequest 从 sync.Pool 获取 CainiaoIotTicketSpCommentAPIRequest
func GetCainiaoIotTicketSpCommentAPIRequest() *CainiaoIotTicketSpCommentAPIRequest {
	return poolCainiaoIotTicketSpCommentAPIRequest.Get().(*CainiaoIotTicketSpCommentAPIRequest)
}

// ReleaseCainiaoIotTicketSpCommentAPIRequest 将 CainiaoIotTicketSpCommentAPIRequest 放入 sync.Pool
func ReleaseCainiaoIotTicketSpCommentAPIRequest(v *CainiaoIotTicketSpCommentAPIRequest) {
	v.Reset()
	poolCainiaoIotTicketSpCommentAPIRequest.Put(v)
}
