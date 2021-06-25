package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询任务类工单是否退款 APIRequest
tmall.servicecenter.task.queryrefund

查询任务类工单是否退款
*/
type TmallServicecenterTaskQueryrefundRequest struct {
    model.Params

    // 工单id列表
    workcardList   []Number 

}

func NewTmallServicecenterTaskQueryrefundRequest() *TmallServicecenterTaskQueryrefundRequest{
    return &TmallServicecenterTaskQueryrefundRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterTaskQueryrefundRequest) GetApiMethodName() string {
    return "tmall.servicecenter.task.queryrefund"
}

func (r TmallServicecenterTaskQueryrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterTaskQueryrefundRequest) SetWorkcardList(workcardList []Number) error {
    r.workcardList = workcardList
    r.Set("workcard_list", workcardList)
    return nil
}

func (r TmallServicecenterTaskQueryrefundRequest) GetWorkcardList() []Number {
    return r.workcardList
}

