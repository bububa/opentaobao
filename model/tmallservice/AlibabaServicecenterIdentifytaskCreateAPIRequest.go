package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterIdentifytaskCreateAPIRequest 创建核销单 API请求
// alibaba.servicecenter.identifytask.create
//
// 创建核销单
type AlibabaServicecenterIdentifytaskCreateAPIRequest struct {
	model.Params
	// 工单集合
	_workcardIds []int64
	// 核销单外部标识，服务商保证唯一。如果创建核销单时传入重复的outer_id，系统直接会返回服务商下该outer_id对应的核销单，不会重新创建新核销单。
	_outerId string
}

// NewAlibabaServicecenterIdentifytaskCreateRequest 初始化AlibabaServicecenterIdentifytaskCreateAPIRequest对象
func NewAlibabaServicecenterIdentifytaskCreateRequest() *AlibabaServicecenterIdentifytaskCreateAPIRequest {
	return &AlibabaServicecenterIdentifytaskCreateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaServicecenterIdentifytaskCreateAPIRequest) Reset() {
	r._workcardIds = r._workcardIds[:0]
	r._outerId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterIdentifytaskCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.identifytask.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterIdentifytaskCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaServicecenterIdentifytaskCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardIds is WorkcardIds Setter
// 工单集合
func (r *AlibabaServicecenterIdentifytaskCreateAPIRequest) SetWorkcardIds(_workcardIds []int64) error {
	r._workcardIds = _workcardIds
	r.Set("workcard_ids", _workcardIds)
	return nil
}

// GetWorkcardIds WorkcardIds Getter
func (r AlibabaServicecenterIdentifytaskCreateAPIRequest) GetWorkcardIds() []int64 {
	return r._workcardIds
}

// SetOuterId is OuterId Setter
// 核销单外部标识，服务商保证唯一。如果创建核销单时传入重复的outer_id，系统直接会返回服务商下该outer_id对应的核销单，不会重新创建新核销单。
func (r *AlibabaServicecenterIdentifytaskCreateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaServicecenterIdentifytaskCreateAPIRequest) GetOuterId() string {
	return r._outerId
}

var poolAlibabaServicecenterIdentifytaskCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaServicecenterIdentifytaskCreateRequest()
	},
}

// GetAlibabaServicecenterIdentifytaskCreateRequest 从 sync.Pool 获取 AlibabaServicecenterIdentifytaskCreateAPIRequest
func GetAlibabaServicecenterIdentifytaskCreateAPIRequest() *AlibabaServicecenterIdentifytaskCreateAPIRequest {
	return poolAlibabaServicecenterIdentifytaskCreateAPIRequest.Get().(*AlibabaServicecenterIdentifytaskCreateAPIRequest)
}

// ReleaseAlibabaServicecenterIdentifytaskCreateAPIRequest 将 AlibabaServicecenterIdentifytaskCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaServicecenterIdentifytaskCreateAPIRequest(v *AlibabaServicecenterIdentifytaskCreateAPIRequest) {
	v.Reset()
	poolAlibabaServicecenterIdentifytaskCreateAPIRequest.Put(v)
}
