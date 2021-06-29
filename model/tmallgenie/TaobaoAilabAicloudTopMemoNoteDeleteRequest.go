package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵备忘录删除 API请求
taobao.ailab.aicloud.top.memo.note.delete

删除天猫精灵用户设置的备忘录
*/
type TaobaoAilabAicloudTopMemoNoteDeleteRequest struct {
    model.Params
    // schema
    schema   string
    // 企业用户ID
    userId   string
    // 手持设备ID
    utdId   string
    // 扩展信息json段，用于存放APP类型，APP版本等等信息。
    ext   string
    // 备忘录ID
    memoId   int64
}

// 初始化TaobaoAilabAicloudTopMemoNoteDeleteRequest对象
func NewTaobaoAilabAicloudTopMemoNoteDeleteRequest() *TaobaoAilabAicloudTopMemoNoteDeleteRequest{
    return &TaobaoAilabAicloudTopMemoNoteDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.note.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetExt() string {
    return r.ext
}
// MemoId Setter
// 备忘录ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetMemoId(memoId int64) error {
    r.memoId = memoId
    r.Set("memo_id", memoId)
    return nil
}

// MemoId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetMemoId() int64 {
    return r.memoId
}
