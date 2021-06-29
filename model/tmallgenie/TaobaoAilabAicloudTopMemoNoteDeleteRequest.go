package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵备忘录删除 APIRequest
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

func NewTaobaoAilabAicloudTopMemoNoteDeleteRequest() *TaobaoAilabAicloudTopMemoNoteDeleteRequest{
    return &TaobaoAilabAicloudTopMemoNoteDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.note.delete"
}

func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopMemoNoteDeleteRequest) SetMemoId(memoId int64) error {
    r.memoId = memoId
    r.Set("memo_id", memoId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteDeleteRequest) GetMemoId() int64 {
    return r.memoId
}

