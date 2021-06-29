package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵闹钟删除 API请求
taobao.ailab.aicloud.top.memo.alarm.delete

天猫精灵闹钟删除
*/
type TaobaoAilabAicloudTopMemoAlarmDeleteRequest struct {
    model.Params
    // schema
    schema   string
    // 手持设备ID
    utdId   string
    // 扩展信息json段，用于存放APP类型，APP版本等等信息。
    ext   string
    // 企业用户ID
    userId   string
    // 闹钟ID
    memoId   int64
}

// 初始化TaobaoAilabAicloudTopMemoAlarmDeleteRequest对象
func NewTaobaoAilabAicloudTopMemoAlarmDeleteRequest() *TaobaoAilabAicloudTopMemoAlarmDeleteRequest{
    return &TaobaoAilabAicloudTopMemoAlarmDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoAlarmDeleteRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.alarm.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoAlarmDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteRequest) GetSchema() string {
    return r.schema
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteRequest) GetExt() string {
    return r.ext
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteRequest) GetUserId() string {
    return r.userId
}
// MemoId Setter
// 闹钟ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteRequest) SetMemoId(memoId int64) error {
    r.memoId = memoId
    r.Set("memo_id", memoId)
    return nil
}

// MemoId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteRequest) GetMemoId() int64 {
    return r.memoId
}
