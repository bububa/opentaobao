package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripitemaddschemagetAPIRequest 获取商品发布模板 API请求
// alitrip.item.add.schema.get
//
// 发布飞猪度假商品时，需要先调用此接口获取商品发布的模板schema。目前支持类目：出境自由行(50278002)、境内自由行(50272002)、出境跟团游(50258005)、境内跟团游(50258004)、境外一日游/多日游(50276003)
type AlitripitemaddschemagetAPIRequest struct {
	model.Params
	// 类目id
	_catId int64
}

// NewAlitripitemaddschemagetRequest 初始化AlitripitemaddschemagetAPIRequest对象
func NewAlitripitemaddschemagetRequest() *AlitripitemaddschemagetAPIRequest {
	return &AlitripitemaddschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripitemaddschemagetAPIRequest) GetApiMethodName() string {
	return "alitrip.item.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripitemaddschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripitemaddschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 类目id
func (r *AlitripitemaddschemagetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlitripitemaddschemagetAPIRequest) GetCatId() int64 {
	return r._catId
}
