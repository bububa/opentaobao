package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpRangeAddAPIRequest 添加活动范围 API请求
// taobao.ump.range.add
//
// 指定某项活动中，某个商家的某些类型物品（指定商品或者别的）参加或者不参加活动。当活动详情的参与类型为部分或者部分不参加的时候，需要指定具体哪部分参加或者不参加，使用本接口完成操作。比如部分商品满就送，这里的range用来指定具体哪些商品参加满就送活动。
type TaobaoUmpRangeAddAPIRequest struct {
	model.Params
	// id列表，当范围类型为商品时，该id为商品id.多个id用逗号隔开，一次不超过50个
	_ids string
	// 活动id
	_actId int64
	// 范围的类型，比如是全店，商品，见：MarketingConstants.PARTICIPATE_TYPE_*
	_type int64
}

// NewTaobaoUmpRangeAddRequest 初始化TaobaoUmpRangeAddAPIRequest对象
func NewTaobaoUmpRangeAddRequest() *TaobaoUmpRangeAddAPIRequest {
	return &TaobaoUmpRangeAddAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoUmpRangeAddAPIRequest) Reset() {
	r._ids = ""
	r._actId = 0
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoUmpRangeAddAPIRequest) GetApiMethodName() string {
	return "taobao.ump.range.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoUmpRangeAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoUmpRangeAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIds is Ids Setter
// id列表，当范围类型为商品时，该id为商品id.多个id用逗号隔开，一次不超过50个
func (r *TaobaoUmpRangeAddAPIRequest) SetIds(_ids string) error {
	r._ids = _ids
	r.Set("ids", _ids)
	return nil
}

// GetIds Ids Getter
func (r TaobaoUmpRangeAddAPIRequest) GetIds() string {
	return r._ids
}

// SetActId is ActId Setter
// 活动id
func (r *TaobaoUmpRangeAddAPIRequest) SetActId(_actId int64) error {
	r._actId = _actId
	r.Set("act_id", _actId)
	return nil
}

// GetActId ActId Getter
func (r TaobaoUmpRangeAddAPIRequest) GetActId() int64 {
	return r._actId
}

// SetType is Type Setter
// 范围的类型，比如是全店，商品，见：MarketingConstants.PARTICIPATE_TYPE_*
func (r *TaobaoUmpRangeAddAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoUmpRangeAddAPIRequest) GetType() int64 {
	return r._type
}

var poolTaobaoUmpRangeAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoUmpRangeAddRequest()
	},
}

// GetTaobaoUmpRangeAddRequest 从 sync.Pool 获取 TaobaoUmpRangeAddAPIRequest
func GetTaobaoUmpRangeAddAPIRequest() *TaobaoUmpRangeAddAPIRequest {
	return poolTaobaoUmpRangeAddAPIRequest.Get().(*TaobaoUmpRangeAddAPIRequest)
}

// ReleaseTaobaoUmpRangeAddAPIRequest 将 TaobaoUmpRangeAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoUmpRangeAddAPIRequest(v *TaobaoUmpRangeAddAPIRequest) {
	v.Reset()
	poolTaobaoUmpRangeAddAPIRequest.Put(v)
}
