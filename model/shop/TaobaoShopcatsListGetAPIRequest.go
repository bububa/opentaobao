package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoShopcatsListGetAPIRequest 获取前台展示的店铺类目 API请求
// taobao.shopcats.list.get
//
// 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
type TaobaoShopcatsListGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
	_fields []string
}

// NewTaobaoShopcatsListGetRequest 初始化TaobaoShopcatsListGetAPIRequest对象
func NewTaobaoShopcatsListGetRequest() *TaobaoShopcatsListGetAPIRequest {
	return &TaobaoShopcatsListGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoShopcatsListGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoShopcatsListGetAPIRequest) GetApiMethodName() string {
	return "taobao.shopcats.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoShopcatsListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoShopcatsListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
func (r *TaobaoShopcatsListGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoShopcatsListGetAPIRequest) GetFields() []string {
	return r._fields
}

var poolTaobaoShopcatsListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoShopcatsListGetRequest()
	},
}

// GetTaobaoShopcatsListGetRequest 从 sync.Pool 获取 TaobaoShopcatsListGetAPIRequest
func GetTaobaoShopcatsListGetAPIRequest() *TaobaoShopcatsListGetAPIRequest {
	return poolTaobaoShopcatsListGetAPIRequest.Get().(*TaobaoShopcatsListGetAPIRequest)
}

// ReleaseTaobaoShopcatsListGetAPIRequest 将 TaobaoShopcatsListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoShopcatsListGetAPIRequest(v *TaobaoShopcatsListGetAPIRequest) {
	v.Reset()
	poolTaobaoShopcatsListGetAPIRequest.Put(v)
}
