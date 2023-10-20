package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumpactivitieslistgetAPIRequest 营销活动列表查询 API请求
// taobao.ump.activities.list.get
//
// 按照营销活动id的列表ids，查询对应的营销活动列表。
type TaobaoumpactivitieslistgetAPIRequest struct {
	model.Params
	// 营销活动id列表。
	_ids []int64
}

// NewTaobaoumpactivitieslistgetRequest 初始化TaobaoumpactivitieslistgetAPIRequest对象
func NewTaobaoumpactivitieslistgetRequest() *TaobaoumpactivitieslistgetAPIRequest {
	return &TaobaoumpactivitieslistgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumpactivitieslistgetAPIRequest) GetApiMethodName() string {
	return "taobao.ump.activities.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumpactivitieslistgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumpactivitieslistgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// 营销活动id列表。
func (r *TaobaoumpactivitieslistgetAPIRequest) SetIds(_ids []int64) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaoumpactivitieslistgetAPIRequest) GetIds() []int64 {
	return r._ids
}
