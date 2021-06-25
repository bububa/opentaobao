package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定工单信息同步 APIRequest
taobao.rdc.aligenius.identification.case.update

同步商家鉴定工单信息
*/
type TaobaoRdcAligeniusIdentificationCaseUpdateRequest struct {
    model.Params

    // 请求参数
    param   *SyncIdentifyRefundCaseDto 

}

func NewTaobaoRdcAligeniusIdentificationCaseUpdateRequest() *TaobaoRdcAligeniusIdentificationCaseUpdateRequest{
    return &TaobaoRdcAligeniusIdentificationCaseUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdcAligeniusIdentificationCaseUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.identification.case.update"
}

func (r TaobaoRdcAligeniusIdentificationCaseUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdcAligeniusIdentificationCaseUpdateRequest) SetParam(param *SyncIdentifyRefundCaseDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoRdcAligeniusIdentificationCaseUpdateRequest) GetParam() *SyncIdentifyRefundCaseDto {
    return r.param
}

