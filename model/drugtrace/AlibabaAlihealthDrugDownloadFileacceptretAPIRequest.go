package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugDownloadFileacceptretAPIRequest
企业上传回执 API请求
alibaba.alihealth.drug.download.fileacceptret

拿到企业下载回执，将企业已下载的和未下载成功的条目都相应的改变状态 */
type AlibabaAlihealthDrugDownloadFileacceptretAPIRequest struct {
	model.Params
	// appKey
	_appKeyN string
	// fileResultJson
	_fileResultJson string
}

// New
