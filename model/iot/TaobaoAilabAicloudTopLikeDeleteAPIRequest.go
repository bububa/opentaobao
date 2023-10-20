package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtoplikedeleteAPIRequest 取消收藏 API请求
// taobao.ailab.aicloud.top.like.delete
//
// 取消收藏
type TaobaoailabaicloudtoplikedeleteAPIRequest struct {
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
	// 资源的记录ID
	_itemId string
}

// NewTaobaoailabaicloudtoplikedeleteRequest 初始化TaobaoailabaicloudtoplikedeleteAPIRequest对象
func NewTaobaoailabaicloudtoplikedeleteRequest() *TaobaoailabaicloudtoplikedeleteAPIRequest {
	return &TaobaoailabaicloudtoplikedeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.like.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtoplikedeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetExt() string {
	return r._ext
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtoplikedeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtoplikedeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtoplikedeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetSchema() string {
	return r._schema
}

// SetType is Type Setter
// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
func (r *TaobaoailabaicloudtoplikedeleteAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetType() string {
	return r._type
}

// SetSource is Source Setter
// 来源
func (r *TaobaoailabaicloudtoplikedeleteAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetSource() string {
	return r._source
}

// SetItemId is ItemId Setter
// 资源的记录ID
func (r *TaobaoailabaicloudtoplikedeleteAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoailabaicloudtoplikedeleteAPIRequest) GetItemId() string {
	return r._itemId
}
