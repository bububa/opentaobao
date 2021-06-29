package pentraprism

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推进单条任务进度 API请求
taobao.pentaprism.task.trigger

外网用户推进单条五棱镜任务进度
*/
type TaobaoPentaprismTaskTriggerRequest struct {
    model.Params
    // TOP接口标准入参
    openPo   *OpenTaskPo
}

// 初始化TaobaoPentaprismTaskTriggerRequest对象
func NewTaobaoPentaprismTaskTriggerRequest() *TaobaoPentaprismTaskTriggerRequest{
    return &TaobaoPentaprismTaskTriggerRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPentaprismTaskTriggerRequest) GetApiMethodName() string {
    return "taobao.pentaprism.task.trigger"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPentaprismTaskTriggerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenPo Setter
// TOP接口标准入参
func (r *TaobaoPentaprismTaskTriggerRequest) SetOpenPo(openPo *OpenTaskPo) error {
    r.openPo = openPo
    r.Set("open_po", openPo)
    return nil
}

// OpenPo Getter
func (r TaobaoPentaprismTaskTriggerRequest) GetOpenPo() *OpenTaskPo {
    return r.openPo
}
