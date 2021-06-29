package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家创建换货留言 APIRequest
tmall.exchange.message.add

卖家创建换货留言
*/
type TmallExchangeMessageAddRequest struct {
    model.Params

    // 留言内容
    content   string 

    // 换货单号ID
    disputeId   int64 

    // 凭证图片列表
    messagePics   []*model.File 

    // 返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
    fields   []string 

}

func NewTmallExchangeMessageAddRequest() *TmallExchangeMessageAddRequest{
    return &TmallExchangeMessageAddRequest{
        Params: model.NewParams(),
    }
}

func (r TmallExchangeMessageAddRequest) GetApiMethodName() string {
    return "tmall.exchange.message.add"
}

func (r TmallExchangeMessageAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallExchangeMessageAddRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TmallExchangeMessageAddRequest) GetContent() string {
    return r.content
}

func (r *TmallExchangeMessageAddRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

func (r TmallExchangeMessageAddRequest) GetDisputeId() int64 {
    return r.disputeId
}

func (r *TmallExchangeMessageAddRequest) SetMessagePics(messagePics []*model.File) error {
    r.messagePics = messagePics
    r.Set("message_pics", messagePics)
    return nil
}

func (r TmallExchangeMessageAddRequest) GetMessagePics() []*model.File {
    return r.messagePics
}

func (r *TmallExchangeMessageAddRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r TmallExchangeMessageAddRequest) GetFields() []string {
    return r.fields
}

