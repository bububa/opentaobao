package tmallsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardTpconfirmAPIRequest 确认服务完成 API请求
// tmall.servicecenter.workcard.tpconfirm
//
// 服务商确认服务完成
type TmallServicecenterWorkcardTpconfirmAPIRequest struct {
	model.Params
	// 服务商确认服务完成信息
	_tpConfirmRequest *TpConfirmRequest
}

// NewTmallServicecenterWorkcardTpconfirmRequest 初始化TmallServicecenterWorkcardTpconfirmAPIRequest对象
func NewTmallServicecenterWorkcardTpconfirmRequest() *TmallServicecenterWorkcardTpconfirmAPIRequest {
	return &TmallServicecenterWorkcardTpconfirmAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallServicecenterWorkcardTpconfirmAPIRequest) Reset() {
	r._tpConfirmRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardTpconfirmAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.tpconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardTpconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardTpconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTpConfirmRequest is TpConfirmRequest Setter
// 服务商确认服务完成信息
func (r *TmallServicecenterWorkcardTpconfirmAPIRequest) SetTpConfirmRequest(_tpConfirmRequest *TpConfirmRequest) error {
	r._tpConfirmRequest = _tpConfirmRequest
	r.Set("tp_confirm_request", _tpConfirmRequest)
	return nil
}

// GetTpConfirmRequest TpConfirmRequest Getter
func (r TmallServicecenterWorkcardTpconfirmAPIRequest) GetTpConfirmRequest() *TpConfirmRequest {
	return r._tpConfirmRequest
}

var poolTmallServicecenterWorkcardTpconfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallServicecenterWorkcardTpconfirmRequest()
	},
}

// GetTmallServicecenterWorkcardTpconfirmRequest 从 sync.Pool 获取 TmallServicecenterWorkcardTpconfirmAPIRequest
func GetTmallServicecenterWorkcardTpconfirmAPIRequest() *TmallServicecenterWorkcardTpconfirmAPIRequest {
	return poolTmallServicecenterWorkcardTpconfirmAPIRequest.Get().(*TmallServicecenterWorkcardTpconfirmAPIRequest)
}

// ReleaseTmallServicecenterWorkcardTpconfirmAPIRequest 将 TmallServicecenterWorkcardTpconfirmAPIRequest 放入 sync.Pool
func ReleaseTmallServicecenterWorkcardTpconfirmAPIRequest(v *TmallServicecenterWorkcardTpconfirmAPIRequest) {
	v.Reset()
	poolTmallServicecenterWorkcardTpconfirmAPIRequest.Put(v)
}
