package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarAdgroupAddAPIRequest (新)创建一个推广组 API请求
// taobao.simba.salestar.adgroup.add
//
// 创建一个推广组
type TaobaoSimbaSalestarAdgroupAddAPIRequest struct {
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

// NewTaobaoSimbaSalestarAdgroupAddRequest 初始化TaobaoSimbaSalestarAdgroupAddAPIRequest对象
func NewTaobaoSimbaSalestarAdgroupAddRequest() *TaobaoSimbaSalestarAdgroupAddAPIRequest {
	return &TaobaoSimbaSalestarAdgroupAddAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSimbaSalestarAdgroupAddAPIRequest) Reset() {
	r._title = ""
	r._imgUrl = ""
	r._campaignId = 0
	r._itemId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.adgroup.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitle is Title Setter
// 创意标题，最多20个汉字
func (r *TaobaoSimbaSalestarAdgroupAddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetTitle() string {
	return r._title
}

// SetImgUrl is ImgUrl Setter
// 创意图片地址，必须是商品的图片之一
func (r *TaobaoSimbaSalestarAdgroupAddAPIRequest) SetImgUrl(_imgUrl string) error {
	r._imgUrl = _imgUrl
	r.Set("img_url", _imgUrl)
	return nil
}

// GetImgUrl ImgUrl Getter
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetImgUrl() string {
	return r._imgUrl
}

// SetCampaignId is CampaignId Setter
// 推广计划Id
func (r *TaobaoSimbaSalestarAdgroupAddAPIRequest) SetCampaignId(_campaignId int64) error {
	r._campaignId = _campaignId
	r.Set("campaign_id", _campaignId)
	return nil
}

// GetCampaignId CampaignId Getter
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetCampaignId() int64 {
	return r._campaignId
}

// SetItemId is ItemId Setter
// 商品Id
func (r *TaobaoSimbaSalestarAdgroupAddAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoSimbaSalestarAdgroupAddAPIRequest) GetItemId() int64 {
	return r._itemId
}

var poolTaobaoSimbaSalestarAdgroupAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSimbaSalestarAdgroupAddRequest()
	},
}

// GetTaobaoSimbaSalestarAdgroupAddRequest 从 sync.Pool 获取 TaobaoSimbaSalestarAdgroupAddAPIRequest
func GetTaobaoSimbaSalestarAdgroupAddAPIRequest() *TaobaoSimbaSalestarAdgroupAddAPIRequest {
	return poolTaobaoSimbaSalestarAdgroupAddAPIRequest.Get().(*TaobaoSimbaSalestarAdgroupAddAPIRequest)
}

// ReleaseTaobaoSimbaSalestarAdgroupAddAPIRequest 将 TaobaoSimbaSalestarAdgroupAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoSimbaSalestarAdgroupAddAPIRequest(v *TaobaoSimbaSalestarAdgroupAddAPIRequest) {
	v.Reset()
	poolTaobaoSimbaSalestarAdgroupAddAPIRequest.Put(v)
}
