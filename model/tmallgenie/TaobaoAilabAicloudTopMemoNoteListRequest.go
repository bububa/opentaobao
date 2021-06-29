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
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetSchema(_schema string) error {
    r._schema = _schema
    r.Set("schema", _schema)
    return nil
}

// Schema Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetSchema() string {
    return r._schema
}
// UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetUserId() string {
    return r._userId
}
// UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetUtdId(_utdId string) error {
    r._utdId = _utdId
    r.Set("utd_id", _utdId)
    return nil
}

// UtdId Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetUtdId() string {
    return r._utdId
}
// Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetExt(_ext string) error {
    r._ext = _ext
    r.Set("ext", _ext)
    return nil
}

// Ext Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetExt() string {
    return r._ext
}
// MemoId Setter
// 备忘录ID
func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetMemoId(_memoId int64) error {
    r._memoId = _memoId
    r.Set("memo_id", _memoId)
    return nil
}

// MemoId Getter
func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetMemoId() int64 {
    return r._memoId
}
