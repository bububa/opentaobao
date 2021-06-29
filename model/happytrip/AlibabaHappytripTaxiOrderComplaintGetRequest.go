package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
投诉详情 APIRequest
alibaba.happytrip.taxi.order.complaint.get

获取投诉处理进度详情
*/
type AlibabaHappytripTaxiOrderComplaintGetRequest struct {
    model.Params

    // 供应商订单号
    orderId   string 

    // 供应商工单号
    caseId   string 

}

func NewAlibabaHappytripTaxiOrderComplaintGetRequest() *AlibabaHappytripTaxiOrderComplaintGetRequest{
    return &AlibabaHappytripTaxiOrderComplaintGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiOrderComplaintGetRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.complaint.get"
}

func (r AlibabaHappytripTaxiOrderComplaintGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiOrderComplaintGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiOrderComplaintGetRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaHappytripTaxiOrderComplaintGetRequest) SetCaseId(caseId string) error {
    r.caseId = caseId
    r.Set("case_id", caseId)
    return nil
}

func (r AlibabaHappytripTaxiOrderComplaintGetRequest) GetCaseId() string {
    return r.caseId
}

