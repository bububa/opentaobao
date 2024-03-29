package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupAddAPIRequest 创建一个推广组 API请求
// taobao.simba.adgroup.add
//
// 创建一个推广组
type TaobaoSimbaAdgroupAddAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 创意标题，最多20个汉字
	_title string
	// 创意图片地址，必须是商品的图片之一
	_imgUrl string
	// 推广计划Id
	_campaignId int64
	// 商品Id
	_itemId int64
	// 推广组默认出价，单位为分，不能小于5 不能大于日最高限额
	_defaultPrice int64
}

// NewTaobaoSimbaAdgroupAddRequest 初始化TaobaoSimbaAdgroupAddAPIRequest对象
func NewTaobaoSimbaAdgroupAddRequest() *TaobaoSimbaAdgroupAddAPIRequest {
	return &TaobaoSimbaAdgroupAddAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaAdgroupAddAPIRequest) Reset() {
	r._nick = ""
	r._title = ""
	r._imgUrl = ""
	r._campaignId = 0
	r._itemId = 0
	r._defaultPrice = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaAdgroupAddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.adgroup.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaAdgroupAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaAdgroupAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaAdgroupAddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaAdgroupAddAPIRequest) GetNick() string {
	return r._nick
}

// SetTitle is Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaAdgroupAddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoSimbaAdgroupAddAPIRequest) GetTitle() string {
	return r._title
}

// SetImgUrl is ImgUrl Setter
// 创意图片地址，必须是商品的图片之一
func (r *TaobaoSimbaAdgroupAddAPIRequest) SetImgUrl(_imgUrl string) error {
	r._imgUrl = _imgUrl
	r.Set("img_url", _imgUrl)
	return nil
}

// GetImgUrl ImgUrl Getter
func (r TaobaoSimbaAdgroupAddAPIRequest) GetImgUrl() string {
	return r._imgUrl
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaAdgroupAddAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaAdgroupAddAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetItemId is ItemId Setter
// 商品Id
func (r *TaobaoSimbaAdgroupAddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoSimbaAdgroupAddAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetDefaultPrice is DefaultPrice Setter
// 推广组默认出价，单位为分，不能小于5 不能大于日最高限额
func (r *TaobaoSimbaAdgroupAddAPIRequest) SetDefaultPrice(_defaultPrice int64) error {
	r._defaultPrice = _defaultPrice
	r.Set("default_price", _defaultPrice)
	return nil
}

// GetDefaultPrice DefaultPrice Getter
func (r TaobaoSimbaAdgroupAddAPIRequest) GetDefaultPrice() int64 {
	return r._defaultPrice
}

var poolTaobaoSimbaAdgroupAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaAdgroupAddRequest()
	},
}

// GetTaobaoSimbaAdgroupAddRequest 从 sync.Pool 获取 TaobaoSimbaAdgroupAddAPIRequest
func GetTaobaoSimbaAdgroupAddAPIRequest() *TaobaoSimbaAdgroupAddAPIRequest {
	return poolTaobaoSimbaAdgroupAddAPIRequest.Get().(*TaobaoSimbaAdgroupAddAPIRequest)
}

// ReleaseTaobaoSimbaAdgroupAddAPIRequest 将 TaobaoSimbaAdgroupAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaAdgroupAddAPIRequest(v *TaobaoSimbaAdgroupAddAPIRequest) {
	v.Reset()
	poolTaobaoSimbaAdgroupAddAPIRequest.Put(v)
}
