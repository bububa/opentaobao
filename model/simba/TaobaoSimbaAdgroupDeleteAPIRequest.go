package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSimbaAdgroupDeleteAPIRequest
删除一个推广组 API请求
taobao.simba.adgroup.delete

删除一个推广组 */
type TaobaoSimbaAdgroupDeleteAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaAdgroupDeleteRequest 初始化TaobaoSimbaAdgroupDeleteAPIRequest对象
func NewTaobaoSimbaAdgroupDeleteRequest() *TaobaoSimbaAdgroupDeleteAPIRequest {
	return &TaobaoSimbaAdgroupDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupDeleteAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// Get Nick Getter
func (r TaobaoSimbaAdgroupDeleteAPIRequest) GetNick() string {
	return r._nick
}

// Set is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaAdgroupDeleteAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoSimbaAdgroupDeleteAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
