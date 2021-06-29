package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵闹钟创建 API请求
taobao.ailab.aicloud.top.memo.alarm.create

天猫精灵闹钟创建
*/
type TaobaoAilabAicloudTopMemoAlarmCreateRequest struct {
    model.Params
    // 扩展信息json段，用于存放APP类型，APP版本等等信息。
    _ext   string
    // schema
    _schema   string
    // 企业用户ID
    _userId   string
    // 手持设备ID
    _utdId   string
    // 创建闹钟入参
    _paramCreateAlarmParam   *CreateAlarmParam
}

// 初始化TaobaoAilabAicloudTopMemoAlarmCreateRequest对象
func NewTaobaoAilabAicloudTopMemoAlarmCreateRequest() *TaobaoAilabAicloudTopMemoAlarmCreateRequest{
    return &TaobaoAilabAicloudTopMemoAlarmCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.alarm.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetExt() string {
    return r._ext
}
// Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetUtdId() string {
    return r._utdId
}
// ParamCreateAlarmParam Setter
// 创建闹钟入参
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetParamCreateAlarmParam(_paramCreateAlarmParam *CreateAlarmParam) error {
    r._paramCreateAlarmParam = _paramCreateAlarmParam
    r.Set("param_create_alarm_param", _paramCreateAlarmParam)
    return nil
}

// ParamCreateAlarmParam Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetParamCreateAlarmParam() *CreateAlarmParam {
    return r._paramCreateAlarmParam
}