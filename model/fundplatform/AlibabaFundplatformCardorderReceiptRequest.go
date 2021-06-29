package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知确认收货 API请求
alibaba.fundplatform.cardorder.receipt

告知卡商这一批储值卡已经被用户确认收货
*/
type AlibabaFundplatformCardorderReceiptRequest struct {
    model.Params
    // 通知制卡成功的子卡子单号
    _cardOrderId   int64
    // 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
    _ownSign   string
}

// 初始化AlibabaFundplatformCardorderReceiptRequest对象
func NewAlibabaFundplatformCardorderReceiptRequest() *AlibabaFundplatformCardorderReceiptRequest{
    return &AlibabaFundplatformCardorderReceiptRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderReceiptRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.receipt"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderReceiptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CardOrderId Setter
// 通知制卡成功的子卡子单号
func (r *AlibabaFundplatformCardorderReceiptRequest) SetCardOrderId(_cardOrderId int64) error {
    r._cardOrderId = _cardOrderId
    r.Set("card_order_id", _cardOrderId)
    return nil
}

// CardOrderId Getter
func (r AlibabaFundplatformCardorderReceiptRequest) GetCardOrderId() int64 {
    return r._cardOrderId
}
// OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabaFundplatformCardorderReceiptRequest) SetOwnSign(_ownSign string) error {
    r._ownSign = _ownSign
    r.Set("own_sign", _ownSign)
    return nil
}

// OwnSign Getter
func (r AlibabaFundplatformCardorderReceiptRequest) GetOwnSign() string {
    return r._ownSign
}
