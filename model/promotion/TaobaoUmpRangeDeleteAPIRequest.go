package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpRangeDeleteAPIRequest 删除活动范围 API请求
// taobao.ump.range.delete
//
// 去指先前指定在某项活动中，某些类型的物品参加或者不参加活动的设置
type TaobaoUmpRangeDeleteAPIRequest struct {
	model.Params
	// id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id
	_ids string
	// 活动id
	_actId int64
	// 范围的类型，比如是全店，商品，类目见：MarketingConstants.PARTICIPATE_TYPE_*
	_type int64
}

// NewTaobaoUmpRangeDeleteRequest 初始化TaobaoUmpRangeDeleteAPIRequest对象
func NewTaobaoUmpRangeDeleteRequest() *TaobaoUmpRangeDeleteAPIRequest {
	return &TaobaoUmpRangeDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpRangeDeleteAPIRequest) Reset() {
	r._ids = ""
	r._actId = 0
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpRangeDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ump.range.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpRangeDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpRangeDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// id列表，当范围类型为商品时，该id为商品id；当范围类型为类目时，该id为类目id
func (r *TaobaoUmpRangeDeleteAPIRequest) SetIds(_ids string) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaoUmpRangeDeleteAPIRequest) GetIds() string {
	return r._ids
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoUmpRangeDeleteAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpRangeDeleteAPIRequest) GetActId() int64 {
	return r._actId
}

// SetType is Type Setter
// 范围的类型，比如是全店，商品，类目见：MarketingConstants.PARTICIPATE_TYPE_*
func (r *TaobaoUmpRangeDeleteAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoUmpRangeDeleteAPIRequest) GetType() int64 {
	return r._type
}

var poolTaobaoUmpRangeDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpRangeDeleteRequest()
	},
}

// GetTaobaoUmpRangeDeleteRequest 从 sync.Pool 获取 TaobaoUmpRangeDeleteAPIRequest
func GetTaobaoUmpRangeDeleteAPIRequest() *TaobaoUmpRangeDeleteAPIRequest {
	return poolTaobaoUmpRangeDeleteAPIRequest.Get().(*TaobaoUmpRangeDeleteAPIRequest)
}

// ReleaseTaobaoUmpRangeDeleteAPIRequest 将 TaobaoUmpRangeDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpRangeDeleteAPIRequest(v *TaobaoUmpRangeDeleteAPIRequest) {
	v.Reset()
	poolTaobaoUmpRangeDeleteAPIRequest.Put(v)
}
