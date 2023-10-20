package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardsigninAPIRequest 服务商确认工人签到成功 API请求
// tmall.servicecenter.workcard.signin
//
// 服务商确认工人签到成功。需要服务商自己保证工人是在现场服务中。否则虚假回传签到而引起的后续问题全部由服务商自己承担
type TmallservicecenterworkcardsigninAPIRequest struct {
	model.Params
	// 核销单外部id
	_outerId string
	// 图片地址回传集合
	_picUrls string
	// 工单id
	_workcardId int64
	// 核销单id
	_fulfilTaskId int64
}

// NewTmallservicecenterworkcardsigninRequest 初始化TmallservicecenterworkcardsigninAPIRequest对象
func NewTmallservicecenterworkcardsigninRequest() *TmallservicecenterworkcardsigninAPIRequest {
	return &TmallservicecenterworkcardsigninAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardsigninAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.signin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardsigninAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardsigninAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 核销单外部id
func (r *TmallservicecenterworkcardsigninAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallservicecenterworkcardsigninAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetPicUrls is PicUrls Setter
// 图片地址回传集合
func (r *TmallservicecenterworkcardsigninAPIRequest) SetPicUrls(_picUrls string) error {
	r._picUrls = _picUrls
	r.Set("pic_urls", _picUrls)
	return nil
}

// GetPicUrls PicUrls Getter
func (r TmallservicecenterworkcardsigninAPIRequest) GetPicUrls() string {
	return r._picUrls
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallservicecenterworkcardsigninAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterworkcardsigninAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetFulfilTaskId is FulfilTaskId Setter
// 核销单id
func (r *TmallservicecenterworkcardsigninAPIRequest) SetFulfilTaskId(_fulfilTaskId int64) error {
	r._fulfilTaskId = _fulfilTaskId
	r.Set("fulfil_task_id", _fulfilTaskId)
	return nil
}

// GetFulfilTaskId FulfilTaskId Getter
func (r TmallservicecenterworkcardsigninAPIRequest) GetFulfilTaskId() int64 {
	return r._fulfilTaskId
}
