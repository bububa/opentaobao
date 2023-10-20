package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBanamadpcItemRenderAPIRequest 新发商品发布页 API请求
// taobao.banamadpc.item.render
//
// 巴拿马供应商通过此接口新发商品发布页
type TaobaoBanamadpcItemRenderAPIRequest struct {
	model.Params
	// 类目ID
	_catId int64
}

// NewTaobaoBanamadpcItemRenderRequest 初始化TaobaoBanamadpcItemRenderAPIRequest对象
func NewTaobaoBanamadpcItemRenderRequest() *TaobaoBanamadpcItemRenderAPIRequest {
	return &TaobaoBanamadpcItemRenderAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBanamadpcItemRenderAPIRequest) Reset() {
	r._catId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemRenderAPIRequest) GetApiMethodName() string {
	return "taobao.banamadpc.item.render"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemRenderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBanamadpcItemRenderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 类目ID
func (r *TaobaoBanamadpcItemRenderAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaoBanamadpcItemRenderAPIRequest) GetCatId() int64 {
	return r._catId
}

var poolTaobaoBanamadpcItemRenderAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBanamadpcItemRenderRequest()
	},
}

// GetTaobaoBanamadpcItemRenderRequest 从 sync.Pool 获取 TaobaoBanamadpcItemRenderAPIRequest
func GetTaobaoBanamadpcItemRenderAPIRequest() *TaobaoBanamadpcItemRenderAPIRequest {
	return poolTaobaoBanamadpcItemRenderAPIRequest.Get().(*TaobaoBanamadpcItemRenderAPIRequest)
}

// ReleaseTaobaoBanamadpcItemRenderAPIRequest 将 TaobaoBanamadpcItemRenderAPIRequest 放入 sync.Pool
func ReleaseTaobaoBanamadpcItemRenderAPIRequest(v *TaobaoBanamadpcItemRenderAPIRequest) {
	v.Reset()
	poolTaobaoBanamadpcItemRenderAPIRequest.Put(v)
}
