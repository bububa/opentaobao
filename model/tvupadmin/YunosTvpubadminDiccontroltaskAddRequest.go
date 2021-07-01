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
type YunosTvpubadminDiccontroltaskAddAPIRequest struct {
    model.Params
    // 任务信息
    _task   *DicControlTaskDO
}

// 初始化YunosTvpubadminDiccontroltaskAddAPIRequest对象
func NewYunosTvpubadminDiccontroltaskAddRequest() *YunosTvpubadminDiccontroltaskAddAPIRequest{
    return &YunosTvpubadminDiccontroltaskAddAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosTvpubadminDiccontroltaskAddAPIRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.diccontroltask.add"
}

// IRequest interface 方法, 获取API参数
func (r YunosTvpubadminDiccontroltaskAddAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Task Setter
// 任务信息
func (r *YunosTvpubadminDiccontroltaskAddAPIRequest) SetTask(_task *DicControlTaskDO) error {
    r._task = _task
    r.Set("task", _task)
    return nil
}

// Task Getter
func (r YunosTvpubadminDiccontroltaskAddAPIRequest) GetTask() *DicControlTaskDO {
    return r._task
}
