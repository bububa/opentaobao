package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugDownloadGetentauthentAPIRequest
获取授权企业列表 API请求
alibaba.alihealth.drug.download.getentauthent

D2D数据落地获取授权企业列表 */
type AlibabaAlihealthDrugDownloadGetentauthentAPIRequest struct {
	model.Params
	// 授权开始时间
	_authBeginDate string
	// 授权结束时间
	_authEndDate string
}

// New
