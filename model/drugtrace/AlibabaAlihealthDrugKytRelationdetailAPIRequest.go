package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytRelationdetailAPIRequest
关联关系处理详情 API请求
alibaba.alihealth.drug.kyt.relationdetail

关联关系处理详情 */
type AlibabaAlihealthDrugKytRelationdetailAPIRequest struct {
	model.Params
	// 码激活文件上传信息标识
	_codeActiveInfoId string
	// 企业ID
	_refEntId string
	// 客户端ID【默认写2】
	_clientType string
}

// New
