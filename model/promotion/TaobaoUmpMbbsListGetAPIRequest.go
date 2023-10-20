package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpMbbsListGetAPIRequest 通过ids列表获取营销积木块列表 API请求
// taobao.ump.mbbs.list.get
//
// 通过营销积木id列表来获取营销积木块列表。
type TaobaoUmpMbbsListGetAPIRequest struct {
	model.Params
	// 营销积木块id组成的字符串。
	_ids []int64
}

// NewTaobaoUmpMbbsListGetRequest 初始化TaobaoUmpMbbsListGetAPIRequest对象
func NewTaobaoUmpMbbsListGetRequest() *TaobaoUmpMbbsListGetAPIRequest {
	return &TaobaoUmpMbbsListGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpMbbsListGetAPIRequest) Reset() {
	r._ids = r._ids[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpMbbsListGetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.mbbs.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpMbbsListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpMbbsListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// 营销积木块id组成的字符串。
func (r *TaobaoUmpMbbsListGetAPIRequest) SetIds(_ids []int64) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaoUmpMbbsListGetAPIRequest) GetIds() []int64 {
	return r._ids
}

var poolTaobaoUmpMbbsListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpMbbsListGetRequest()
	},
}

// GetTaobaoUmpMbbsListGetRequest 从 sync.Pool 获取 TaobaoUmpMbbsListGetAPIRequest
func GetTaobaoUmpMbbsListGetAPIRequest() *TaobaoUmpMbbsListGetAPIRequest {
	return poolTaobaoUmpMbbsListGetAPIRequest.Get().(*TaobaoUmpMbbsListGetAPIRequest)
}

// ReleaseTaobaoUmpMbbsListGetAPIRequest 将 TaobaoUmpMbbsListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpMbbsListGetAPIRequest(v *TaobaoUmpMbbsListGetAPIRequest) {
	v.Reset()
	poolTaobaoUmpMbbsListGetAPIRequest.Put(v)
}
