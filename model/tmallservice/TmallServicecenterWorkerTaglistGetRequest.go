package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取工人标签 API请求
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

// 初始化TmallServicecenterWorkerTaglistGetRequest对象
func NewTmallServicecenterWorkerTaglistGetRequest() *TmallServicecenterWorkerTaglistGetRequest{
    return &TmallServicecenterWorkerTaglistGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkerTaglistGetRequest) GetApiMethodName() string {
    return "tmall.servicecenter.worker.taglist.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkerTaglistGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IdNumber Setter
// 工人注册勤鸽时的身份证
func (r *TmallServicecenterWorkerTaglistGetRequest) SetIdNumber(idNumber string) error {
    r.idNumber = idNumber
    r.Set("id_number", idNumber)
    return nil
}

// IdNumber Getter
func (r TmallServicecenterWorkerTaglistGetRequest) GetIdNumber() string {
    return r.idNumber
}
// Mobile Setter
// 工人注册勤鸽时的手机号码
func (r *TmallServicecenterWorkerTaglistGetRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r TmallServicecenterWorkerTaglistGetRequest) GetMobile() string {
    return r.mobile
}
