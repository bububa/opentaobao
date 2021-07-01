package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytGetbyrefentidAPIRequest
根据企业唯一标识查看企业详细信息 API请求
alibaba.alihealth.drug.kyt.getbyrefentid

根据企业唯一标识查看企业详细信息 */
type AlibabaAlihealthDrugKytGetbyrefentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业唯一标识（返回该唯一标识企业的详细信息）
	_destRefEntId string
}

// New
