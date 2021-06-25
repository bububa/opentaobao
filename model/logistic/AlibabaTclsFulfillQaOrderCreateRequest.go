package logistic

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创单接口 APIRequest
alibaba.tcls.fulfill.qa.order.create

根据历史测试履约单号，复制一个同样镜像的履约单号并下发给大润发仓（api实现已经限制了测试数据）
*/
type AlibabaTclsFulfillQaOrderCreateRequest struct {
    model.Params

    // 原始履约单号
    fulfillOrderId   string 

    // 目标ip
    targetIp   string 

    // 执行人姓名
    creator   string 

    // 执行人工号
    jobNo   string 

}

func NewAlibabaTclsFulfillQaOrderCreateRequest() *AlibabaTclsFulfillQaOrderCreateRequest{
    return &AlibabaTclsFulfillQaOrderCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTclsFulfillQaOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.tcls.fulfill.qa.order.create"
}

func (r AlibabaTclsFulfillQaOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTclsFulfillQaOrderCreateRequest) SetFulfillOrderId(fulfillOrderId string) error {
    r.fulfillOrderId = fulfillOrderId
    r.Set("fulfill_order_id", fulfillOrderId)
    return nil
}

func (r AlibabaTclsFulfillQaOrderCreateRequest) GetFulfillOrderId() string {
    return r.fulfillOrderId
}

func (r *AlibabaTclsFulfillQaOrderCreateRequest) SetTargetIp(targetIp string) error {
    r.targetIp = targetIp
    r.Set("target_ip", targetIp)
    return nil
}

func (r AlibabaTclsFulfillQaOrderCreateRequest) GetTargetIp() string {
    return r.targetIp
}

func (r *AlibabaTclsFulfillQaOrderCreateRequest) SetCreator(creator string) error {
    r.creator = creator
    r.Set("creator", creator)
    return nil
}

func (r AlibabaTclsFulfillQaOrderCreateRequest) GetCreator() string {
    return r.creator
}

func (r *AlibabaTclsFulfillQaOrderCreateRequest) SetJobNo(jobNo string) error {
    r.jobNo = jobNo
    r.Set("job_no", jobNo)
    return nil
}

func (r AlibabaTclsFulfillQaOrderCreateRequest) GetJobNo() string {
    return r.jobNo
}

