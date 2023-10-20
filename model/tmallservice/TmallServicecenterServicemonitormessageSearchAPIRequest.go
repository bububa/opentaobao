package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicemonitormessagesearchAPIRequest 根据时间段查询服务商的服务预警消息列表(15分钟内) API请求
// tmall.servicecenter.servicemonitormessage.search
//
// 根据时间段查询服务商的服务预警消息列表(15分钟内)
type TmallservicecenterservicemonitormessagesearchAPIRequest struct {
	model.Params
	// 开始时间long值
	_start int64
	// 结束时间long值，与start相差15分钟
	_end int64
}

// NewTmallservicecenterservicemonitormessagesearchRequest 初始化TmallservicecenterservicemonitormessagesearchAPIRequest对象
func NewTmallservicecenterservicemonitormessagesearchRequest() *TmallservicecenterservicemonitormessagesearchAPIRequest {
	return &TmallservicecenterservicemonitormessagesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterservicemonitormessagesearchAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicemonitormessage.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterservicemonitormessagesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterservicemonitormessagesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStart is Start Setter
// 开始时间long值
func (r *TmallservicecenterservicemonitormessagesearchAPIRequest) SetStart(_start int64) error {
	r._start = _start
	r.Set("start", _start)
	return nil
}

// GetStart Start Getter
func (r TmallservicecenterservicemonitormessagesearchAPIRequest) GetStart() int64 {
	return r._start
}

// SetEnd is End Setter
// 结束时间long值，与start相差15分钟
func (r *TmallservicecenterservicemonitormessagesearchAPIRequest) SetEnd(_end int64) error {
	r._end = _end
	r.Set("end", _end)
	return nil
}

// GetEnd End Getter
func (r TmallservicecenterservicemonitormessagesearchAPIRequest) GetEnd() int64 {
	return r._end
}
