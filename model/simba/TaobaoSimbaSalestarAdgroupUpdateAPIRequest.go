package simba

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.adgroup.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
