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
type TaobaoTbkDgVegasTljInstanceReportRequest struct {
    model.Params
    // 实例ID
    rightsId   string
}

// 初始化TaobaoTbkDgVegasTljInstanceReportRequest对象
func NewTaobaoTbkDgVegasTljInstanceReportRequest() *TaobaoTbkDgVegasTljInstanceReportRequest{
    return &TaobaoTbkDgVegasTljInstanceReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasTljInstanceReportRequest) GetApiMethodName() string {
    return "taobao.tbk.dg.vegas.tlj.instance.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasTljInstanceReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RightsId Setter
// 实例ID
func (r *TaobaoTbkDgVegasTljInstanceReportRequest) SetRightsId(rightsId string) error {
    r.rightsId = rightsId
    r.Set("rights_id", rightsId)
    return nil
}

// RightsId Getter
func (r TaobaoTbkDgVegasTljInstanceReportRequest) GetRightsId() string {
    return r.rightsId
}
