package tbk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
淘宝客-推广者-淘礼金发放及使用报表 API请求
taobao.tbk.dg.vegas.tlj.instance.report

淘礼金实例维度相关报表数据查询
*/
type TaobaoTbkDgVegasTljInstanceReportAPIRequest struct {
    model.Params
    // 实例ID
    _rightsId   string
}

// 初始化TaobaoTbkDgVegasTljInstanceReportAPIRequest对象
func NewTaobaoTbkDgVegasTljInstanceReportRequest() *TaobaoTbkDgVegasTljInstanceReportAPIRequest{
    return &TaobaoTbkDgVegasTljInstanceReportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasTljInstanceReportAPIRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.vegas.tlj.instance.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasTljInstanceReportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RightsId Setter
// 实例ID
func (r *TaobaoTbkDgVegasTljInstanceReportAPIRequest) SetRightsId(_rightsId string) error {
    r._rightsId = _rightsId
    r.Set("rights_id", _rightsId)
    return nil
}

// RightsId Getter
func (r TaobaoTbkDgVegasTljInstanceReportAPIRequest) GetRightsId() string {
    return r._rightsId
}
