package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取工人标签 APIRequest
tmall.servicecenter.worker.taglist.get

服务商获取对应工人的标签
*/
type TmallServicecenterWorkerTaglistGetRequest struct {
    model.Params

    // 工人注册勤鸽时的身份证
    idNumber   string 

    // 工人注册勤鸽时的手机号码
    mobile   string 

}

func NewTmallServicecenterWorkerTaglistGetRequest() *TmallServicecenterWorkerTaglistGetRequest{
    return &TmallServicecenterWorkerTaglistGetRequest{
        Params: model.NewParams(),
    }
}

func (r TmallServicecenterWorkerTaglistGetRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.taglist.get"
}

func (r TmallServicecenterWorkerTaglistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallServicecenterWorkerTaglistGetRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

func (r TmallServicecenterWorkerTaglistGetRequest) GetIdNumber() string {
    return r.idNumber
}

func (r *TmallServicecenterWorkerTaglistGetRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r TmallServicecenterWorkerTaglistGetRequest) GetMobile() string {
    return r.mobile
}

