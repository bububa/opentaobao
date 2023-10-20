package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestaradgroupaddAPIRequest (新)创建一个推广组 API请求
// taobao.simba.salestar.adgroup.add
//
// 创建一个推广组
type TaobaosimbasalestaradgroupaddAPIRequest struct {
	model.Params
	// 创意标题，最多20个汉字
	_title string
	// 创意图片地址，必须是商品的图片之一
	_imgUrl string
	// 推广计划Id
	_campaignId int64
	// 商品Id
	_itemId int64
}

// NewTaobaosimbasalestaradgroupaddRequest 初始化TaobaosimbasalestaradgroupaddAPIRequest对象
func NewTaobaosimbasalestaradgroupaddRequest() *TaobaosimbasalestaradgroupaddAPIRequest {
	return &TaobaosimbasalestaradgroupaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbasalestaradgroupaddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.adgroup.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbasalestaradgroupaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbasalestaradgroupaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitle is Title Setter
// 创意标题，最多20个汉字
func (r *TaobaosimbasalestaradgroupaddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaosimbasalestaradgroupaddAPIRequest) GetTitle() string {
	return r._title
}

// SetImgUrl is ImgUrl Setter
// 创意图片地址，必须是商品的图片之一
func (r *TaobaosimbasalestaradgroupaddAPIRequest) SetImgUrl(_imgUrl string) error {
	r._imgUrl = _imgUrl
	r.Set("img_url", _imgUrl)
	return nil
}

// GetImgUrl ImgUrl Getter
func (r TaobaosimbasalestaradgroupaddAPIRequest) GetImgUrl() string {
	return r._imgUrl
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaosimbasalestaradgroupaddAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaosimbasalestaradgroupaddAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetItemId is ItemId Setter
// 商品Id
func (r *TaobaosimbasalestaradgroupaddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaosimbasalestaradgroupaddAPIRequest) GetItemId() int64 {
	return r._itemId
}
