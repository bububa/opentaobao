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
type TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest struct {
    model.Params
    // schema
    _schema   string
    // 手持设备ID
    _utdId   string
    // 扩展信息json段，用于存放APP类型，APP版本等等信息。
    _ext   string
    // 企业用户ID
    _userId   string
    // 闹钟ID
    _memoId   int64
}

// 初始化TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest对象
func NewTaobaoAilabAicloudTopMemoAlarmDeleteRequest() *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest{
    return &TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.alarm.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetSchema() string {
    return r._schema
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetExt() string {
    return r._ext
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetUserId() string {
    return r._userId
}
// MemoId Setter
// 闹钟ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetMemoId(_memoId int64) error {
    r._memoId = _memoId
    r.Set("memo_id", _memoId)
    return nil
}

// MemoId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetMemoId() int64 {
    return r._memoId
}
