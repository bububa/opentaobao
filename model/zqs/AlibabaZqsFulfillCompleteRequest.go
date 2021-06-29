package zqs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
周期购履约完成接口 APIRequest
alibaba.zqs.fulfill.complete

周期购履约完成接口
*/
type AlibabaZqsFulfillCompleteRequest struct {
    model.Params

    // 第几期
    sequenceNo   int64 

    // 交易单号
    mainBizOrderId   int64 

    // 交易子单号
    subBizOrderId   int64 

}

func NewAlibabaZqsFulfillCompleteRequest() *AlibabaZqsFulfillCompleteRequest{
    return &AlibabaZqsFulfillCompleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaZqsFulfillCompleteRequest) GetApiMethodName() string {
    return "alibaba.zqs.fulfill.complete"
}

func (r AlibabaZqsFulfillCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaZqsFulfillCompleteRequest) SetSequenceNo(sequenceNo int64) error {
    r.sequenceNo = sequenceNo
    r.Set("sequence_no", sequenceNo)
    return nil
}

func (r AlibabaZqsFulfillCompleteRequest) GetSequenceNo() int64 {
    return r.sequenceNo
}

func (r *AlibabaZqsFulfillCompleteRequest) SetMainBizOrderId(mainBizOrderId int64) error {
    r.mainBizOrderId = mainBizOrderId
    r.Set("main_biz_order_id", mainBizOrderId)
    return nil
}

func (r AlibabaZqsFulfillCompleteRequest) GetMainBizOrderId() int64 {
    return r.mainBizOrderId
}

func (r *AlibabaZqsFulfillCompleteRequest) SetSubBizOrderId(subBizOrderId int64) error {
    r.subBizOrderId = subBizOrderId
    r.Set("sub_biz_order_id", subBizOrderId)
    return nil
}

func (r AlibabaZqsFulfillCompleteRequest) GetSubBizOrderId() int64 {
    return r.subBizOrderId
}

