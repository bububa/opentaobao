package pentraprism

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
推进单条任务进度 APIRequest
taobao.pentaprism.task.trigger

外网用户推进单条五棱镜任务进度
*/
type TaobaoPentaprismTaskTriggerRequest struct {
    model.Params

    // TOP接口标准入参
    openPo   *OpenTaskPo 

}

func NewTaobaoPentaprismTaskTriggerRequest() *TaobaoPentaprismTaskTriggerRequest{
    return &TaobaoPentaprismTaskTriggerRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPentaprismTaskTriggerRequest) GetApiMethodName() string {
    return "taobao.pentaprism.task.trigger"
}

func (r TaobaoPentaprismTaskTriggerRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPentaprismTaskTriggerRequest) SetOpenPo(openPo *OpenTaskPo) error {
    r.openPo = openPo
    r.Set("open_po", openPo)
    return nil
}

func (r TaobaoPentaprismTaskTriggerRequest) GetOpenPo() *OpenTaskPo {
    return r.openPo
}

