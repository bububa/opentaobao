package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopLikeDeleteAPIRequest 取消收藏 API请求
// taobao.ailab.aicloud.top.like.delete
//
// 取消收藏
type TaobaoAilabAicloudTopLikeDeleteAPIRequest struct {
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

// NewTaobaoAilabAicloudTopLikeDeleteRequest 初始化TaobaoAilabAicloudTopLikeDeleteAPIRequest对象
func NewTaobaoAilabAicloudTopLikeDeleteRequest() *TaobaoAilabAicloudTopLikeDeleteAPIRequest {
	return &TaobaoAilabAicloudTopLikeDeleteAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopLikeDeleteAPIRequest) Reset() {
	r._ext = ""
	r._utdId = ""
	r._userId = ""
	r._schema = ""
	r._type = ""
	r._source = ""
	r._itemId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.like.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopLikeDeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetExt() string {
	return r._ext
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopLikeDeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopLikeDeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopLikeDeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetSchema() string {
	return r._schema
}

// SetType is Type Setter
// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
func (r *TaobaoAilabAicloudTopLikeDeleteAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetType() string {
	return r._type
}

// SetSource is Source Setter
// 来源
func (r *TaobaoAilabAicloudTopLikeDeleteAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetSource() string {
	return r._source
}

// SetItemId is ItemId Setter
// 资源的记录ID
func (r *TaobaoAilabAicloudTopLikeDeleteAPIRequest) SetItemId(_itemId string) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoAilabAicloudTopLikeDeleteAPIRequest) GetItemId() string {
	return r._itemId
}

var poolTaobaoAilabAicloudTopLikeDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopLikeDeleteRequest()
	},
}

// GetTaobaoAilabAicloudTopLikeDeleteRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopLikeDeleteAPIRequest
func GetTaobaoAilabAicloudTopLikeDeleteAPIRequest() *TaobaoAilabAicloudTopLikeDeleteAPIRequest {
	return poolTaobaoAilabAicloudTopLikeDeleteAPIRequest.Get().(*TaobaoAilabAicloudTopLikeDeleteAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopLikeDeleteAPIRequest 将 TaobaoAilabAicloudTopLikeDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopLikeDeleteAPIRequest(v *TaobaoAilabAicloudTopLikeDeleteAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopLikeDeleteAPIRequest.Put(v)
}
