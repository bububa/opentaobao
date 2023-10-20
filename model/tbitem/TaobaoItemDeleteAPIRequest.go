package tbitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemDeleteAPIRequest 删除单条商品 API请求
// taobao.item.delete
//
// 删除单条商品
type TaobaoItemDeleteAPIRequest struct {
	model.Params
	// 商品数字ID，该参数必须
	_numIid int64
}

// NewTaobaoItemDeleteRequest 初始化TaobaoItemDeleteAPIRequest对象
func NewTaobaoItemDeleteRequest() *TaobaoItemDeleteAPIRequest {
	return &TaobaoItemDeleteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoItemDeleteAPIRequest) Reset() {
	r._numIid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoItemDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.item.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoItemDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoItemDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNumIid is NumIid Setter
// 商品数字ID，该参数必须
func (r *TaobaoItemDeleteAPIRequest) SetNumIid(_numIid int64) error {
	r._numIid = _numIid
	r.Set("num_iid", _numIid)
	return nil
}

// GetNumIid NumIid Getter
func (r TaobaoItemDeleteAPIRequest) GetNumIid() int64 {
	return r._numIid
}

var poolTaobaoItemDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoItemDeleteRequest()
	},
}

// GetTaobaoItemDeleteRequest 从 sync.Pool 获取 TaobaoItemDeleteAPIRequest
func GetTaobaoItemDeleteAPIRequest() *TaobaoItemDeleteAPIRequest {
	return poolTaobaoItemDeleteAPIRequest.Get().(*TaobaoItemDeleteAPIRequest)
}

// ReleaseTaobaoItemDeleteAPIRequest 将 TaobaoItemDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoItemDeleteAPIRequest(v *TaobaoItemDeleteAPIRequest) {
	v.Reset()
	poolTaobaoItemDeleteAPIRequest.Put(v)
}
