package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBanamadpcItemAddAPIRequest 新发商品 API请求
// taobao.banamadpc.item.add
//
// 巴拿马供应商通过此接口新发商品
type TaobaoBanamadpcItemAddAPIRequest struct {
	model.Params
	// 商品的schema xml
	_xml string
	// 类目id
	_catId int64
}

// NewTaobaoBanamadpcItemAddRequest 初始化TaobaoBanamadpcItemAddAPIRequest对象
func NewTaobaoBanamadpcItemAddRequest() *TaobaoBanamadpcItemAddAPIRequest {
	return &TaobaoBanamadpcItemAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBanamadpcItemAddAPIRequest) Reset() {
	r._xml = ""
	r._catId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBanamadpcItemAddAPIRequest) GetApiMethodName() string {
	return "taobao.banamadpc.item.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBanamadpcItemAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBanamadpcItemAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetXml is Xml Setter
// 商品的schema xml
func (r *TaobaoBanamadpcItemAddAPIRequest) SetXml(_xml string) error {
	r._xml = _xml
	r.Set("xml", _xml)
	return nil
}

// GetXml Xml Getter
func (r TaobaoBanamadpcItemAddAPIRequest) GetXml() string {
	return r._xml
}

// SetCatId is CatId Setter
// 类目id
func (r *TaobaoBanamadpcItemAddAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r TaobaoBanamadpcItemAddAPIRequest) GetCatId() int64 {
	return r._catId
}

var poolTaobaoBanamadpcItemAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBanamadpcItemAddRequest()
	},
}

// GetTaobaoBanamadpcItemAddRequest 从 sync.Pool 获取 TaobaoBanamadpcItemAddAPIRequest
func GetTaobaoBanamadpcItemAddAPIRequest() *TaobaoBanamadpcItemAddAPIRequest {
	return poolTaobaoBanamadpcItemAddAPIRequest.Get().(*TaobaoBanamadpcItemAddAPIRequest)
}

// ReleaseTaobaoBanamadpcItemAddAPIRequest 将 TaobaoBanamadpcItemAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoBanamadpcItemAddAPIRequest(v *TaobaoBanamadpcItemAddAPIRequest) {
	v.Reset()
	poolTaobaoBanamadpcItemAddAPIRequest.Put(v)
}
