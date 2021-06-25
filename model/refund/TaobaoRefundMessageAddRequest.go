package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建退款留言/凭证 APIRequest
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
    image   []byte 

}

func NewTaobaoRefundMessageAddRequest() *TaobaoRefundMessageAddRequest{
    return &TaobaoRefundMessageAddRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRefundMessageAddRequest) GetApiMethodName() string {
    return "taobao.refund.message.add"
}

func (r TaobaoRefundMessageAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRefundMessageAddRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

func (r TaobaoRefundMessageAddRequest) GetRefundId() int64 {
    return r.refundId
}

func (r *TaobaoRefundMessageAddRequest) SetContent(content string) error {
    r.content = content
    r.Set("content", content)
    return nil
}

func (r TaobaoRefundMessageAddRequest) GetContent() string {
    return r.content
}

func (r *TaobaoRefundMessageAddRequest) SetImage(image []byte) error {
    r.image = image
    r.Set("image", image)
    return nil
}

func (r TaobaoRefundMessageAddRequest) GetImage() []byte {
    return r.image
}

