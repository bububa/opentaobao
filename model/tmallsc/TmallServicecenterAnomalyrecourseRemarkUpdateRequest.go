package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫服务平台一键求助单服务商备注更新接口 APIRequest
tmall.servicecenter.anomalyrecourse.remark.update

一键求助服务商可以回传备注
*/
type TmallServicecenterAnomalyrecourseRemarkUpdateRequest struct {
    model.Params

    // 需要更新的一键求助单id
    id   int64 

    // 需要更新的服务商备注
    remark   string 

}

func NewTmallServicecenterAnomalyrecourseRemarkUpdateRequest() *TmallServicecenterAnomalyrecourseRemarkUpdateRequest{
    return &TmallServicecenterAnomalyrecourseRemarkUpdateRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterAnomalyrecourseRemarkUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.anomalyrecourse.remark.update"
}

func (r TmallServicecenterAnomalyrecourseRemarkUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterAnomalyrecourseRemarkUpdateRequest) SetId(id int64) error {
    r.id = id
    r.Set("id", id)
    return nil
}

func (r TmallServicecenterAnomalyrecourseRemarkUpdateRequest) GetId() int64 {
    return r.id
}

func (r *TmallServicecenterAnomalyrecourseRemarkUpdateRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

func (r TmallServicecenterAnomalyrecourseRemarkUpdateRequest) GetRemark() string {
    return r.remark
}

