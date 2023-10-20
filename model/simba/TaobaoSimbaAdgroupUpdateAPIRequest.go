package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupUpdateAPIRequest 更新一个推广组的信息 API请求
// taobao.simba.adgroup.update
//
// 更新一个推广组的信息，可以设置默认出价、是否上线、非搜索出价、非搜索是否使用默认出价
type TaobaoSimbaAdgroupUpdateAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 非搜索是否使用默认出价，false-不用；true-使用；默认为true;
	_useNonsearchDefaultPrice string
	// 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
	_onlineStatus string
	// 推广组Id
	_adgroupId int64
	// 默认出价，单位是分，不能小于5
	_defaultPrice int64
	// 非搜索出价，单位是分，不能小于5，如果use_nonseatch_default_price为使用默认出价，则此nonsearch_max_price字段传入的数据不起作用，商品将使用默认非搜索出价
	_nonsearchMaxPrice int64
}

// NewTaobaoSimbaAdgroupUpdateRequest 初始化TaobaoSimbaAdgroupUpdateAPIRequest对象
func NewTaobaoSimbaAdgroupUpdateRequest() *TaobaoSimbaAdgroupUpdateAPIRequest {
	return &TaobaoSimbaAdgroupUpdateAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) Reset() {
	r._nick = ""
	r._useNonsearchDefaultPrice = ""
	r._onlineStatus = ""
	r._adgroupId = 0
	r._defaultPrice = 0
	r._nonsearchMaxPrice = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetNick() string {
	return r._nick
}

// SetUseNonsearchDefaultPrice is UseNonsearchDefaultPrice Setter
// 非搜索是否使用默认出价，false-不用；true-使用；默认为true;
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetUseNonsearchDefaultPrice(_useNonsearchDefaultPrice string) error {
	r._useNonsearchDefaultPrice = _useNonsearchDefaultPrice
	r.Set("use_nonsearch_default_price", _useNonsearchDefaultPrice)
	return nil
}

// GetUseNonsearchDefaultPrice UseNonsearchDefaultPrice Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetUseNonsearchDefaultPrice() string {
	return r._useNonsearchDefaultPrice
}

// SetOnlineStatus is OnlineStatus Setter
// 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetOnlineStatus(_onlineStatus string) error {
	r._onlineStatus = _onlineStatus
	r.Set("online_status", _onlineStatus)
	return nil
}

// GetOnlineStatus OnlineStatus Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetOnlineStatus() string {
	return r._onlineStatus
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetDefaultPrice is DefaultPrice Setter
// 默认出价，单位是分，不能小于5
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetDefaultPrice(_defaultPrice int64) error {
	r._defaultPrice = _defaultPrice
	r.Set("default_price", _defaultPrice)
	return nil
}

// GetDefaultPrice DefaultPrice Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetDefaultPrice() int64 {
	return r._defaultPrice
}

// SetNonsearchMaxPrice is NonsearchMaxPrice Setter
// 非搜索出价，单位是分，不能小于5，如果use_nonseatch_default_price为使用默认出价，则此nonsearch_max_price字段传入的数据不起作用，商品将使用默认非搜索出价
func (r *TaobaoSimbaAdgroupUpdateAPIRequest) SetNonsearchMaxPrice(_nonsearchMaxPrice int64) error {
	r._nonsearchMaxPrice = _nonsearchMaxPrice
	r.Set("nonsearch_max_price", _nonsearchMaxPrice)
	return nil
}

// GetNonsearchMaxPrice NonsearchMaxPrice Getter
func (r TaobaoSimbaAdgroupUpdateAPIRequest) GetNonsearchMaxPrice() int64 {
	return r._nonsearchMaxPrice
}

var poolTaobaoSimbaAdgroupUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaAdgroupUpdateRequest()
	},
}

// GetTaobaoSimbaAdgroupUpdateRequest 从 sync.Pool 获取 TaobaoSimbaAdgroupUpdateAPIRequest
func GetTaobaoSimbaAdgroupUpdateAPIRequest() *TaobaoSimbaAdgroupUpdateAPIRequest {
	return poolTaobaoSimbaAdgroupUpdateAPIRequest.Get().(*TaobaoSimbaAdgroupUpdateAPIRequest)
}

// ReleaseTaobaoSimbaAdgroupUpdateAPIRequest 将 TaobaoSimbaAdgroupUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaAdgroupUpdateAPIRequest(v *TaobaoSimbaAdgroupUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSimbaAdgroupUpdateAPIRequest.Put(v)
}
