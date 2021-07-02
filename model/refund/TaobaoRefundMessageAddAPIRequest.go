package refund

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoRefundMessageAddAPIRequest 创建退款留言/凭证 API请求
// taobao.refund.message.add
//
// 创建退款留言/凭证
type TaobaoRefundMessageAddAPIRequest struct {
	model.Params
	// 退款编号。
	_refundId int64
	// 留言内容。最大长度: 400个字节
	_content string
	// 图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K
	_image *model.File
}

// NewTaobaoRefundMessageAddRequest 初始化TaobaoRefundMessageAddAPIRequest对象
func NewTaobaoRefundMessageAddRequest() *TaobaoRefundMessageAddAPIRequest {
	return &TaobaoRefundMessageAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoRefundMessageAddAPIRequest) GetApiMethodName() string {
	return "taobao.refund.message.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoRefundMessageAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefundId Setter
// 退款编号。
func (r *TaobaoRefundMessageAddAPIRequest) SetRefundId(_refundId int64) error {
	r._refundId = _refundId
	r.Set("refund_id", _refundId)
	return nil
}

// Get RefundId Getter
func (r TaobaoRefundMessageAddAPIRequest) GetRefundId() int64 {
	return r._refundId
}

// Set is Content Setter
// 留言内容。最大长度: 400个字节
func (r *TaobaoRefundMessageAddAPIRequest) SetContent(_content string) error {
	r._content = _content
	r.Set("content", _content)
	return nil
}

// Get Content Getter
func (r TaobaoRefundMessageAddAPIRequest) GetContent() string {
	return r._content
}

// Set is Image Setter
// 图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K
func (r *TaobaoRefundMessageAddAPIRequest) SetImage(_image *model.File) error {
	r._image = _image
	r.Set("image", _image)
	return nil
}

// Get Image Getter
func (r TaobaoRefundMessageAddAPIRequest) GetImage() *model.File {
	return r._image
}
