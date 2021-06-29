package elife

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询对账文件地址接口 API请求
taobao.elife.lifecard.recon

查询对账文件地址接口
*/
type TaobaoElifeLifecardReconRequest struct {
    model.Params
    // 对账日期(YYYYMMDD)
    _opDate   string
}

// 初始化TaobaoElifeLifecardReconRequest对象
func NewTaobaoElifeLifecardReconRequest() *TaobaoElifeLifecardReconRequest{
    return &TaobaoElifeLifecardReconRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoElifeLifecardReconRequest) GetApiMethodName() string {
    return "taobao.elife.lifecard.recon"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoElifeLifecardReconRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpDate Setter
// 对账日期(YYYYMMDD)
func (r *TaobaoElifeLifecardReconRequest) SetOpDate(_opDate string) error {
    r._opDate = _opDate
    r.Set("op_date", _opDate)
    return nil
}

// OpDate Getter
func (r TaobaoElifeLifecardReconRequest) GetOpDate() string {
    return r._opDate
}
