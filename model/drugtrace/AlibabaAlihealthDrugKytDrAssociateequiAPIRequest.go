package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrAssociateequiAPIRequest
疫苗单据与设备绑定 API请求
alibaba.alihealth.drug.kyt.dr.associateequi

疫苗单据与设备绑定 */
type AlibabaAlihealthDrugKytDrAssociateequiAPIRequest struct {
	model.Params
	// 企业refentid
	_refEntId string
	// 单据编号，多个用逗号分隔
	_billCodes string
	// 设备ID
	_vaEquipmentId string
}

// New
