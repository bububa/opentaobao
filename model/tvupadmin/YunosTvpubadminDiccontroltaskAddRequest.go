package tvupadmin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增停开服任务 APIRequest
yunos.tvpubadmin.diccontroltask.add

新增停开服任务
*/
type YunosTvpubadminDiccontroltaskAddRequest struct {
    model.Params

    // 任务信息
    task   *DicControlTaskDO 

}

func NewYunosTvpubadminDiccontroltaskAddRequest() *YunosTvpubadminDiccontroltaskAddRequest{
    return &YunosTvpubadminDiccontroltaskAddRequest{
        Params: model.NewParams(),
    }
}

func (r YunosTvpubadminDiccontroltaskAddRequest) GetApiMethodName() string {
    return "yunos.tvpubadmin.diccontroltask.add"
}

func (r YunosTvpubadminDiccontroltaskAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosTvpubadminDiccontroltaskAddRequest) SetTask(task *DicControlTaskDO) error {
    r.task = task
    r.Set("task", task)
    return nil
}

func (r YunosTvpubadminDiccontroltaskAddRequest) GetTask() *DicControlTaskDO {
    return r.task
}

