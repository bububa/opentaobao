package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugDownloadEntlistAPIRequest
企业下载列表 API请求
alibaba.alihealth.drug.download.entlist

获取企业的下载文件列表 */
type AlibabaAlihealthDrugDownloadEntlistAPIRequest struct {
	model.Params
	// appKey
	_appKeyN string
}

// New
