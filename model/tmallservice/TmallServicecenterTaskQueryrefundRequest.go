package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询任务类工单是否退款 API请求
tmall.servicecenter.task.queryrefund

查询任务类工单是否退款
*/
type TmallServicecenterTaskQueryrefundRequest struct {
    model.Params
    // 工单id列表
    workcardList   []int64
}

// 初始化TmallServicecenterTaskQueryrefundRequest对象
func NewTmallServicecenterTaskQueryrefundRequest() *TmallServicecenterTaskQueryrefundRequest{
    return &TmallServicecenterTaskQueryrefundRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterTaskQueryrefundRequest) GetApiMethodName() string {
    return "tmall.servicecenter.task.queryrefund"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterTaskQueryrefundRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkcardList Setter
// 工单id列表
func (r *TmallServicecenterTaskQueryrefundRequest) SetWorkcardList(workcardList []int64) error {
    r.workcardList = workcardList
    r.Set("workcard_list", workcardList)
    return nil
}

// WorkcardList Getter
func (r TmallServicecenterTaskQueryrefundRequest) GetWorkcardList() []int64 {
    return r.workcardList
}
