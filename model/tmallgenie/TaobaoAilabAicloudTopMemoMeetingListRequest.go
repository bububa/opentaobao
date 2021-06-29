package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵会议查询 API请求
taobao.ailab.aicloud.top.memo.meeting.list

查询天猫精灵用户设置的所有会议
*/
type TaobaoAilabAicloudTopMemoMeetingListRequest struct {
    model.Params
    // schema
    schema   string
    // 企业用户ID
    userId   string
    // 手持设备ID
    utdId   string
    // 扩展信息json段，用于存放APP类型，APP版本等等信息。
    ext   string
    // 闹钟ID
    memoId   int64
}

// 初始化TaobaoAilabAicloudTopMemoMeetingListRequest对象
func NewTaobaoAilabAicloudTopMemoMeetingListRequest() *TaobaoAilabAicloudTopMemoMeetingListRequest{
    return &TaobaoAilabAicloudTopMemoMeetingListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoMeetingListRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.meeting.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoMeetingListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoMeetingListRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoMeetingListRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoMeetingListRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoMeetingListRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoMeetingListRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoMeetingListRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoMeetingListRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoMeetingListRequest) GetExt() string {
    return r.ext
}
// MemoId Setter
// 闹钟ID
func (r *TaobaoAilabAicloudTopMemoMeetingListRequest) SetMemoId(memoId int64) error {
    r.memoId = memoId
    r.Set("memo_id", memoId)
    return nil
}

// MemoId Getter
func (r TaobaoAilabAicloudTopMemoMeetingListRequest) GetMemoId() int64 {
    return r.memoId
}
