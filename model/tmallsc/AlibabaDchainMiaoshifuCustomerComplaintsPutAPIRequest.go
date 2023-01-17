package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest 服务商工人客诉数据上传 API请求
// alibaba.dchain.miaoshifu.customer.complaints.put
//
// 数字服务供应链平台提供给服务商上传工人客诉数据
type AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest struct {
	model.Params
	// 服务工人客诉对象
	_workerCustomerComplaintSaveCmd *WorkerCustomerComplaintSaveCmd
}

// NewAlibabaDchainMiaoshifuCustomerComplaintsPutRequest 初始化AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest对象
func NewAlibabaDchainMiaoshifuCustomerComplaintsPutRequest() *AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest {
	return &AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.miaoshifu.customer.complaints.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerCustomerComplaintSaveCmd is WorkerCustomerComplaintSaveCmd Setter
// 服务工人客诉对象
func (r *AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest) SetWorkerCustomerComplaintSaveCmd(_workerCustomerComplaintSaveCmd *WorkerCustomerComplaintSaveCmd) error {
	r._workerCustomerComplaintSaveCmd = _workerCustomerComplaintSaveCmd
	r.Set("worker_customer_complaint_save_cmd", _workerCustomerComplaintSaveCmd)
	return nil
}

// GetWorkerCustomerComplaintSaveCmd WorkerCustomerComplaintSaveCmd Getter
func (r AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest) GetWorkerCustomerComplaintSaveCmd() *WorkerCustomerComplaintSaveCmd {
	return r._workerCustomerComplaintSaveCmd
}
