package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询天猫精灵用户设置的所有备忘录 APIRequest
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

func NewTaobaoAilabAicloudTopMemoNoteListRequest() *TaobaoAilabAicloudTopMemoNoteListRequest{
    return &TaobaoAilabAicloudTopMemoNoteListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.note.list"
}

func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopMemoNoteListRequest) SetMemoId(memoId int64) error {
    r.memoId = memoId
    r.Set("memo_id", memoId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoNoteListRequest) GetMemoId() int64 {
    return r.memoId
}

