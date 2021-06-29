package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵会议删除 APIRequest
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

func NewTaobaoAilabAicloudTopMemoMeetingDeleteRequest() *TaobaoAilabAicloudTopMemoMeetingDeleteRequest{
    return &TaobaoAilabAicloudTopMemoMeetingDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.meeting.delete"
}

func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopMemoMeetingDeleteRequest) SetMemoId(memoId int64) error {
    r.memoId = memoId
    r.Set("memo_id", memoId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoMeetingDeleteRequest) GetMemoId() int64 {
    return r.memoId
}

