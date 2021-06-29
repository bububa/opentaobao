package tmallgenie

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫精灵闹钟查询 APIRequest
taobao.ailab.aicloud.top.memo.alarm.list

查询天猫精灵用户设置的所有闹钟
*/
type TaobaoAilabAicloudTopMemoAlarmListRequest struct {
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

func NewTaobaoAilabAicloudTopMemoAlarmListRequest() *TaobaoAilabAicloudTopMemoAlarmListRequest{
    return &TaobaoAilabAicloudTopMemoAlarmListRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAilabAicloudTopMemoAlarmListRequest) GetApiMethodName() string {
    return "taobao.ailab.aicloud.top.memo.alarm.list"
}

func (r TaobaoAilabAicloudTopMemoAlarmListRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAilabAicloudTopMemoAlarmListRequest) SetSchema(schema string) error {
    r.schema = schema
    r.Set("schema", schema)
    return nil
}

func (r TaobaoAilabAicloudTopMemoAlarmListRequest) GetSchema() string {
    return r.schema
}

func (r *TaobaoAilabAicloudTopMemoAlarmListRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoAlarmListRequest) GetUserId() string {
    return r.userId
}

func (r *TaobaoAilabAicloudTopMemoAlarmListRequest) SetUtdId(utdId string) error {
    r.utdId = utdId
    r.Set("utd_id", utdId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoAlarmListRequest) GetUtdId() string {
    return r.utdId
}

func (r *TaobaoAilabAicloudTopMemoAlarmListRequest) SetExt(ext string) error {
    r.ext = ext
    r.Set("ext", ext)
    return nil
}

func (r TaobaoAilabAicloudTopMemoAlarmListRequest) GetExt() string {
    return r.ext
}

func (r *TaobaoAilabAicloudTopMemoAlarmListRequest) SetMemoId(memoId int64) error {
    r.memoId = memoId
    r.Set("memo_id", memoId)
    return nil
}

func (r TaobaoAilabAicloudTopMemoAlarmListRequest) GetMemoId() int64 {
    return r.memoId
}

