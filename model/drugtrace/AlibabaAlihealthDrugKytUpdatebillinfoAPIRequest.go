package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest
零售端平台单据更新 API请求
alibaba.alihealth.drug.kyt.updatebillinfo

零售端平台单据更新 */
type AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 企业ID
	_entId string
	// 操作人编码
	_icCode string
	// 单据ID
	_billId string
	// 单据类型
	_billType string
	// 单据编码
	_billCode string
	// 发货单位ID
	_partnerIdSend string
	// 收货单信ID
	_partnerIdRecv string
	// 详情
	_note string
}

// New
