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
    ext   string
    // schema
    schema   string
    // 企业用户ID
    userId   string
    // 手持设备ID
    utdId   string
    // 创建闹钟入参
    paramCreateAlarmParam   *CreateAlarmParam
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
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetExt() string {
    return r.ext
}
// Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetUtdId() string {
    return r.utdId
}
// ParamCreateAlarmParam Setter
// 创建闹钟入参
func (r *TaobaoAilabAicloudTopMemoAlarmCreateRequest) SetParamCreateAlarmParam(paramCreateAlarmParam *CreateAlarmParam) error {
    r.paramCreateAlarmParam = paramCreateAlarmParam
    r.Set("param_create_alarm_param", paramCreateAlarmParam)
    return nil
}

// ParamCreateAlarmParam Getter
func (r TaobaoAilabAicloudTopMemoAlarmCreateRequest) GetParamCreateAlarmParam() *CreateAlarmParam {
    return r.paramCreateAlarmParam
}
