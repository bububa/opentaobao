package travel

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelpoisearchAPIRequest POI信息查询 API请求
// alitrip.travel.poi.search
//
// POI信息查询，用于商品更新使用
type AlitriptravelpoisearchAPIRequest struct {
	model.Params
	// POI信息名称
	_name string
	// POI信息ID；
	_id int64
	// POI类型；1->机场,2->火车站,3->汽车站,4->酒店,5->景点,6->购物，7->美食，9->玩乐，10->阿里旅行服务站，11->服务，100->其他
	_type int64
}

// NewAlitriptravelpoisearchRequest 初始化AlitriptravelpoisearchAPIRequest对象
func NewAlitriptravelpoisearchRequest() *AlitriptravelpoisearchAPIRequest {
	return &AlitriptravelpoisearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelpoisearchAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.poi.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelpoisearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelpoisearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// POI信息名称
func (r *AlitriptravelpoisearchAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlitriptravelpoisearchAPIRequest) GetName() string {
	return r._name
}

// SetId is Id Setter
// POI信息ID；
func (r *AlitriptravelpoisearchAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r AlitriptravelpoisearchAPIRequest) GetId() int64 {
	return r._id
}

// SetType is Type Setter
// POI类型；1-&gt;机场,2-&gt;火车站,3-&gt;汽车站,4-&gt;酒店,5-&gt;景点,6-&gt;购物，7-&gt;美食，9-&gt;玩乐，10-&gt;阿里旅行服务站，11-&gt;服务，100-&gt;其他
func (r *AlitriptravelpoisearchAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r AlitriptravelpoisearchAPIRequest) GetType() int64 {
	return r._type
}
