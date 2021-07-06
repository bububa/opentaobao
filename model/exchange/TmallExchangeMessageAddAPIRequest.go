package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeMessageAddAPIRequest 卖家创建换货留言 API请求
// tmall.exchange.message.add
//
// 卖家创建换货留言
type TmallExchangeMessageAddAPIRequest struct {
	model.Params
	// 返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
	_fields []string
	// 留言内容
	_content string
	// 换货单号ID
	_disputeId int64
	// 凭证图片列表
	_messagePics *model.File
}

// NewTmallExchangeMessageAddRequest 初始化TmallExchangeMessageAddAPIRequest对象
func NewTmallExchangeMessageAddRequest() *TmallExchangeMessageAddAPIRequest {
	return &TmallExchangeMessageAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallExchangeMessageAddAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.message.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallExchangeMessageAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetFields is Fields Setter
// 返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
func (r *TmallExchangeMessageAddAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallExchangeMessageAddAPIRequest) GetFields() []string {
	return r._fields
}

// SetContent is Content Setter
// 留言内容
func (r *TmallExchangeMessageAddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TmallExchangeMessageAddAPIRequest) GetContent() string {
	return r._content
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallExchangeMessageAddAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallExchangeMessageAddAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetMessagePics is MessagePics Setter
// 凭证图片列表
func (r *TmallExchangeMessageAddAPIRequest) SetMessagePics(_messagePics *model.File) error {
	r._messagePics = _messagePics
	r.Set("message_pics", _messagePics)
	return nil
}

// GetMessagePics MessagePics Getter
func (r TmallExchangeMessageAddAPIRequest) GetMessagePics() *model.File {
	return r._messagePics
}
