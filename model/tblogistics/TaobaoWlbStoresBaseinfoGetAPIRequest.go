package tblogistics

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbStoresBaseinfoGetAPIRequest 商家按照仓的类型获取仓库接口 API请求
// taobao.wlb.stores.baseinfo.get
//
// 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
type TaobaoWlbStoresBaseinfoGetAPIRequest struct {
	model.Params
	// 0.商家仓库.1.菜鸟仓库.2全部被划分的仓库
	_type int64
}

// NewTaobaoWlbStoresBaseinfoGetRequest 初始化TaobaoWlbStoresBaseinfoGetAPIRequest对象
func NewTaobaoWlbStoresBaseinfoGetRequest() *TaobaoWlbStoresBaseinfoGetAPIRequest {
	return &TaobaoWlbStoresBaseinfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoWlbStoresBaseinfoGetAPIRequest) Reset() {
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbStoresBaseinfoGetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.stores.baseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbStoresBaseinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbStoresBaseinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 0.商家仓库.1.菜鸟仓库.2全部被划分的仓库
func (r *TaobaoWlbStoresBaseinfoGetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoWlbStoresBaseinfoGetAPIRequest) GetType() int64 {
	return r._type
}

var poolTaobaoWlbStoresBaseinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoWlbStoresBaseinfoGetRequest()
	},
}

// GetTaobaoWlbStoresBaseinfoGetRequest 从 sync.Pool 获取 TaobaoWlbStoresBaseinfoGetAPIRequest
func GetTaobaoWlbStoresBaseinfoGetAPIRequest() *TaobaoWlbStoresBaseinfoGetAPIRequest {
	return poolTaobaoWlbStoresBaseinfoGetAPIRequest.Get().(*TaobaoWlbStoresBaseinfoGetAPIRequest)
}

// ReleaseTaobaoWlbStoresBaseinfoGetAPIRequest 将 TaobaoWlbStoresBaseinfoGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoWlbStoresBaseinfoGetAPIRequest(v *TaobaoWlbStoresBaseinfoGetAPIRequest) {
	v.Reset()
	poolTaobaoWlbStoresBaseinfoGetAPIRequest.Put(v)
}
