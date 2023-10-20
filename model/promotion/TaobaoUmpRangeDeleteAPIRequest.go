package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumprangedeleteAPIRequest 删除活动范围 API请求
// taobao.ump.range.delete
//
// 去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置
type TaobaoumprangedeleteAPIRequest struct {
	model.Params
	// id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id
	_ids string
	// 活动id
	_actId int64
	// 范围的类型，比如是全店，商品，类目见：MarketingConstants.PARTICIPATE_TYPE_*
	_type int64
}

// NewTaobaoumprangedeleteRequest 初始化TaobaoumprangedeleteAPIRequest对象
func NewTaobaoumprangedeleteRequest() *TaobaoumprangedeleteAPIRequest {
	return &TaobaoumprangedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoumprangedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ump.range.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoumprangedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoumprangedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id
func (r *TaobaoumprangedeleteAPIRequest) SetIds(_ids string) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaoumprangedeleteAPIRequest) GetIds() string {
	return r._ids
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoumprangedeleteAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoumprangedeleteAPIRequest) GetActId() int64 {
	return r._actId
}

// SetType is Type Setter
// 范围的类型，比如是全店，商品，类目见：MarketingConstants.PARTICIPATE_TYPE_*
func (r *TaobaoumprangedeleteAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoumprangedeleteAPIRequest) GetType() int64 {
	return r._type
}
