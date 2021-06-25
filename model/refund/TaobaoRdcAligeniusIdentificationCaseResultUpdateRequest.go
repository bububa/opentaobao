package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定工单结果同步 APIRequest
taobao.rdc.aligenius.identification.case.result.update

同步鉴定工单结果信息
*/
type TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest struct {
    model.Params

    // 请求参数
    param   *SyncIdentifyRefundCaseResultDto 

}

func NewTaobaoRdcAligeniusIdentificationCaseResultUpdateRequest() *TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest{
    return &TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.identification.case.result.update"
}

func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest) SetParam(param *SyncIdentifyRefundCaseResultDto) error {
    r.param = param
    r.Set("param", param)
    return nil
}

func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest) GetParam() *SyncIdentifyRefundCaseResultDto {
    return r.param
}

