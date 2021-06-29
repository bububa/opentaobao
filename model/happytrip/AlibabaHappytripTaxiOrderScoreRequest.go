package happytrip

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单打分和评价 APIRequest
alibaba.happytrip.taxi.order.score

对司机进行评分，只有订单结束后，才能进行。
*/
type AlibabaHappytripTaxiOrderScoreRequest struct {
    model.Params

    // 订单id
    orderId   string 

    // 司机评分 星级(1-5)
    level   int64 

    // 司机评价最多40个汉字
    comment   string 

}

func NewAlibabaHappytripTaxiOrderScoreRequest() *AlibabaHappytripTaxiOrderScoreRequest{
    return &AlibabaHappytripTaxiOrderScoreRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaHappytripTaxiOrderScoreRequest) GetApiMethodName() string {
    return "alibaba.happytrip.taxi.order.score"
}

func (r AlibabaHappytripTaxiOrderScoreRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaHappytripTaxiOrderScoreRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaHappytripTaxiOrderScoreRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaHappytripTaxiOrderScoreRequest) SetLevel(level int64) error {
    r.level = level
    r.Set("level", level)
    return nil
}

func (r AlibabaHappytripTaxiOrderScoreRequest) GetLevel() int64 {
    return r.level
}

func (r *AlibabaHappytripTaxiOrderScoreRequest) SetComment(comment string) error {
    r.comment = comment
    r.Set("comment", comment)
    return nil
}

func (r AlibabaHappytripTaxiOrderScoreRequest) GetComment() string {
    return r.comment
}

