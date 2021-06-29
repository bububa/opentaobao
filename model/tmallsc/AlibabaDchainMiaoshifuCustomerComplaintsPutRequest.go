package tmallsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
服务商工人客诉数据上传 API请求
alibaba.dchain.miaoshifu.customer.complaints.put

数字服务供应链平台提供给服务商上传工人客诉数据
*/
type AlibabaDchainMiaoshifuCustomerComplaintsPutRequest struct {
    model.Params
    // 服务工人客诉对象
    _workerCustomerComplaintSaveCmd   *WorkerCustomerComplaintSaveCmd
}

// 初始化AlibabaDchainMiaoshifuCustomerComplaintsPutRequest对象
func NewAlibabaDchainMiaoshifuCustomerComplaintsPutRequest() *AlibabaDchainMiaoshifuCustomerComplaintsPutRequest{
    return &AlibabaDchainMiaoshifuCustomerComplaintsPutRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDchainMiaoshifuCustomerComplaintsPutRequest) GetApiMethodName() string {
    return "alibaba.dchain.miaoshifu.customer.complaints.put"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDchainMiaoshifuCustomerComplaintsPutRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// WorkerCustomerComplaintSaveCmd Setter
// 服务工人客诉对象
func (r *AlibabaDchainMiaoshifuCustomerComplaintsPutRequest) SetWorkerCustomerComplaintSaveCmd(_workerCustomerComplaintSaveCmd *WorkerCustomerComplaintSaveCmd) error {
    r._workerCustomerComplaintSaveCmd = _workerCustomerComplaintSaveCmd
    r.Set("worker_customer_complaint_save_cmd", _workerCustomerComplaintSaveCmd)
    return nil
}

// WorkerCustomerComplaintSaveCmd Getter
func (r AlibabaDchainMiaoshifuCustomerComplaintsPutRequest) GetWorkerCustomerComplaintSaveCmd() *WorkerCustomerComplaintSaveCmd {
    return r._workerCustomerComplaintSaveCmd
}
