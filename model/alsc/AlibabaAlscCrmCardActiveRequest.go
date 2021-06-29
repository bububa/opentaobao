package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标准激活卡 API请求
alibaba.alsc.crm.card.active

激活卡
*/
type AlibabaAlscCrmCardActiveRequest struct {
    model.Params
    // 请求参数
    _paramActiveCardOpenReq   *ActiveCardOpenReq
}

// 初始化AlibabaAlscCrmCardActiveRequest对象
func NewAlibabaAlscCrmCardActiveRequest() *AlibabaAlscCrmCardActiveRequest{
    return &AlibabaAlscCrmCardActiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardActiveRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.active"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardActiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamActiveCardOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmCardActiveRequest) SetParamActiveCardOpenReq(_paramActiveCardOpenReq *ActiveCardOpenReq) error {
    r._paramActiveCardOpenReq = _paramActiveCardOpenReq
    r.Set("param_active_card_open_req", _paramActiveCardOpenReq)
    return nil
}

// ParamActiveCardOpenReq Getter
func (r AlibabaAlscCrmCardActiveRequest) GetParamActiveCardOpenReq() *ActiveCardOpenReq {
    return r._paramActiveCardOpenReq
}
