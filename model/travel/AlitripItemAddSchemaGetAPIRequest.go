package travel

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripItemAddSchemaGetAPIRequest 获取商品发布模板 API请求
// alitrip.item.add.schema.get
//
// 发布飞猪度假商品时，需要先调用此接口获取商品发布的模板schema。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripItemAddSchemaGetAPIRequest struct {
	model.Params
	// 类目id
	_catId int64
}

// NewAlitripItemAddSchemaGetRequest 初始化AlitripItemAddSchemaGetAPIRequest对象
func NewAlitripItemAddSchemaGetRequest() *AlitripItemAddSchemaGetAPIRequest {
	return &AlitripItemAddSchemaGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripItemAddSchemaGetAPIRequest) Reset() {
	r._catId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripItemAddSchemaGetAPIRequest) GetApiMethodName() string {
	return "alitrip.item.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripItemAddSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripItemAddSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 类目id
func (r *AlitripItemAddSchemaGetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlitripItemAddSchemaGetAPIRequest) GetCatId() int64 {
	return r._catId
}

var poolAlitripItemAddSchemaGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripItemAddSchemaGetRequest()
	},
}

// GetAlitripItemAddSchemaGetRequest 从 sync.Pool 获取 AlitripItemAddSchemaGetAPIRequest
func GetAlitripItemAddSchemaGetAPIRequest() *AlitripItemAddSchemaGetAPIRequest {
	return poolAlitripItemAddSchemaGetAPIRequest.Get().(*AlitripItemAddSchemaGetAPIRequest)
}

// ReleaseAlitripItemAddSchemaGetAPIRequest 将 AlitripItemAddSchemaGetAPIRequest 放入 sync.Pool
func ReleaseAlitripItemAddSchemaGetAPIRequest(v *AlitripItemAddSchemaGetAPIRequest) {
	v.Reset()
	poolAlitripItemAddSchemaGetAPIRequest.Put(v)
}
