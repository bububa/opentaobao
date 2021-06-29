package aliexpress

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV更新INS私信发送的结果 APIRequest
aliexpress.social.ins.directresult.update

ISV更新INS私信发送的结果
*/
type AliexpressSocialInsDirectresultUpdateRequest struct {
    model.Params

    // 回调id,在获取图片时会返回
    id   int64 

    // 接受消息人INS_ID，也就是查询图片时的request_ins_id
    receiveInsId   string 

    // ISV发送私信人的INS_ID
    senderInsId   string 

    // 1.成功，2.失败。
    result   int64 

}

func NewAliexpressSocialInsDirectresultUpdateRequest() *AliexpressSocialInsDirectresultUpdateRequest{
    return &AliexpressSocialInsDirectresultUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSocialInsDirectresultUpdateRequest) GetApiMethodName() string {
    return "aliexpress.social.ins.directresult.update"
}

func (r AliexpressSocialInsDirectresultUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSocialInsDirectresultUpdateRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r AliexpressSocialInsDirectresultUpdateRequest) GetId() int64 {
    return r.id
}

func (r *AliexpressSocialInsDirectresultUpdateRequest) SetReceiveInsId(receiveInsId string) error {
    r.receiveInsId = receiveInsId
    r.Set("receive_ins_id", receiveInsId)
    return nil
}

func (r AliexpressSocialInsDirectresultUpdateRequest) GetReceiveInsId() string {
    return r.receiveInsId
}

func (r *AliexpressSocialInsDirectresultUpdateRequest) SetSenderInsId(senderInsId string) error {
    r.senderInsId = senderInsId
    r.Set("sender_ins_id", senderInsId)
    return nil
}

func (r AliexpressSocialInsDirectresultUpdateRequest) GetSenderInsId() string {
    return r.senderInsId
}

func (r *AliexpressSocialInsDirectresultUpdateRequest) SetResult(result int64) error {
    r.result = result
    r.Set("result", result)
    return nil
}

func (r AliexpressSocialInsDirectresultUpdateRequest) GetResult() int64 {
    return r.result
}

