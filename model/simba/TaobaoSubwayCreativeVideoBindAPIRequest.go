package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwaycreativevideobindAPIRequest 绑定视频到创意上 API请求
// taobao.subway.creative.video.bind
//
// 将用户上传的视频绑定到指定的创意上
type TaobaosubwaycreativevideobindAPIRequest struct {
	model.Params
	// 淘宝用户昵称
	_nick string
	// 审核通过的视频（状态6）
	_videoId int64
	// 视频类型
	_sizeType int64
	// ItemId
	_itemId int64
	// 创意ID
	_creativeId int64
}

// NewTaobaosubwaycreativevideobindRequest 初始化TaobaosubwaycreativevideobindAPIRequest对象
func NewTaobaosubwaycreativevideobindRequest() *TaobaosubwaycreativevideobindAPIRequest {
	return &TaobaosubwaycreativevideobindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosubwaycreativevideobindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.creative.video.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosubwaycreativevideobindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosubwaycreativevideobindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 淘宝用户昵称
func (r *TaobaosubwaycreativevideobindAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosubwaycreativevideobindAPIRequest) GetNick() string {
	return r._nick
}

// SetVideoId is VideoId Setter
// 审核通过的视频（状态6）
func (r *TaobaosubwaycreativevideobindAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r TaobaosubwaycreativevideobindAPIRequest) GetVideoId() int64 {
	return r._videoId
}

// SetSizeType is SizeType Setter
// 视频类型
func (r *TaobaosubwaycreativevideobindAPIRequest) SetSizeType(_sizeType int64) error {
	r._sizeType = _sizeType
	r.Set("size_type", _sizeType)
	return nil
}

// GetSizeType SizeType Getter
func (r TaobaosubwaycreativevideobindAPIRequest) GetSizeType() int64 {
	return r._sizeType
}

// SetItemId is ItemId Setter
// ItemId
func (r *TaobaosubwaycreativevideobindAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaosubwaycreativevideobindAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCreativeId is CreativeId Setter
// 创意ID
func (r *TaobaosubwaycreativevideobindAPIRequest) SetCreativeId(_creativeId int64) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaosubwaycreativevideobindAPIRequest) GetCreativeId() int64 {
	return r._creativeId
}
