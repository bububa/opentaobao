package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteridentifytaskcreateAPIRequest 服务商创建核销单 API请求
// tmall.servicecenter.identifytask.create
//
// 服务商调用该接口进行创建核销单操作
type TmallservicecenteridentifytaskcreateAPIRequest struct {
	model.Params
	// 工单列表
	_workcardIds []int64
	// 服务商自定义的外部核销单id
	_outerId string
	// 是否改派
	_reassign bool
}

// NewTmallservicecenteridentifytaskcreateRequest 初始化TmallservicecenteridentifytaskcreateAPIRequest对象
func NewTmallservicecenteridentifytaskcreateRequest() *TmallservicecenteridentifytaskcreateAPIRequest {
	return &TmallservicecenteridentifytaskcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteridentifytaskcreateAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.identifytask.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteridentifytaskcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteridentifytaskcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardIds is WorkcardIds Setter
// 工单列表
func (r *TmallservicecenteridentifytaskcreateAPIRequest) SetWorkcardIds(_workcardIds []int64) error {
	r._workcardIds = _workcardIds
	r.Set("workcard_ids", _workcardIds)
	return nil
}

// GetWorkcardIds WorkcardIds Getter
func (r TmallservicecenteridentifytaskcreateAPIRequest) GetWorkcardIds() []int64 {
	return r._workcardIds
}

// SetOuterId is OuterId Setter
// 服务商自定义的外部核销单id
func (r *TmallservicecenteridentifytaskcreateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r TmallservicecenteridentifytaskcreateAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetReassign is Reassign Setter
// 是否改派
func (r *TmallservicecenteridentifytaskcreateAPIRequest) SetReassign(_reassign bool) error {
	r._reassign = _reassign
	r.Set("reassign", _reassign)
	return nil
}

// GetReassign Reassign Getter
func (r TmallservicecenteridentifytaskcreateAPIRequest) GetReassign() bool {
	return r._reassign
}
