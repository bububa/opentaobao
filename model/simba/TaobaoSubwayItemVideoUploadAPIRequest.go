package simba

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayItemVideoUploadAPIRequest 创意视频上传 API请求
// taobao.subway.item.video.upload
//
// 为用户提供视频上传的功能
type TaobaoSubwayItemVideoUploadAPIRequest struct {
	model.Params
	// 账号昵称
	_nick string
	// video的url
	_videoUrl string
	// itemId
	_itemId int64
	// 视频类型，1是方视频
	_type int64
}

// NewTaobaoSubwayItemVideoUploadRequest 初始化TaobaoSubwayItemVideoUploadAPIRequest对象
func NewTaobaoSubwayItemVideoUploadRequest() *TaobaoSubwayItemVideoUploadAPIRequest {
	return &TaobaoSubwayItemVideoUploadAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoSubwayItemVideoUploadAPIRequest) Reset() {
	r._nick = ""
	r._videoUrl = ""
	r._itemId = 0
	r._type = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSubwayItemVideoUploadAPIRequest) GetApiMethodName() string {
	return "taobao.subway.item.video.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSubwayItemVideoUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoSubwayItemVideoUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNick is Nick Setter
// 账号昵称
func (r *TaobaoSubwayItemVideoUploadAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSubwayItemVideoUploadAPIRequest) GetNick() string {
	return r._nick
}

// SetVideoUrl is VideoUrl Setter
// video的url
func (r *TaobaoSubwayItemVideoUploadAPIRequest) SetVideoUrl(_videoUrl string) error {
	r._videoUrl = _videoUrl
	r.Set("video_url", _videoUrl)
	return nil
}

// GetVideoUrl VideoUrl Getter
func (r TaobaoSubwayItemVideoUploadAPIRequest) GetVideoUrl() string {
	return r._videoUrl
}

// SetItemId is ItemId Setter
// itemId
func (r *TaobaoSubwayItemVideoUploadAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoSubwayItemVideoUploadAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetType is Type Setter
// 视频类型，1是方视频
func (r *TaobaoSubwayItemVideoUploadAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoSubwayItemVideoUploadAPIRequest) GetType() int64 {
	return r._type
}

var poolTaobaoSubwayItemVideoUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoSubwayItemVideoUploadRequest()
	},
}

// GetTaobaoSubwayItemVideoUploadRequest 从 sync.Pool 获取 TaobaoSubwayItemVideoUploadAPIRequest
func GetTaobaoSubwayItemVideoUploadAPIRequest() *TaobaoSubwayItemVideoUploadAPIRequest {
	return poolTaobaoSubwayItemVideoUploadAPIRequest.Get().(*TaobaoSubwayItemVideoUploadAPIRequest)
}

// ReleaseTaobaoSubwayItemVideoUploadAPIRequest 将 TaobaoSubwayItemVideoUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoSubwayItemVideoUploadAPIRequest(v *TaobaoSubwayItemVideoUploadAPIRequest) {
	v.Reset()
	poolTaobaoSubwayItemVideoUploadAPIRequest.Put(v)
}
