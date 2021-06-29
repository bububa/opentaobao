package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询天猫精灵用户设置的所有备忘录 API请求
taobao.ailab.aicloud.top.memo.note.list

查询天猫精灵用户设置的所有备忘录
*/
type TaobaoAilabAicloudTopMemoNoteListRequest struct {
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

// 初始化TaobaoAilabAicloudTopMemoNoteListRequest对象
func NewTaobaoAilabAicloudTopMemoNoteListRequest() *TaobaoAilabAicloudTopMemoNoteListRequest{
    return &TaobaoAilabAicloudTopMemoNoteListRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.note.list"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetSchema() string {
    return r.schema
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetUserId() string {
    return r.userId
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetUtdId() string {
    return r.utdId
}
// Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetExt() string {
    return r.ext
}
// MemoId Setter
// 备忘录ID
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetMemoId(memoId int64) error {
    r.memoId = memoId
    r.Set("memo_id", memoId)
    return nil
}

// MemoId Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetMemoId() int64 {
    return r.memoId
}
