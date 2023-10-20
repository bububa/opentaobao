package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbaadgroupsitemexistAPIRequest 商品是否推广 API请求
// taobao.simba.adgroups.item.exist
//
// 判断在一个推广计划中是否已经推广了一个商品
type TaobaosimbaadgroupsitemexistAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 产品类型 101001005 代表普通推广，101001014代表销量明星
	_productId int64
	// 推广计划Id
	_campaignId int64
	// 商品Id
	_itemId int64
}

// NewTaobaosimbaadgroupsitemexistRequest 初始化TaobaosimbaadgroupsitemexistAPIRequest对象
func NewTaobaosimbaadgroupsitemexistRequest() *TaobaosimbaadgroupsitemexistAPIRequest {
	return &TaobaosimbaadgroupsitemexistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbaadgroupsitemexistAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroups.item.exist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbaadgroupsitemexistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbaadgroupsitemexistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbaadgroupsitemexistAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbaadgroupsitemexistAPIRequest) GetNick() string {
	return r._nick
}

// SetProductId is ProductId Setter
// 产品类型 101001005 代表普通推广，101001014代表销量明星
func (r *TaobaosimbaadgroupsitemexistAPIRequest) SetProductId(_productId int64) error {
	r._productId = _productId
	r.Set("product_id", _productId)
	return nil
}

// GetProductId ProductId Getter
func (r TaobaosimbaadgroupsitemexistAPIRequest) GetProductId() int64 {
	return r._productId
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbaadgroupsitemexistAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbaadgroupsitemexistAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetItemId is ItemId Setter
// 商品Id
func (r *TaobaosimbaadgroupsitemexistAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaosimbaadgroupsitemexistAPIRequest) GetItemId() int64 {
	return r._itemId
}
