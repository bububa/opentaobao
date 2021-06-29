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
    _schema   string
    // 企业用户ID
    _userId   string
    // 手持设备ID
    _utdId   string
    // 扩展信息json段，用于存放APP类型，APP版本等等信息。
    _ext   string
    // 备忘录ID
    _memoId   int64
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
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetExt() string {
    return r._ext
}
// MemoId Setter
// 备忘录ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetMemoId(_memoId int64) error {
    r._memoId = _memoId
    r.Set("memo_id", _memoId)
    return nil
}

// MemoId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetMemoId() int64 {
    return r._memoId
}
