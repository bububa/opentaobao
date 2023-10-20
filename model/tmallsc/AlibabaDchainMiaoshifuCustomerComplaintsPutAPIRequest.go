package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainmiaoshifucustomercomplaintsputAPIRequest 服务商工人客诉数据上传 API请求
// alibaba.dchain.miaoshifu.customer.complaints.put
//
// 数字服务供应链平台提供给服务商上传工人客诉数据
type AlibabadchainmiaoshifucustomercomplaintsputAPIRequest struct {
	model.Params
	// 服务工人客诉对象
	_workerCustomerComplaintSaveCmd *WorkerCustomerComplaintSaveCmd
}

// NewAlibabadchainmiaoshifucustomercomplaintsputRequest 初始化AlibabadchainmiaoshifucustomercomplaintsputAPIRequest对象
func NewAlibabadchainmiaoshifucustomercomplaintsputRequest() *AlibabadchainmiaoshifucustomercomplaintsputAPIRequest {
	return &AlibabadchainmiaoshifucustomercomplaintsputAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainmiaoshifucustomercomplaintsputAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.miaoshifu.customer.complaints.put"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainmiaoshifucustomercomplaintsputAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainmiaoshifucustomercomplaintsputAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetWorkerCustomerComplaintSaveCmd is WorkerCustomerComplaintSaveCmd Setter
// 服务工人客诉对象
func (r *AlibabadchainmiaoshifucustomercomplaintsputAPIRequest) SetWorkerCustomerComplaintSaveCmd(_workerCustomerComplaintSaveCmd *WorkerCustomerComplaintSaveCmd) error {
	r._workerCustomerComplaintSaveCmd = _workerCustomerComplaintSaveCmd
	r.Set("worker_customer_complaint_save_cmd", _workerCustomerComplaintSaveCmd)
	return nil
}

// GetWorkerCustomerComplaintSaveCmd WorkerCustomerComplaintSaveCmd Getter
func (r AlibabadchainmiaoshifucustomercomplaintsputAPIRequest) GetWorkerCustomerComplaintSaveCmd() *WorkerCustomerComplaintSaveCmd {
	return r._workerCustomerComplaintSaveCmd
}
