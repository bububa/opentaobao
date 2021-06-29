package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
绑定物理卡 API请求
alibaba.alsc.crm.card.bindcard

将会员卡和实例物理卡绑定在一起
*/
type AlibabaAlscCrmCardBindcardRequest struct {
    model.Params
    // 请求参数
    paramBindPhysicalCardOpenReq   *BindPhysicalCardOpenReq
}

// 初始化AlibabaAlscCrmCardBindcardRequest对象
func NewAlibabaAlscCrmCardBindcardRequest() *AlibabaAlscCrmCardBindcardRequest{
    return &AlibabaAlscCrmCardBindcardRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBindcardRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.bindcard"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBindcardRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBindPhysicalCardOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmCardBindcardRequest) SetParamBindPhysicalCardOpenReq(paramBindPhysicalCardOpenReq *BindPhysicalCardOpenReq) error {
    r.paramBindPhysicalCardOpenReq = paramBindPhysicalCardOpenReq
    r.Set("param_bind_physical_card_open_req", paramBindPhysicalCardOpenReq)
    return nil
}

// ParamBindPhysicalCardOpenReq Getter
func (r AlibabaAlscCrmCardBindcardRequest) GetParamBindPhysicalCardOpenReq() *BindPhysicalCardOpenReq {
    return r.paramBindPhysicalCardOpenReq
}
