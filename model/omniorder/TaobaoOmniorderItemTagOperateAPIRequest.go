package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOmniorderItemTagOperateAPIRequest 全渠道商品打标与去标 API请求
// taobao.omniorder.item.tag.operate
//
// 用于对全渠道商品进行打标、去标（门店发货标，门店自提标，前置拆单标）操作。另外还包括增加、删除、修改分单系统，接单系统配置。
type TaobaoOmniorderItemTagOperateAPIRequest struct {
	model.Params
	// 商品标,storeDeliver代表门店发货, AllocateByFront代表前置拆单, storeCollect代表门店自提
	_types []string
	// 商品ID
	_itemId int64
	// 操作状态， 填 1 代表打标，填 -1 代表去标
	_status int64
	// 分单&接单设置
	_omniSetting *OmniSettingDto
}

// NewTaobaoOmniorderItemTagOperateRequest 初始化TaobaoOmniorderItemTagOperateAPIRequest对象
func NewTaobaoOmniorderItemTagOperateRequest() *TaobaoOmniorderItemTagOperateAPIRequest {
	return &TaobaoOmniorderItemTagOperateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderItemTagOperateAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.item.tag.operate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderItemTagOperateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTypes is Types Setter
// 商品标,storeDeliver代表门店发货, AllocateByFront代表前置拆单, storeCollect代表门店自提
func (r *TaobaoOmniorderItemTagOperateAPIRequest) SetTypes(_types []string) error {
	r._types = _types
	r.Set("types", _types)
	return nil
}

// GetTypes Types Getter
func (r TaobaoOmniorderItemTagOperateAPIRequest) GetTypes() []string {
	return r._types
}

// SetItemId is ItemId Setter
// 商品ID
func (r *TaobaoOmniorderItemTagOperateAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoOmniorderItemTagOperateAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetStatus is Status Setter
// 操作状态， 填 1 代表打标，填 -1 代表去标
func (r *TaobaoOmniorderItemTagOperateAPIRequest) SetStatus(_status int64) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoOmniorderItemTagOperateAPIRequest) GetStatus() int64 {
	return r._status
}

// SetOmniSetting is OmniSetting Setter
// 分单&amp;接单设置
func (r *TaobaoOmniorderItemTagOperateAPIRequest) SetOmniSetting(_omniSetting *OmniSettingDto) error {
	r._omniSetting = _omniSetting
	r.Set("omni_setting", _omniSetting)
	return nil
}

// GetOmniSetting OmniSetting Getter
func (r TaobaoOmniorderItemTagOperateAPIRequest) GetOmniSetting() *OmniSettingDto {
	return r._omniSetting
}
