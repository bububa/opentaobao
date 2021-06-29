package exchange

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家创建换货留言 API请求
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

// 初始化TmallExchangeMessageAddRequest对象
func NewTmallExchangeMessageAddRequest() *TmallExchangeMessageAddRequest{
    return &TmallExchangeMessageAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallExchangeMessageAddRequest) GetApiMethodName() string {
    return "tmall.exchange.message.add"
}

// IRequest interface 方法, 获取API参数
func (r TmallExchangeMessageAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Content Setter
// 留言内容
func (r *TmallExchangeMessageAddRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r TmallExchangeMessageAddRequest) GetContent() string {
    return r.content
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeMessageAddRequest) SetDisputeId(disputeId int64) error {
    r.disputeId = disputeId
    r.Set("dispute_id", disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeMessageAddRequest) GetDisputeId() int64 {
    return r.disputeId
}
// MessagePics Setter
// 凭证图片列表
func (r *TmallExchangeMessageAddRequest) SetMessagePics(messagePics []*model.File) error {
    r.messagePics = messagePics
    r.Set("message_pics", messagePics)
    return nil
}

// MessagePics Getter
func (r TmallExchangeMessageAddRequest) GetMessagePics() []*model.File {
    return r.messagePics
}
// Fields Setter
// 返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
func (r *TmallExchangeMessageAddRequest) SetFields(fields []string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r TmallExchangeMessageAddRequest) GetFields() []string {
    return r.fields
}
