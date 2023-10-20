package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBanamadpcItemSelectPropAPIRequest 获取子属性 API请求
// taobao.banamadpc.item.select.prop
//
// 巴拿马供应商通过此接口获取子属性
type TaobaoBanamadpcItemSelectPropAPIRequest struct {
	model.Params
	// 子属性的schema xml
	_xml string
	// 类目id
	_catId int64
	// 属性id
	_propId int64
}

// NewTaobaoBanamadpcItemSelectPropRequest 初始化TaobaoBanamadpcItemSelectPropAPIRequest对象
func NewTaobaoBanamadpcItemSelectPropRequest() *TaobaoBanamadpcItemSelectPropAPIRequest {
	return &TaobaoBanamadpcItemSelectPropAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBanamadpcItemSelectPropAPIRequest) Reset() {
	r._xml = ""
	r._catId = 0
	r._propId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemSelectPropAPIRequest) GetApiMethodName() string {
	return "taobao.banamadpc.item.select.prop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemSelectPropAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBanamadpcItemSelectPropAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXml is Xml Setter
// 子属性的schema xml
func (r *TaobaoBanamadpcItemSelectPropAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// GetXml Xml Getter
func (r TaobaoBanamadpcItemSelectPropAPIRequest) GetXml() string {
	return r._xml
}

// SetCatId is CatId Setter
// 类目id
func (r *TaobaoBanamadpcItemSelectPropAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaoBanamadpcItemSelectPropAPIRequest) GetCatId() int64 {
	return r._catId
}

// SetPropId is PropId Setter
// 属性id
func (r *TaobaoBanamadpcItemSelectPropAPIRequest) SetPropId(_propId int64) error {
	r._propId = _propId
	r.Set("prop_id", _propId)
	return nil
}

// GetPropId PropId Getter
func (r TaobaoBanamadpcItemSelectPropAPIRequest) GetPropId() int64 {
	return r._propId
}

var poolTaobaoBanamadpcItemSelectPropAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBanamadpcItemSelectPropRequest()
	},
}

// GetTaobaoBanamadpcItemSelectPropRequest 从 sync.Pool 获取 TaobaoBanamadpcItemSelectPropAPIRequest
func GetTaobaoBanamadpcItemSelectPropAPIRequest() *TaobaoBanamadpcItemSelectPropAPIRequest {
	return poolTaobaoBanamadpcItemSelectPropAPIRequest.Get().(*TaobaoBanamadpcItemSelectPropAPIRequest)
}

// ReleaseTaobaoBanamadpcItemSelectPropAPIRequest 将 TaobaoBanamadpcItemSelectPropAPIRequest 放入 sync.Pool
func ReleaseTaobaoBanamadpcItemSelectPropAPIRequest(v *TaobaoBanamadpcItemSelectPropAPIRequest) {
	v.Reset()
	poolTaobaoBanamadpcItemSelectPropAPIRequest.Put(v)
}
