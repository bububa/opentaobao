package tbrefund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaorefundmessageaddAPIRequest 创建退款留言/凭证 API请求
// taobao.refund.message.add
//
// 创建退款留言/凭证
type TaobaorefundmessageaddAPIRequest struct {
	model.Params
	// 留言内容。最大长度: 400个字节
	_content string
	// 退款编号。
	_refundId int64
	// 图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K
	_image *model.File
}

// NewTaobaorefundmessageaddRequest 初始化TaobaorefundmessageaddAPIRequest对象
func NewTaobaorefundmessageaddRequest() *TaobaorefundmessageaddAPIRequest {
	return &TaobaorefundmessageaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaorefundmessageaddAPIRequest) GetApiMethodName() string {
	return "taobao.refund.message.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaorefundmessageaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaorefundmessageaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetContent is Content Setter
// 留言内容。最大长度: 400个字节
func (r *TaobaorefundmessageaddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// GetContent Content Getter
func (r TaobaorefundmessageaddAPIRequest) GetContent() string {
	return r._content
}

// SetRefundId is RefundId Setter
// 退款编号。
func (r *TaobaorefundmessageaddAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// GetRefundId RefundId Getter
func (r TaobaorefundmessageaddAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// SetImage is Image Setter
// 图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K
func (r *TaobaorefundmessageaddAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// GetImage Image Getter
func (r TaobaorefundmessageaddAPIRequest) GetImage() *model.File {
	return r._image
}
