package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtoplikeaddAPIRequest 增加收藏 API请求
// taobao.ailab.aicloud.top.like.add
//
// 将制定内容加入收藏
type TaobaoailabaicloudtoplikeaddAPIRequest struct {
	model.Params
	// 扩展信息，用于存放APP类型等
	_ext string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 账户体系隔离
	_schema string
	// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
	_type string
	// 来源
	_source string
	// 收藏的资源的ID
	_itemId string
	// 内容，必须要是一个json格式：{"song":"走过1999","singer":"张学友","album":"走过1999"}
	_content string
}

// NewTaobaoailabaicloudtoplikeaddRequest 初始化TaobaoailabaicloudtoplikeaddAPIRequest对象
func NewTaobaoailabaicloudtoplikeaddRequest() *TaobaoailabaicloudtoplikeaddAPIRequest {
	return &TaobaoailabaicloudtoplikeaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.like.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtoplikeaddAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetExt() string {
	return r._ext
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtoplikeaddAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtoplikeaddAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetUserId() string {
	return r._userId
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtoplikeaddAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetSchema() string {
	return r._schema
}

// SetType is Type Setter
// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
func (r *TaobaoailabaicloudtoplikeaddAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetType() string {
	return r._type
}

// SetSource is Source Setter
// 来源
func (r *TaobaoailabaicloudtoplikeaddAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetSource() string {
	return r._source
}

// SetItemId is ItemId Setter
// 收藏的资源的ID
func (r *TaobaoailabaicloudtoplikeaddAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetItemId() string {
	return r._itemId
}

// SetContent is Content Setter
// 内容，必须要是一个json格式：{&#34;song&#34;:&#34;走过1999&#34;,&#34;singer&#34;:&#34;张学友&#34;,&#34;album&#34;:&#34;走过1999&#34;}
func (r *TaobaoailabaicloudtoplikeaddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaoailabaicloudtoplikeaddAPIRequest) GetContent() string {
	return r._content
}
