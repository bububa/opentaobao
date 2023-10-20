package exchange

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallexchangemessageaddAPIRequest 卖家创建换货留言 API请求
// tmall.exchange.message.add
//
// 卖家创建换货留言
type TmallexchangemessageaddAPIRequest struct {
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

// NewTmallexchangemessageaddRequest 初始化TmallexchangemessageaddAPIRequest对象
func NewTmallexchangemessageaddRequest() *TmallexchangemessageaddAPIRequest {
	return &TmallexchangemessageaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallexchangemessageaddAPIRequest) GetApiMethodName() string {
	return "tmall.exchange.message.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallexchangemessageaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallexchangemessageaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFields is Fields Setter
// 返回字段。目前支持id,refund_id,owner_id,owner_nick,owner_role,content,pic_urls,created,message_type
func (r *TmallexchangemessageaddAPIRequest) SetFields(_fields []string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r TmallexchangemessageaddAPIRequest) GetFields() []string {
	return r._fields
}

// SetContent is Content Setter
// 留言内容
func (r *TmallexchangemessageaddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TmallexchangemessageaddAPIRequest) GetContent() string {
	return r._content
}

// SetDisputeId is DisputeId Setter
// 换货单号ID
func (r *TmallexchangemessageaddAPIRequest) SetDisputeId(_disputeId int64) error {
	r._disputeId = _disputeId
	r.Set("dispute_id", _disputeId)
	return nil
}

// GetDisputeId DisputeId Getter
func (r TmallexchangemessageaddAPIRequest) GetDisputeId() int64 {
	return r._disputeId
}

// SetMessagePics is MessagePics Setter
// 凭证图片列表
func (r *TmallexchangemessageaddAPIRequest) SetMessagePics(_messagePics *model.File) error {
	r._messagePics = _messagePics
	r.Set("message_pics", _messagePics)
	return nil
}

// GetMessagePics MessagePics Getter
func (r TmallexchangemessageaddAPIRequest) GetMessagePics() *model.File {
	return r._messagePics
}
