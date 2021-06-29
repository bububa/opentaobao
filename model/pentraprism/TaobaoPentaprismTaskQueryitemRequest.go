package pentraprism

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询任务当前进度 API请求
taobao.pentaprism.task.queryitem

外网用户查询五棱镜任务系统当前进度
*/
type TaobaoPentaprismTaskQueryitemRequest struct {
    model.Params
    // TOP接口标准入参
    openPo   *OpenTaskPo
}

// 初始化TaobaoPentaprismTaskQueryitemRequest对象
func NewTaobaoPentaprismTaskQueryitemRequest() *TaobaoPentaprismTaskQueryitemRequest{
    return &TaobaoPentaprismTaskQueryitemRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoPentaprismTaskQueryitemRequest) GetApiMethodName() string {
    return "taobao.pentaprism.task.queryitem"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoPentaprismTaskQueryitemRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OpenPo Setter
// TOP接口标准入参
func (r *TaobaoPentaprismTaskQueryitemRequest) SetOpenPo(openPo *OpenTaskPo) error {
    r.openPo = openPo
    r.Set("open_po", openPo)
    return nil
}

// OpenPo Getter
func (r TaobaoPentaprismTaskQueryitemRequest) GetOpenPo() *OpenTaskPo {
    return r.openPo
}
