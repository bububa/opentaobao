package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest
服务商工人客诉数据上传 API请求
alibaba.dchain.miaoshifu.customer.complaints.put

数字服务供应链平台提供给服务商上传工人客诉数据 */
type AlibabaDchainMiaoshifuCustomerComplaintsPutAPIRequest struct {
	model.Params
	// 服务工人客诉对象
	_workerCustomerComplaintSaveCmd *WorkerCustomerComplaintSaveCmd
}

// New
