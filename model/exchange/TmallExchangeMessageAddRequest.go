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
    _content   string
    // 换货单号ID
    _disputeId   int64
    // 凭证图片列表
    _messagePics   []*model.File
    // 返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
    _fields   []string
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
func (r *TmallExchangeMessageAddRequest) SetContent(_content string) error {
    r._content = _content
    r.Set("content", _content)
    return nil
}

// Content Getter
func (r TmallExchangeMessageAddRequest) GetContent() string {
    return r._content
}
// DisputeId Setter
// 换货单号ID
func (r *TmallExchangeMessageAddRequest) SetDisputeId(_disputeId int64) error {
    r._disputeId = _disputeId
    r.Set("dispute_id", _disputeId)
    return nil
}

// DisputeId Getter
func (r TmallExchangeMessageAddRequest) GetDisputeId() int64 {
    return r._disputeId
}
// MessagePics Setter
// 凭证图片列表
func (r *TmallExchangeMessageAddRequest) SetMessagePics(_messagePics []*model.File) error {
    r._messagePics = _messagePics
    r.Set("message_pics", _messagePics)
    return nil
}

// MessagePics Getter
func (r TmallExchangeMessageAddRequest) GetMessagePics() []*model.File {
    return r._messagePics
}
// Fields Setter
// 返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
func (r *TmallExchangeMessageAddRequest) SetFields(_fields []string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r TmallExchangeMessageAddRequest) GetFields() []string {
    return r._fields
}
