package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增停开服任务 API请求
yunos.tvpubadmin.diccontroltask.add

新增停开服任务
*/
type YunosTvpubadminDiccontroltaskAddRequest struct {
    model.Params
    // 任务信息
    _task   *DicControlTaskDO
}

// 初始化YunosTvpubadminDiccontroltaskAddRequest对象
func NewYunosTvpubadminDiccontroltaskAddRequest() *YunosTvpubadminDiccontroltaskAddRequest{
    return &YunosTvpubadminDiccontroltaskAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskAddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.diccontroltask.add"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Task Setter
// 任务信息
func (r *YunosTvpubadminDiccontroltaskAddRequest) SetTask(_task *DicControlTaskDO) error {
    r._task = _task
    r.Set("task", _task)
    return nil
}

// Task Getter
func (r YunosTvpubadminDiccontroltaskAddRequest) GetTask() *DicControlTaskDO {
    return r._task
}
