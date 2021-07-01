package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytFiledownloadAPIRequest
处理失败单据下载 API请求
alibaba.alihealth.drug.kyt.filedownload

处理失败单据下载 */
type AlibabaAlihealthDrugKytFiledownloadAPIRequest struct {
	model.Params
	// 企业ID
	_refUserId string
	// 文件地址
	_url string
	// 单据类型
	_billType string
	// 单据队列ID
	_billQueueId string
}

// New
