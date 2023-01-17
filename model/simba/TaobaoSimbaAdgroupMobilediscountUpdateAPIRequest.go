package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest 对推广组进行单独移动溢价 API请求
// taobao.simba.adgroup.mobilediscount.update
//
// 对推广组进行单独移动溢价
type TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest struct {
	model.Params
	// 推广组id数组(推广组id集合元素个数在1-200个之间，推广组id需要在同一个推广计划中)
	_adgroupIds []string
	// 昵称
	_nick string
	// 折扣（折扣值在1-400之间）
	_mobileDiscount int64
}

// NewTaobaoSimbaAdgroupMobilediscountUpdateRequest 初始化TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest对象
func NewTaobaoSimbaAdgroupMobilediscountUpdateRequest() *TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest {
	return &TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.mobilediscount.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroupIds is AdgroupIds Setter
// 推广组id数组(推广组id集合元素个数在1-200个之间，推广组id需要在同一个推广计划中)
func (r *TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) SetAdgroupIds(_adgroupIds []string) error {
	r._adgroupIds = _adgroupIds
	r.Set("adgroup_ids", _adgroupIds)
	return nil
}

// GetAdgroupIds AdgroupIds Getter
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetAdgroupIds() []string {
	return r._adgroupIds
}

// SetNick is Nick Setter
// 昵称
func (r *TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetNick() string {
	return r._nick
}

// SetMobileDiscount is MobileDiscount Setter
// 折扣（折扣值在1-400之间）
func (r *TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) SetMobileDiscount(_mobileDiscount int64) error {
	r._mobileDiscount = _mobileDiscount
	r.Set("mobile_discount", _mobileDiscount)
	return nil
}

// GetMobileDiscount MobileDiscount Getter
func (r TaobaoSimbaAdgroupMobilediscountUpdateAPIRequest) GetMobileDiscount() int64 {
	return r._mobileDiscount
}
