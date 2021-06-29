package tmallservice

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
喵师傅工人任务批量完成接口 API请求
tmall.servicecenter.msf.tasks.update

喵师傅工人任务批量完成接口
*/
type TmallServicecenterMsfTasksUpdateRequest struct {
    model.Params
    // 工人手机号
    workerMobile   int64
    // 服务编码
    serviceCode   string
    // 调用来源
    source   string
    // 子订单号列表。最多100个
    bizOrderIds   []int64
    // 周期序号。必须与子订单列表对应
    seqs   []int64
}

// 初始化TmallServicecenterMsfTasksUpdateRequest对象
func NewTmallServicecenterMsfTasksUpdateRequest() *TmallServicecenterMsfTasksUpdateRequest{
    return &TmallServicecenterMsfTasksUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallServicecenterMsfTasksUpdateRequest) GetApiMethodName() string {
    return "tmall.servicecenter.msf.tasks.update"
}

// IRequest interface 方法, 获取API参数
func (r TmallServicecenterMsfTasksUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkerMobile Setter
// 工人手机号
func (r *TmallServicecenterMsfTasksUpdateRequest) SetWorkerMobile(workerMobile int64) error {
    r.workerMobile = workerMobile
    r.Set("worker_mobile", workerMobile)
    return nil
}

// WorkerMobile Getter
func (r TmallServicecenterMsfTasksUpdateRequest) GetWorkerMobile() int64 {
    return r.workerMobile
}
// ServiceCode Setter
// 服务编码
func (r *TmallServicecenterMsfTasksUpdateRequest) SetServiceCode(serviceCode string) error {
    r.serviceCode = serviceCode
    r.Set("service_code", serviceCode)
    return nil
}

// ServiceCode Getter
func (r TmallServicecenterMsfTasksUpdateRequest) GetServiceCode() string {
    return r.serviceCode
}
// Source Setter
// 调用来源
func (r *TmallServicecenterMsfTasksUpdateRequest) SetSource(source string) error {
    r.source = source
    r.Set("source", source)
    return nil
}

// Source Getter
func (r TmallServicecenterMsfTasksUpdateRequest) GetSource() string {
    return r.source
}
// BizOrderIds Setter
// 子订单号列表。最多100个
func (r *TmallServicecenterMsfTasksUpdateRequest) SetBizOrderIds(bizOrderIds []int64) error {
    r.bizOrderIds = bizOrderIds
    r.Set("biz_order_ids", bizOrderIds)
    return nil
}

// BizOrderIds Getter
func (r TmallServicecenterMsfTasksUpdateRequest) GetBizOrderIds() []int64 {
    return r.bizOrderIds
}
// Seqs Setter
// 周期序号。必须与子订单列表对应
func (r *TmallServicecenterMsfTasksUpdateRequest) SetSeqs(seqs []int64) error {
    r.seqs = seqs
    r.Set("seqs", seqs)
    return nil
}

// Seqs Getter
func (r TmallServicecenterMsfTasksUpdateRequest) GetSeqs() []int64 {
    return r.seqs
}
