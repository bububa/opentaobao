package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
制卡商通知制卡完成 APIRequest
alibaba.fundplatform.cardorders.status.make.finish

当制卡完成后，制卡商需要调用该接口，通知我们制卡已完成。
*/
type AlibabaFundplatformCardordersStatusMakeFinishRequest struct {
    model.Params

    // 子制卡单ID
    cardOrderId   int64 

}

func NewAlibabaFundplatformCardordersStatusMakeFinishRequest() *AlibabaFundplatformCardordersStatusMakeFinishRequest{
    return &AlibabaFundplatformCardordersStatusMakeFinishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardordersStatusMakeFinishRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorders.status.make.finish"
}

func (r AlibabaFundplatformCardordersStatusMakeFinishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardordersStatusMakeFinishRequest) SetCardOrderId(cardOrderId int64) error {
    r.cardOrderId = cardOrderId
    r.Set("card_order_id", cardOrderId)
    return nil
}

func (r AlibabaFundplatformCardordersStatusMakeFinishRequest) GetCardOrderId() int64 {
    return r.cardOrderId
}

