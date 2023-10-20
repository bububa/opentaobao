package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpMbbGetbyidAPIRequest 获取营销积木块 API请求
// taobao.ump.mbb.getbyid
//
// 根据积木块id获取营销积木块。
type TaobaoUmpMbbGetbyidAPIRequest struct {
	model.Params
	// 积木块的id
	_id int64
}

// NewTaobaoUmpMbbGetbyidRequest 初始化TaobaoUmpMbbGetbyidAPIRequest对象
func NewTaobaoUmpMbbGetbyidRequest() *TaobaoUmpMbbGetbyidAPIRequest {
	return &TaobaoUmpMbbGetbyidAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpMbbGetbyidAPIRequest) Reset() {
	r._id = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbGetbyidAPIRequest) GetApiMethodName() string {
	return "taobao.ump.mbb.getbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbGetbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpMbbGetbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 积木块的id
func (r *TaobaoUmpMbbGetbyidAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaoUmpMbbGetbyidAPIRequest) GetId() int64 {
	return r._id
}

var poolTaobaoUmpMbbGetbyidAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpMbbGetbyidRequest()
	},
}

// GetTaobaoUmpMbbGetbyidRequest 从 sync.Pool 获取 TaobaoUmpMbbGetbyidAPIRequest
func GetTaobaoUmpMbbGetbyidAPIRequest() *TaobaoUmpMbbGetbyidAPIRequest {
	return poolTaobaoUmpMbbGetbyidAPIRequest.Get().(*TaobaoUmpMbbGetbyidAPIRequest)
}

// ReleaseTaobaoUmpMbbGetbyidAPIRequest 将 TaobaoUmpMbbGetbyidAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpMbbGetbyidAPIRequest(v *TaobaoUmpMbbGetbyidAPIRequest) {
	v.Reset()
	poolTaobaoUmpMbbGetbyidAPIRequest.Put(v)
}
