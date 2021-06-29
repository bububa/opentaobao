package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建退款留言/凭证 API请求
taobao.refund.message.add

创建退款留言/凭证
*/
type TaobaoRefundMessageAddRequest struct {
    model.Params
    // 退款编号。
    refundId   int64
    // 留言内容。最大长度: 400个字节
    content   string
    // 图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K
    image   []*model.File
}

// 初始化TaobaoRefundMessageAddRequest对象
func NewTaobaoRefundMessageAddRequest() *TaobaoRefundMessageAddRequest{
    return &TaobaoRefundMessageAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRefundMessageAddRequest) GetApiMethodName() string {
    return "taobao.refund.message.add"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRefundMessageAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款编号。
func (r *TaobaoRefundMessageAddRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRefundMessageAddRequest) GetRefundId() int64 {
    return r.refundId
}
// Content Setter
// 留言内容。最大长度: 400个字节
func (r *TaobaoRefundMessageAddRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

// Content Getter
func (r TaobaoRefundMessageAddRequest) GetContent() string {
    return r.content
}
// Image Setter
// 图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K
func (r *TaobaoRefundMessageAddRequest) SetImage(image []*model.File) error {
    r.image = image
    r.Set("image", image)
    return nil
}

// Image Getter
func (r TaobaoRefundMessageAddRequest) GetImage() []*model.File {
    return r.image
}
