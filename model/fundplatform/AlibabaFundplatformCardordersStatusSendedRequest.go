package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
制卡商通知实体卡发货完成 API请求
alibaba.fundplatform.cardorders.status.sended

当制卡商将实体卡发货完成后，需要调用该接口，通知我们已发货。
*/
type AlibabaFundplatformCardordersStatusSendedRequest struct {
    model.Params
    // 子制卡单ID
    _cardOrderId   int64
    // 物流单号
    _logisticsOrderId   string
    // 物流商名称
    _logisticsCompany   string
}

// 初始化AlibabaFundplatformCardordersStatusSendedRequest对象
func NewAlibabaFundplatformCardordersStatusSendedRequest() *AlibabaFundplatformCardordersStatusSendedRequest{
    return &AlibabaFundplatformCardordersStatusSendedRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardordersStatusSendedRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorders.status.sended"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardordersStatusSendedRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CardOrderId Setter
// 子制卡单ID
func (r *AlibabaFundplatformCardordersStatusSendedRequest) SetCardOrderId(_cardOrderId int64) error {
    r._cardOrderId = _cardOrderId
    r.Set("card_order_id", _cardOrderId)
    return nil
}

// CardOrderId Getter
func (r AlibabaFundplatformCardordersStatusSendedRequest) GetCardOrderId() int64 {
    return r._cardOrderId
}
// LogisticsOrderId Setter
// 物流单号
func (r *AlibabaFundplatformCardordersStatusSendedRequest) SetLogisticsOrderId(_logisticsOrderId string) error {
    r._logisticsOrderId = _logisticsOrderId
    r.Set("logistics_order_id", _logisticsOrderId)
    return nil
}

// LogisticsOrderId Getter
func (r AlibabaFundplatformCardordersStatusSendedRequest) GetLogisticsOrderId() string {
    return r._logisticsOrderId
}
// LogisticsCompany Setter
// 物流商名称
func (r *AlibabaFundplatformCardordersStatusSendedRequest) SetLogisticsCompany(_logisticsCompany string) error {
    r._logisticsCompany = _logisticsCompany
    r.Set("logistics_company", _logisticsCompany)
    return nil
}

// LogisticsCompany Getter
func (r AlibabaFundplatformCardordersStatusSendedRequest) GetLogisticsCompany() string {
    return r._logisticsCompany
}
