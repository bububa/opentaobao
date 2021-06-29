package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
制卡商通知实体卡发货完成 APIRequest
alibaba.fundplatform.cardorders.status.sended

当制卡商将实体卡发货完成后，需要调用该接口，通知我们已发货。
*/
type AlibabaFundplatformCardordersStatusSendedRequest struct {
    model.Params

    // 子制卡单ID
    cardOrderId   int64 

    // 物流单号
    logisticsOrderId   string 

    // 物流商名称
    logisticsCompany   string 

}

func NewAlibabaFundplatformCardordersStatusSendedRequest() *AlibabaFundplatformCardordersStatusSendedRequest{
    return &AlibabaFundplatformCardordersStatusSendedRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaFundplatformCardordersStatusSendedRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorders.status.sended"
}

func (r AlibabaFundplatformCardordersStatusSendedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaFundplatformCardordersStatusSendedRequest) SetCardOrderId(cardOrderId int64) error {
    r.cardOrderId = cardOrderId
    r.Set("card_order_id", cardOrderId)
    return nil
}

func (r AlibabaFundplatformCardordersStatusSendedRequest) GetCardOrderId() int64 {
    return r.cardOrderId
}

func (r *AlibabaFundplatformCardordersStatusSendedRequest) SetLogisticsOrderId(logisticsOrderId string) error {
    r.logisticsOrderId = logisticsOrderId
    r.Set("logistics_order_id", logisticsOrderId)
    return nil
}

func (r AlibabaFundplatformCardordersStatusSendedRequest) GetLogisticsOrderId() string {
    return r.logisticsOrderId
}

func (r *AlibabaFundplatformCardordersStatusSendedRequest) SetLogisticsCompany(logisticsCompany string) error {
    r.logisticsCompany = logisticsCompany
    r.Set("logistics_company", logisticsCompany)
    return nil
}

func (r AlibabaFundplatformCardordersStatusSendedRequest) GetLogisticsCompany() string {
    return r.logisticsCompany
}

