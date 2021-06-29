package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定工单结果同步 API请求
taobao.rdc.aligenius.identification.case.result.update

同步鉴定工单结果信息
*/
type TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest struct {
    model.Params
    // 请求参数
    _param   *SyncIdentifyRefundCaseResultDto
}

// 初始化TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest对象
func NewTaobaoRdcAligeniusIdentificationCaseResultUpdateRequest() *TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest{
    return &TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.identification.case.result.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest) SetParam(_param *SyncIdentifyRefundCaseResultDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoRdcAligeniusIdentificationCaseResultUpdateRequest) GetParam() *SyncIdentifyRefundCaseResultDto {
    return r._param
}
