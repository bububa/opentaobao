package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCreativeVideoBindAPIRequest 绑定视频到创意上 API请求
// taobao.subway.creative.video.bind
//
// 将用户上传的视频绑定到指定的创意上
type TaobaoSubwayCreativeVideoBindAPIRequest struct {
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

// NewTaobaoSubwayCreativeVideoBindRequest 初始化TaobaoSubwayCreativeVideoBindAPIRequest对象
func NewTaobaoSubwayCreativeVideoBindRequest() *TaobaoSubwayCreativeVideoBindAPIRequest {
	return &TaobaoSubwayCreativeVideoBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayCreativeVideoBindAPIRequest) GetApiMethodName() string {
	return "taobao.subway.creative.video.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayCreativeVideoBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNick is Nick Setter
// 淘宝用户昵称
func (r *TaobaoSubwayCreativeVideoBindAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayCreativeVideoBindAPIRequest) GetNick() string {
	return r._nick
}

// SetVideoId is VideoId Setter
// 审核通过的视频（状态6）
func (r *TaobaoSubwayCreativeVideoBindAPIRequest) SetVideoId(_videoId int64) error {
	r._videoId = _videoId
	r.Set("video_id", _videoId)
	return nil
}

// GetVideoId VideoId Getter
func (r TaobaoSubwayCreativeVideoBindAPIRequest) GetVideoId() int64 {
	return r._videoId
}

// SetSizeType is SizeType Setter
// 视频类型
func (r *TaobaoSubwayCreativeVideoBindAPIRequest) SetSizeType(_sizeType int64) error {
	r._sizeType = _sizeType
	r.Set("size_type", _sizeType)
	return nil
}

// GetSizeType SizeType Getter
func (r TaobaoSubwayCreativeVideoBindAPIRequest) GetSizeType() int64 {
	return r._sizeType
}

// SetItemId is ItemId Setter
// ItemId
func (r *TaobaoSubwayCreativeVideoBindAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoSubwayCreativeVideoBindAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetCreativeId is CreativeId Setter
// 创意ID
func (r *TaobaoSubwayCreativeVideoBindAPIRequest) SetCreativeId(_creativeId int64) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaoSubwayCreativeVideoBindAPIRequest) GetCreativeId() int64 {
	return r._creativeId
}
