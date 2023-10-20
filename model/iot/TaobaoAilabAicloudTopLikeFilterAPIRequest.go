package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtoplikefilterAPIRequest 过滤列表歌曲存在于收藏列表的 API请求
// taobao.ailab.aicloud.top.like.filter
//
// 过滤出传入列表歌曲存在于收藏列表的
type TaobaoailabaicloudtoplikefilterAPIRequest struct {
	model.Params
	// 传入的歌曲列表
	_mediaItems []MediaItem
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 音频收藏类型, 四种类型：music,children_song,program,story
	_type string
}

// NewTaobaoailabaicloudtoplikefilterRequest 初始化TaobaoailabaicloudtoplikefilterAPIRequest对象
func NewTaobaoailabaicloudtoplikefilterRequest() *TaobaoailabaicloudtoplikefilterAPIRequest {
	return &TaobaoailabaicloudtoplikefilterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtoplikefilterAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.like.filter"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtoplikefilterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtoplikefilterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMediaItems is MediaItems Setter
// 传入的歌曲列表
func (r *TaobaoailabaicloudtoplikefilterAPIRequest) SetMediaItems(_mediaItems []MediaItem) error {
	r._mediaItems = _mediaItems
	r.Set("media_items", _mediaItems)
	return nil
}

// GetMediaItems MediaItems Getter
func (r TaobaoailabaicloudtoplikefilterAPIRequest) GetMediaItems() []MediaItem {
	return r._mediaItems
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtoplikefilterAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtoplikefilterAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtoplikefilterAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtoplikefilterAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtoplikefilterAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtoplikefilterAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtoplikefilterAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtoplikefilterAPIRequest) GetExt() string {
	return r._ext
}

// SetType is Type Setter
// 音频收藏类型, 四种类型：music,children_song,program,story
func (r *TaobaoailabaicloudtoplikefilterAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoailabaicloudtoplikefilterAPIRequest) GetType() string {
	return r._type
}
