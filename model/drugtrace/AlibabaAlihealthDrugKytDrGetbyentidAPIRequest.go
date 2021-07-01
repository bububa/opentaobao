package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrGetbyentidAPIRequest
多融通过企业ID得到一个企业的详细信息 API请求
alibaba.alihealth.drug.kyt.dr.getbyentid

根据企业主键查看企业详细信息 */
type AlibabaAlihealthDrugKytDrGetbyentidAPIRequest struct {
	model.Params
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 准备要查询的企业ID（返回该企业ID的详细信息）
	_entId string
}

// New
