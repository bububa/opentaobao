package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
标准开卡流程 API请求
alibaba.alsc.crm.card.open

标准开卡流程
*/
type AlibabaAlscCrmCardOpenRequest struct {
    model.Params
    // 开卡参数
    _paramOpenCardStandardOpenReq   *OpenCardStandardOpenReq
}

// 初始化AlibabaAlscCrmCardOpenRequest对象
func NewAlibabaAlscCrmCardOpenRequest() *AlibabaAlscCrmCardOpenRequest{
    return &AlibabaAlscCrmCardOpenRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardOpenRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.card.open"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardOpenRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOpenCardStandardOpenReq Setter
// 开卡参数
func (r *AlibabaAlscCrmCardOpenRequest) SetParamOpenCardStandardOpenReq(_paramOpenCardStandardOpenReq *OpenCardStandardOpenReq) error {
    r._paramOpenCardStandardOpenReq = _paramOpenCardStandardOpenReq
    r.Set("param_open_card_standard_open_req", _paramOpenCardStandardOpenReq)
    return nil
}

// ParamOpenCardStandardOpenReq Getter
func (r AlibabaAlscCrmCardOpenRequest) GetParamOpenCardStandardOpenReq() *OpenCardStandardOpenReq {
    return r._paramOpenCardStandardOpenReq
}
