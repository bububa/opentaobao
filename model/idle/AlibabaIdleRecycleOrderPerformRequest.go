package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
回收订单履约V2 API请求
alibaba.idle.recycle.order.perform

闲鱼回收业务中,外部服务商作为买家 需要驱动交易节点变化
*/
type AlibabaIdleRecycleOrderPerformAPIRequest struct {
    model.Params
    // 参数
    _param0   *RecycleOrderSynDTO
}

// 初始化AlibabaIdleRecycleOrderPerformAPIRequest对象
func NewAlibabaIdleRecycleOrderPerformRequest() *AlibabaIdleRecycleOrderPerformAPIRequest{
    return &AlibabaIdleRecycleOrderPerformAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaIdleRecycleOrderPerformAPIRequest) GetApiMethodName() string {
    return "alibaba.idle.recycle.order.perform"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaIdleRecycleOrderPerformAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Param0 Setter
// 参数
func (r *AlibabaIdleRecycleOrderPerformAPIRequest) SetParam0(_param0 *RecycleOrderSynDTO) error {
    r._param0 = _param0
    r.Set("param0", _param0)
    return nil
}

// Param0 Getter
func (r AlibabaIdleRecycleOrderPerformAPIRequest) GetParam0() *RecycleOrderSynDTO {
    return r._param0
}
