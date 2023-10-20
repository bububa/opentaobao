package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligeniealbumsgetAPIRequest 专辑详情 API请求
// alibaba.ailabs.aligenie.albums.get
//
// 给予厂商查询专辑下的音频详情
type AlibabaailabsaligeniealbumsgetAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 专辑 id
	_param1 int64
	// 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
	_param2 int64
	// 每页数量（不超过50）
	_param3 int64
}

// NewAlibabaailabsaligeniealbumsgetRequest 初始化AlibabaailabsaligeniealbumsgetAPIRequest对象
func NewAlibabaailabsaligeniealbumsgetRequest() *AlibabaailabsaligeniealbumsgetAPIRequest {
	return &AlibabaailabsaligeniealbumsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.albums.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *AlibabaailabsaligeniealbumsgetAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *AlibabaailabsaligeniealbumsgetAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *AlibabaailabsaligeniealbumsgetAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *AlibabaailabsaligeniealbumsgetAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetExt() string {
	return r._ext
}

// SetParam1 is Param1 Setter
// 专辑 id
func (r *AlibabaailabsaligeniealbumsgetAPIRequest) SetParam1(_param1 int64) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetParam1() int64 {
	return r._param1
}

// SetParam2 is Param2 Setter
// 当前页（从1开始, 目前由于底层引擎限制，实际最多能查5000个，超过5000返回为空，请确保页码*分页大小不超过5000）
func (r *AlibabaailabsaligeniealbumsgetAPIRequest) SetParam2(_param2 int64) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetParam2() int64 {
	return r._param2
}

// SetParam3 is Param3 Setter
// 每页数量（不超过50）
func (r *AlibabaailabsaligeniealbumsgetAPIRequest) SetParam3(_param3 int64) error {
	r._param3 = _param3
	r.Set("param3", _param3)
	return nil
}

// GetParam3 Param3 Getter
func (r AlibabaailabsaligeniealbumsgetAPIRequest) GetParam3() int64 {
	return r._param3
}
