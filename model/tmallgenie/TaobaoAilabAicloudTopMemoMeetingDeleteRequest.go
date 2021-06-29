package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵会议删除 API请求
taobao.ailab.aicloud.top.memo.meeting.delete

天猫精灵会议删除
*/
type TaobaoAilabAicloudTopMemoMeetingDeleteRequest struct {
    model.Params
    // schema
    schema   string
    // 企业用户ID
    userId   string
    // 手持设备ID
    utdId   string
    // 扩展信息json段，用于存放APP类型，APP版本等等信息。
    ext   string
    // 会议ID
    memoId   int64
}

// 初始化TaobaoAilabAicloudTopMemoMeetingDeleteRequest对象
func NewTaobaoAilabAicloudTopMemoMeetingDeleteRequest() *TaobaoAilabAicloudTopMemoMeetingDeleteRequest{
    return &TaobaoAilabAicloudTopMemoMeetingDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.meeting.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetExt() string {
    return r.ext
}
// MemoId Setter
// 会议ID
func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetMemoId(memoId int64) error {
    r.memoId = memoId
    r.Set("memo_id", memoId)
    return nil
}

// MemoId Getter
func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetMemoId() int64 {
    return r.memoId
}
