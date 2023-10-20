package tmallservice

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaServicecenterFulfiltaskCreateAPIRequest 合单生成核销单 API请求
// alibaba.servicecenter.fulfiltask.create
//
// 服务对工单进行合单，合单的结果是生成核销单
type AlibabaServicecenterFulfiltaskCreateAPIRequest struct {
	model.Params
	// 工单id列表
	_workcardIds []int64
	// 外部单号
	_outerId string
}

// NewAlibabaServicecenterFulfiltaskCreateRequest 初始化AlibabaServicecenterFulfiltaskCreateAPIRequest对象
func NewAlibabaServicecenterFulfiltaskCreateRequest() *AlibabaServicecenterFulfiltaskCreateAPIRequest {
	return &AlibabaServicecenterFulfiltaskCreateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaServicecenterFulfiltaskCreateAPIRequest) Reset() {
	r._workcardIds = r._workcardIds[:0]
	r._outerId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaServicecenterFulfiltaskCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.servicecenter.fulfiltask.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaServicecenterFulfiltaskCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaServicecenterFulfiltaskCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkcardIds is WorkcardIds Setter
// 工单id列表
func (r *AlibabaServicecenterFulfiltaskCreateAPIRequest) SetWorkcardIds(_workcardIds []int64) error {
	r._workcardIds = _workcardIds
	r.Set("workcard_ids", _workcardIds)
	return nil
}

// GetWorkcardIds WorkcardIds Getter
func (r AlibabaServicecenterFulfiltaskCreateAPIRequest) GetWorkcardIds() []int64 {
	return r._workcardIds
}

// SetOuterId is OuterId Setter
// 外部单号
func (r *AlibabaServicecenterFulfiltaskCreateAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaServicecenterFulfiltaskCreateAPIRequest) GetOuterId() string {
	return r._outerId
}

var poolAlibabaServicecenterFulfiltaskCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaServicecenterFulfiltaskCreateRequest()
	},
}

// GetAlibabaServicecenterFulfiltaskCreateRequest 从 sync.Pool 获取 AlibabaServicecenterFulfiltaskCreateAPIRequest
func GetAlibabaServicecenterFulfiltaskCreateAPIRequest() *AlibabaServicecenterFulfiltaskCreateAPIRequest {
	return poolAlibabaServicecenterFulfiltaskCreateAPIRequest.Get().(*AlibabaServicecenterFulfiltaskCreateAPIRequest)
}

// ReleaseAlibabaServicecenterFulfiltaskCreateAPIRequest 将 AlibabaServicecenterFulfiltaskCreateAPIRequest 放入 sync.Pool
func ReleaseAlibabaServicecenterFulfiltaskCreateAPIRequest(v *AlibabaServicecenterFulfiltaskCreateAPIRequest) {
	v.Reset()
	poolAlibabaServicecenterFulfiltaskCreateAPIRequest.Put(v)
}
