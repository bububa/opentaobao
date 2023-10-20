package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtoplikelistAPIRequest 列出收藏列表 API请求
// taobao.ailab.aicloud.top.like.list
//
// 列出收藏列表
type TaobaoailabaicloudtoplikelistAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
	_param1 string
	// 页码 从0起
	_param2 int64
	// 每页条目个数
	_param3 int64
}

// NewTaobaoailabaicloudtoplikelistRequest 初始化TaobaoailabaicloudtoplikelistAPIRequest对象
func NewTaobaoailabaicloudtoplikelistRequest() *TaobaoailabaicloudtoplikelistAPIRequest {
	return &TaobaoailabaicloudtoplikelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.like.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoailabaicloudtoplikelistAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoailabaicloudtoplikelistAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoailabaicloudtoplikelistAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoailabaicloudtoplikelistAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetExt() string {
	return r._ext
}

// SetParam1 is Param1 Setter
// 收藏类型，目前支持四种：story,children_song,music,program，分别表示故事、儿童、音乐和节目
func (r *TaobaoailabaicloudtoplikelistAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam2 is Param2 Setter
// 页码 从0起
func (r *TaobaoailabaicloudtoplikelistAPIRequest) SetParam2(_param2 int64) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetParam2() int64 {
	return r._param2
}

// SetParam3 is Param3 Setter
// 每页条目个数
func (r *TaobaoailabaicloudtoplikelistAPIRequest) SetParam3(_param3 int64) error {
	r._param3 = _param3
	r.Set("param3", _param3)
	return nil
}

// GetParam3 Param3 Getter
func (r TaobaoailabaicloudtoplikelistAPIRequest) GetParam3() int64 {
	return r._param3
}
