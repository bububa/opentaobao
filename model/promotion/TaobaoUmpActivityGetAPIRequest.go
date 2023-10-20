package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityGetAPIRequest 查询营销活动 API请求
// taobao.ump.activity.get
//
// 查询营销活动
type TaobaoUmpActivityGetAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
}

// NewTaobaoUmpActivityGetRequest 初始化TaobaoUmpActivityGetAPIRequest对象
func NewTaobaoUmpActivityGetRequest() *TaobaoUmpActivityGetAPIRequest {
	return &TaobaoUmpActivityGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpActivityGetAPIRequest) Reset() {
	r._actId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpActivityGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activity.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpActivityGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpActivityGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoUmpActivityGetAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpActivityGetAPIRequest) GetActId() int64 {
	return r._actId
}

var poolTaobaoUmpActivityGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpActivityGetRequest()
	},
}

// GetTaobaoUmpActivityGetRequest 从 sync.Pool 获取 TaobaoUmpActivityGetAPIRequest
func GetTaobaoUmpActivityGetAPIRequest() *TaobaoUmpActivityGetAPIRequest {
	return poolTaobaoUmpActivityGetAPIRequest.Get().(*TaobaoUmpActivityGetAPIRequest)
}

// ReleaseTaobaoUmpActivityGetAPIRequest 将 TaobaoUmpActivityGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpActivityGetAPIRequest(v *TaobaoUmpActivityGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpActivityGetAPIRequest.Put(v)
}
