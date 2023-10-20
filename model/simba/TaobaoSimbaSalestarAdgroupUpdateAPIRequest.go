package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarAdgroupUpdateAPIRequest 销量明星更新一个推广组的信息 API请求
// taobao.simba.salestar.adgroup.update
//
// 更新一个推广组的信息，可以设置 是否上线
type TaobaoSimbaSalestarAdgroupUpdateAPIRequest struct {
	model.Params
	// 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
	_onlineStatus string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaSalestarAdgroupUpdateRequest 初始化TaobaoSimbaSalestarAdgroupUpdateAPIRequest对象
func NewTaobaoSimbaSalestarAdgroupUpdateRequest() *TaobaoSimbaSalestarAdgroupUpdateAPIRequest {
	return &TaobaoSimbaSalestarAdgroupUpdateAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSalestarAdgroupUpdateAPIRequest) Reset() {
	r._onlineStatus = ""
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.adgroup.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSalestarAdgroupUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOnlineStatus is OnlineStatus Setter
// 用户设置的上下线状态 offline-下线(暂停竞价)； online-上线；默认为online
func (r *TaobaoSimbaSalestarAdgroupUpdateAPIRequest) SetOnlineStatus(_onlineStatus string) error {
	r._onlineStatus = _onlineStatus
	r.Set("online_status", _onlineStatus)
	return nil
}

// GetOnlineStatus OnlineStatus Getter
func (r TaobaoSimbaSalestarAdgroupUpdateAPIRequest) GetOnlineStatus() string {
	return r._onlineStatus
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaSalestarAdgroupUpdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaSalestarAdgroupUpdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoSimbaSalestarAdgroupUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSalestarAdgroupUpdateRequest()
	},
}

// GetTaobaoSimbaSalestarAdgroupUpdateRequest 从 sync.Pool 获取 TaobaoSimbaSalestarAdgroupUpdateAPIRequest
func GetTaobaoSimbaSalestarAdgroupUpdateAPIRequest() *TaobaoSimbaSalestarAdgroupUpdateAPIRequest {
	return poolTaobaoSimbaSalestarAdgroupUpdateAPIRequest.Get().(*TaobaoSimbaSalestarAdgroupUpdateAPIRequest)
}

// ReleaseTaobaoSimbaSalestarAdgroupUpdateAPIRequest 将 TaobaoSimbaSalestarAdgroupUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSalestarAdgroupUpdateAPIRequest(v *TaobaoSimbaSalestarAdgroupUpdateAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSalestarAdgroupUpdateAPIRequest.Put(v)
}
