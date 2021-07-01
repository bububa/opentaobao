package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytListauthsAPIRequest
企业搜索自己授权的物流企业 API请求
alibaba.alihealth.drug.kyt.listauths

企业搜索自己授权的物流企业 */
type AlibabaAlihealthDrugKytListauthsAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 企业名称
	_entName string
	// 页码
	_page int64
	// 页大小
	_pageSize int64
}

// New
