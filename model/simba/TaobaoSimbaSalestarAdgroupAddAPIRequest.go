package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarAdgroupAddAPIRequest (新)创建一个推广组 API请求
// taobao.simba.salestar.adgroup.add
//
// 创建一个推广组
type TaobaoSimbaSalestarAdgroupAddAPIRequest struct {
	model.Params
	// 推广计划Id
	_campaignId int64
	// 商品Id
	_itemId int64
	// 创意标题，最多20个汉字
	_title string
	// 创意图片地址，必须是商品的图片之一
	_imgUrl string
}

// NewTaobaoSimbaSalestarAdgroupAddRequest 初始化TaobaoSimbaSalestarAdgroupAddAPIRequest对象
func NewTaobaoSimbaSalestarAdgroupAddRequest() *TaobaoSimbaSalestarAdgroupAddAPIRequest {
	return &TaobaoSimbaSalestarAdgroupAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.adgroup.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaSalestarAdgroupAddAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// Get CampaignId Getter
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// Set is ItemId Setter
// 商品Id
func (r *TaobaoSimbaSalestarAdgroupAddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// Get ItemId Getter
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetItemId() int64 {
	return r._itemId
}

// Set is Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaSalestarAdgroupAddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// Get Title Getter
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetTitle() string {
	return r._title
}

// Set is ImgUrl Setter
// 创意图片地址，必须是商品的图片之一
func (r *TaobaoSimbaSalestarAdgroupAddAPIRequest) SetImgUrl(_imgUrl string) error {
	r._imgUrl = _imgUrl
	r.Set("img_url", _imgUrl)
	return nil
}

// Get ImgUrl Getter
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetImgUrl() string {
	return r._imgUrl
}
