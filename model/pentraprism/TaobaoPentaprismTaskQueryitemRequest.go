package pentraprism

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询任务当前进度 APIRequest
taobao.pentaprism.task.queryitem

外网用户查询五棱镜任务系统当前进度
*/
type TaobaoPentaprismTaskQueryitemRequest struct {
    model.Params

    // TOP接口标准入参
    openPo   *OpenTaskPo 

}

func NewTaobaoPentaprismTaskQueryitemRequest() *TaobaoPentaprismTaskQueryitemRequest{
    return &TaobaoPentaprismTaskQueryitemRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoPentaprismTaskQueryitemRequest) GetApiMethodName() string {
    return "taobao.pentaprism.task.queryitem"
}

func (r TaobaoPentaprismTaskQueryitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoPentaprismTaskQueryitemRequest) SetOpenPo(openPo *OpenTaskPo) error {
    r.openPo = openPo
    r.Set("open_po", openPo)
    return nil
}

func (r TaobaoPentaprismTaskQueryitemRequest) GetOpenPo() *OpenTaskPo {
    return r.openPo
}

