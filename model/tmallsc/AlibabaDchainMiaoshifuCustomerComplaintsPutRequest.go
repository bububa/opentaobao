package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商工人客诉数据上传 APIRequest
alibaba.dchain.miaoshifu.customer.complaints.put

数字服务供应链平台提供给服务商上传工人客诉数据
*/
type AlibabaDchainMiaoshifuCustomerComplaintsPutRequest struct {
    model.Params

    // 服务工人客诉对象
    workerCustomerComplaintSaveCmd   *WorkerCustomerComplaintSaveCmd 

}

func NewAlibabaDchainMiaoshifuCustomerComplaintsPutRequest() *AlibabaDchainMiaoshifuCustomerComplaintsPutRequest{
    return &AlibabaDchainMiaoshifuCustomerComplaintsPutRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaDchainMiaoshifuCustomerComplaintsPutRequest) GetApiMethodName() string {
    return "alibaba.dchain.miaoshifu.customer.complaints.put"
}

func (r AlibabaDchainMiaoshifuCustomerComplaintsPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaDchainMiaoshifuCustomerComplaintsPutRequest) SetWorkerCustomerComplaintSaveCmd(workerCustomerComplaintSaveCmd *WorkerCustomerComplaintSaveCmd) error {
    r.workerCustomerComplaintSaveCmd = workerCustomerComplaintSaveCmd
    r.Set("worker_customer_complaint_save_cmd", workerCustomerComplaintSaveCmd)
    return nil
}

func (r AlibabaDchainMiaoshifuCustomerComplaintsPutRequest) GetWorkerCustomerComplaintSaveCmd() *WorkerCustomerComplaintSaveCmd {
    return r.workerCustomerComplaintSaveCmd
}

