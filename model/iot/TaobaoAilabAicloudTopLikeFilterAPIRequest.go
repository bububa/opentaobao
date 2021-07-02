package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopLikeFilterAPIRequest 过滤列表歌曲存在于收藏列表的 API请求
// taobao.ailab.aicloud.top.like.filter
//
// 过滤出传入列表歌曲存在于收藏列表的
type TaobaoAilabAicloudTopLikeFilterAPIRequest struct {
	model.Params
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
	// 传入的歌曲列表
	_mediaItems []MediaItem
}

// NewTaobaoAilabAicloudTopLikeFilterRequest 初始化TaobaoAilabAicloudTopLikeFilterAPIRequest对象
func NewTaobaoAilabAicloudTopLikeFilterRequest() *TaobaoAilabAicloudTopLikeFilterAPIRequest {
	return &TaobaoAilabAicloudTopLikeFilterAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopLikeFilterAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.like.filter"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopLikeFilterAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopLikeFilterAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// Get Schema Getter
func (r TaobaoAilabAicloudTopLikeFilterAPIRequest) GetSchema() string {
	return r._schema
}

// Set is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopLikeFilterAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoAilabAicloudTopLikeFilterAPIRequest) GetUserId() string {
	return r._userId
}

// Set is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopLikeFilterAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// Get UtdId Getter
func (r TaobaoAilabAicloudTopLikeFilterAPIRequest) GetUtdId() string {
	return r._utdId
}

// Set is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopLikeFilterAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// Get Ext Getter
func (r TaobaoAilabAicloudTopLikeFilterAPIRequest) GetExt() string {
	return r._ext
}

// Set is Type Setter
// 音频收藏类型, 四种类型：music,children_song,program,story
func (r *TaobaoAilabAicloudTopLikeFilterAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// Get Type Getter
func (r TaobaoAilabAicloudTopLikeFilterAPIRequest) GetType() string {
	return r._type
}

// Set is MediaItems Setter
// 传入的歌曲列表
func (r *TaobaoAilabAicloudTopLikeFilterAPIRequest) SetMediaItems(_mediaItems []MediaItem) error {
	r._mediaItems = _mediaItems
	r.Set("media_items", _mediaItems)
	return nil
}

// Get MediaItems Getter
func (r TaobaoAilabAicloudTopLikeFilterAPIRequest) GetMediaItems() []MediaItem {
	return r._mediaItems
}
