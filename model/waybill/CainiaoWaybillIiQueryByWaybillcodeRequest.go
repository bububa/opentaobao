package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过面单号查询面单信息 API请求
cainiao.waybill.ii.query.by.waybillcode

通过面单号查看面单号的当前状态，如签收、发货、失效等。
*/
type CainiaoWaybillIiQueryByWaybillcodeRequest struct {
    model.Params
    // 系统自动生成
    _paramList   []WaybillDetailQueryByWaybillCodeRequest
}

// 初始化CainiaoWaybillIiQueryByWaybillcodeRequest对象
func NewCainiaoWaybillIiQueryByWaybillcodeRequest() *CainiaoWaybillIiQueryByWaybillcodeRequest{
    return &CainiaoWaybillIiQueryByWaybillcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiQueryByWaybillcodeRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.query.by.waybillcode"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiQueryByWaybillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamList Setter
// 系统自动生成
func (r *CainiaoWaybillIiQueryByWaybillcodeRequest) SetParamList(_paramList []WaybillDetailQueryByWaybillCodeRequest) error {
    r._paramList = _paramList
    r.Set("param_list", _paramList)
    return nil
}

// ParamList Getter
func (r CainiaoWaybillIiQueryByWaybillcodeRequest) GetParamList() []WaybillDetailQueryByWaybillCodeRequest {
    return r._paramList
}
