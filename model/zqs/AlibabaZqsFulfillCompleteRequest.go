package zqs

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
周期购履约完成接口 API请求
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

// 初始化AlibabaZqsFulfillCompleteRequest对象
func NewAlibabaZqsFulfillCompleteRequest() *AlibabaZqsFulfillCompleteRequest{
    return &AlibabaZqsFulfillCompleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaZqsFulfillCompleteRequest) GetApiMethodName() string {
    return "alibaba.zqs.fulfill.complete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaZqsFulfillCompleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SequenceNo Setter
// 第几期
func (r *AlibabaZqsFulfillCompleteRequest) SetSequenceNo(sequenceNo int64) error {
    r.sequenceNo = sequenceNo
    r.Set("sequence_no", sequenceNo)
    return nil
}

// SequenceNo Getter
func (r AlibabaZqsFulfillCompleteRequest) GetSequenceNo() int64 {
    return r.sequenceNo
}
// MainBizOrderId Setter
// 交易单号
func (r *AlibabaZqsFulfillCompleteRequest) SetMainBizOrderId(mainBizOrderId int64) error {
    r.mainBizOrderId = mainBizOrderId
    r.Set("main_biz_order_id", mainBizOrderId)
    return nil
}

// MainBizOrderId Getter
func (r AlibabaZqsFulfillCompleteRequest) GetMainBizOrderId() int64 {
    return r.mainBizOrderId
}
// SubBizOrderId Setter
// 交易子单号
func (r *AlibabaZqsFulfillCompleteRequest) SetSubBizOrderId(subBizOrderId int64) error {
    r.subBizOrderId = subBizOrderId
    r.Set("sub_biz_order_id", subBizOrderId)
    return nil
}

// SubBizOrderId Getter
func (r AlibabaZqsFulfillCompleteRequest) GetSubBizOrderId() int64 {
    return r.subBizOrderId
}
