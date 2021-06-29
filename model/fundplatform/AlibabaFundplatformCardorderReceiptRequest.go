package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知确认收货 APIRequest
alibaba.fundplatform.cardorder.receipt

告知卡商这一批储值卡已经被用户确认收货
*/
type AlibabaFundplatformCardorderReceiptRequest struct {
    model.Params

    // 通知制卡成功的子卡子单号
    cardOrderId   int64 

    // 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
    ownSign   string 

}

func NewAlibabaFundplatformCardorderReceiptRequest() *AlibabaFundplatformCardorderReceiptRequest{
    return &AlibabaFundplatformCardorderReceiptRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardorderReceiptRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.receipt"
}

func (r AlibabaFundplatformCardorderReceiptRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardorderReceiptRequest) SetCardOrderId(cardOrderId int64) error {
    r.cardOrderId = cardOrderId
    r.Set("card_order_id", cardOrderId)
    return nil
}

func (r AlibabaFundplatformCardorderReceiptRequest) GetCardOrderId() int64 {
    return r.cardOrderId
}

func (r *AlibabaFundplatformCardorderReceiptRequest) SetOwnSign(ownSign string) error {
    r.ownSign = ownSign
    r.Set("own_sign", ownSign)
    return nil
}

func (r AlibabaFundplatformCardorderReceiptRequest) GetOwnSign() string {
    return r.ownSign
}

