package xhotelonlineorder

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelMemberrightUpdateAPIRequest 酒店会员权益更新操作 API请求
// taobao.xhotel.memberright.update
//
// 当用户在搜索酒店时，我们需要根据用户是否可享有某项权益来进行相应价格的展示或隐藏，因此我们在酒店搜索时就需要判断用户是否享有某项权益。而由于酒店搜索频率过高，为提高搜索性能并降低第三方接口压力，当用户在搜索酒店时，淘宝会通过读取淘宝本地缓存的用户相关权益信息来进行判断。为提高缓存的准确性，当第三方有用户相关权益有变化时，通过调用淘宝此接口来更新淘宝本地缓存。此接口需要采用Top方式调用。
type TaobaoXhotelMemberrightUpdateAPIRequest struct {
	model.Params
	// 淘宝用户id
	_taobaoUserId int64
	// 会员权益类型，1表示首住权益
	_rightType int64
	// 表示用户是否有对应的权益，取值范围true、false
	_hasRight bool
}

// NewTaobaoXhotelMemberrightUpdateRequest 初始化TaobaoXhotelMemberrightUpdateAPIRequest对象
func NewTaobaoXhotelMemberrightUpdateRequest() *TaobaoXhotelMemberrightUpdateAPIRequest {
	return &TaobaoXhotelMemberrightUpdateAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelMemberrightUpdateAPIRequest) Reset() {
	r._taobaoUserId = 0
	r._rightType = 0
	r._hasRight = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelMemberrightUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.memberright.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelMemberrightUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelMemberrightUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTaobaoUserId is TaobaoUserId Setter
// 淘宝用户id
func (r *TaobaoXhotelMemberrightUpdateAPIRequest) SetTaobaoUserId(_taobaoUserId int64) error {
	r._taobaoUserId = _taobaoUserId
	r.Set("taobao_user_id", _taobaoUserId)
	return nil
}

// GetTaobaoUserId TaobaoUserId Getter
func (r TaobaoXhotelMemberrightUpdateAPIRequest) GetTaobaoUserId() int64 {
	return r._taobaoUserId
}

// SetRightType is RightType Setter
// 会员权益类型，1表示首住权益
func (r *TaobaoXhotelMemberrightUpdateAPIRequest) SetRightType(_rightType int64) error {
	r._rightType = _rightType
	r.Set("right_type", _rightType)
	return nil
}

// GetRightType RightType Getter
func (r TaobaoXhotelMemberrightUpdateAPIRequest) GetRightType() int64 {
	return r._rightType
}

// SetHasRight is HasRight Setter
// 表示用户是否有对应的权益，取值范围true、false
func (r *TaobaoXhotelMemberrightUpdateAPIRequest) SetHasRight(_hasRight bool) error {
	r._hasRight = _hasRight
	r.Set("has_right", _hasRight)
	return nil
}

// GetHasRight HasRight Getter
func (r TaobaoXhotelMemberrightUpdateAPIRequest) GetHasRight() bool {
	return r._hasRight
}

var poolTaobaoXhotelMemberrightUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelMemberrightUpdateRequest()
	},
}

// GetTaobaoXhotelMemberrightUpdateRequest 从 sync.Pool 获取 TaobaoXhotelMemberrightUpdateAPIRequest
func GetTaobaoXhotelMemberrightUpdateAPIRequest() *TaobaoXhotelMemberrightUpdateAPIRequest {
	return poolTaobaoXhotelMemberrightUpdateAPIRequest.Get().(*TaobaoXhotelMemberrightUpdateAPIRequest)
}

// ReleaseTaobaoXhotelMemberrightUpdateAPIRequest 将 TaobaoXhotelMemberrightUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelMemberrightUpdateAPIRequest(v *TaobaoXhotelMemberrightUpdateAPIRequest) {
	v.Reset()
	poolTaobaoXhotelMemberrightUpdateAPIRequest.Put(v)
}
