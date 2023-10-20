package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMcnShopcatsListGetAPIRequest 店铺类目清单 API请求
// taobao.mcn.shopcats.list.get
//
// 无需授权; 获取前台展示的店铺类目;
type TaobaoMcnShopcatsListGetAPIRequest struct {
	model.Params
	// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
	_fields []string
}

// NewTaobaoMcnShopcatsListGetRequest 初始化TaobaoMcnShopcatsListGetAPIRequest对象
func NewTaobaoMcnShopcatsListGetRequest() *TaobaoMcnShopcatsListGetAPIRequest {
	return &TaobaoMcnShopcatsListGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoMcnShopcatsListGetAPIRequest) Reset() {
	r._fields = r._fields[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMcnShopcatsListGetAPIRequest) GetApiMethodName() string {
	return "taobao.mcn.shopcats.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMcnShopcatsListGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMcnShopcatsListGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 需要返回的字段列表，见ShopCat，默认返回：cid,parent_cid,name,is_parent
func (r *TaobaoMcnShopcatsListGetAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TaobaoMcnShopcatsListGetAPIRequest) GetFields() []string {
	return r._fields
}

var poolTaobaoMcnShopcatsListGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoMcnShopcatsListGetRequest()
	},
}

// GetTaobaoMcnShopcatsListGetRequest 从 sync.Pool 获取 TaobaoMcnShopcatsListGetAPIRequest
func GetTaobaoMcnShopcatsListGetAPIRequest() *TaobaoMcnShopcatsListGetAPIRequest {
	return poolTaobaoMcnShopcatsListGetAPIRequest.Get().(*TaobaoMcnShopcatsListGetAPIRequest)
}

// ReleaseTaobaoMcnShopcatsListGetAPIRequest 将 TaobaoMcnShopcatsListGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoMcnShopcatsListGetAPIRequest(v *TaobaoMcnShopcatsListGetAPIRequest) {
	v.Reset()
	poolTaobaoMcnShopcatsListGetAPIRequest.Put(v)
}
