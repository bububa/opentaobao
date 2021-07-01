package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytSynonymauthsAPIRequest
物流企业查询货主企业信息 API请求
alibaba.alihealth.drug.kyt.synonymauths

物流企业查询货主企业信息 */
type AlibabaAlihealthDrugKytSynonymauthsAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 企业名称
	_entName string
	// 货主自定义编号
	_synOwnEntId string
	// 页码
	_pageSize int64
	// 页面大小
	_page int64
}

// New
