package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityDeleteAPIRequest 删除营销活动 API请求
// taobao.ump.activity.delete
//
// 删除营销活动。对应的活动详情等将会被全部删除。
type TaobaoUmpActivityDeleteAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// NewTaobaoUmpActivityDeleteRequest 初始化TaobaoUmpActivityDeleteAPIRequest对象
func NewTaobaoUmpActivityDeleteRequest() *TaobaoUmpActivityDeleteAPIRequest {
	return &TaobaoUmpActivityDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpActivityDeleteAPIRequest) Reset() {
	r._actId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivityDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivityDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpActivityDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoUmpActivityDeleteAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpActivityDeleteAPIRequest) GetActId() int64 {
	return r._actId
}

var poolTaobaoUmpActivityDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpActivityDeleteRequest()
	},
}

// GetTaobaoUmpActivityDeleteRequest 从 sync.Pool 获取 TaobaoUmpActivityDeleteAPIRequest
func GetTaobaoUmpActivityDeleteAPIRequest() *TaobaoUmpActivityDeleteAPIRequest {
	return poolTaobaoUmpActivityDeleteAPIRequest.Get().(*TaobaoUmpActivityDeleteAPIRequest)
}

// ReleaseTaobaoUmpActivityDeleteAPIRequest 将 TaobaoUmpActivityDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpActivityDeleteAPIRequest(v *TaobaoUmpActivityDeleteAPIRequest) {
	v.Reset()
	poolTaobaoUmpActivityDeleteAPIRequest.Put(v)
}
