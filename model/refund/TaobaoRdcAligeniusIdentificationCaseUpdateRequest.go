package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
鉴定工单信息同步 API请求
taobao.rdc.aligenius.identification.case.update

同步商家鉴定工单信息
*/
type TaobaoRdcAligeniusIdentificationCaseUpdateRequest struct {
    model.Params
    // 请求参数
    _param   *SyncIdentifyRefundCaseDto
}

// 初始化TaobaoRdcAligeniusIdentificationCaseUpdateRequest对象
func NewTaobaoRdcAligeniusIdentificationCaseUpdateRequest() *TaobaoRdcAligeniusIdentificationCaseUpdateRequest{
    return &TaobaoRdcAligeniusIdentificationCaseUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRdcAligeniusIdentificationCaseUpdateRequest) GetApiMethodName() string {
    return "taobao.rdc.aligenius.identification.case.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRdcAligeniusIdentificationCaseUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param Setter
// 请求参数
func (r *TaobaoRdcAligeniusIdentificationCaseUpdateRequest) SetParam(_param *SyncIdentifyRefundCaseDto) error {
    r._param = _param
    r.Set("param", _param)
    return nil
}

// Param Getter
func (r TaobaoRdcAligeniusIdentificationCaseUpdateRequest) GetParam() *SyncIdentifyRefundCaseDto {
    return r._param
}
