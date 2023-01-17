package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest 天猫精灵闹钟创建 API请求
// taobao.ailab.aicloud.top.memo.alarm.create
//
// 天猫精灵闹钟创建
type TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest struct {
	model.Params
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// schema
	_schema string
	// 企业用户ID
	_userId string
	// 手持设备ID
	_utdId string
	// 创建闹钟入参
	_paramCreateAlarmParam *CreateAlarmParam
}

// NewTaobaoAilabAicloudTopMemoAlarmCreateRequest 初始化TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest对象
func NewTaobaoAilabAicloudTopMemoAlarmCreateRequest() *TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest {
	return &TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.alarm.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) GetExt() string {
	return r._ext
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetParamCreateAlarmParam is ParamCreateAlarmParam Setter
// 创建闹钟入参
func (r *TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) SetParamCreateAlarmParam(_paramCreateAlarmParam *CreateAlarmParam) error {
	r._paramCreateAlarmParam = _paramCreateAlarmParam
	r.Set("param_create_alarm_param", _paramCreateAlarmParam)
	return nil
}

// GetParamCreateAlarmParam ParamCreateAlarmParam Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest) GetParamCreateAlarmParam() *CreateAlarmParam {
	return r._paramCreateAlarmParam
}
