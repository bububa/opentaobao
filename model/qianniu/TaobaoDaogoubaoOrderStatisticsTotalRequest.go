package qianniu

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
销售订单总额统计 API请求
taobao.daogoubao.order.statistics.total

对接千牛端数字中心
*/
type TaobaoDaogoubaoOrderStatisticsTotalRequest struct {
    model.Params
    // 调试时用的传入id
    _debugId   string
    // 需要的字段名
    _field   string
}

// 初始化TaobaoDaogoubaoOrderStatisticsTotalRequest对象
func NewTaobaoDaogoubaoOrderStatisticsTotalRequest() *TaobaoDaogoubaoOrderStatisticsTotalRequest{
    return &TaobaoDaogoubaoOrderStatisticsTotalRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoDaogoubaoOrderStatisticsTotalRequest) GetApiMethodName() string {
    return "taobao.daogoubao.order.statistics.total"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoDaogoubaoOrderStatisticsTotalRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DebugId Setter
// 调试时用的传入id
func (r *TaobaoDaogoubaoOrderStatisticsTotalRequest) SetDebugId(_debugId string) error {
    r._debugId = _debugId
    r.Set("debug_id", _debugId)
    return nil
}

// DebugId Getter
func (r TaobaoDaogoubaoOrderStatisticsTotalRequest) GetDebugId() string {
    return r._debugId
}
// Field Setter
// 需要的字段名
func (r *TaobaoDaogoubaoOrderStatisticsTotalRequest) SetField(_field string) error {
    r._field = _field
    r.Set("field", _field)
    return nil
}

// Field Getter
func (r TaobaoDaogoubaoOrderStatisticsTotalRequest) GetField() string {
    return r._field
}
