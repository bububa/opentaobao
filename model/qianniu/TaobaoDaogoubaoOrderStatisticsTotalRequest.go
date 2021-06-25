package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售订单总额统计 APIRequest
taobao.daogoubao.order.statistics.total

对接千牛端数字中心
*/
type TaobaoDaogoubaoOrderStatisticsTotalRequest struct {
    model.Params

    // 调试时用的传入id
    debugId   string 

    // 需要的字段名
    field   string 

}

func NewTaobaoDaogoubaoOrderStatisticsTotalRequest() *TaobaoDaogoubaoOrderStatisticsTotalRequest{
    return &TaobaoDaogoubaoOrderStatisticsTotalRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoDaogoubaoOrderStatisticsTotalRequest) GetApiMethodName() string {
    return "taobao.daogoubao.order.statistics.total"
}

func (r TaobaoDaogoubaoOrderStatisticsTotalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoDaogoubaoOrderStatisticsTotalRequest) SetDebugId(debugId string) error {
    r.debugId = debugId
    r.Set("debug_id", debugId)
    return nil
}

func (r TaobaoDaogoubaoOrderStatisticsTotalRequest) GetDebugId() string {
    return r.debugId
}

func (r *TaobaoDaogoubaoOrderStatisticsTotalRequest) SetField(field string) error {
    r.field = field
    r.Set("field", field)
    return nil
}

func (r TaobaoDaogoubaoOrderStatisticsTotalRequest) GetField() string {
    return r.field
}

