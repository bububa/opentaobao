package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUploadb2cbillAPIRequest
快易通零售B2C API请求
alibaba.alihealth.drug.kyt.uploadb2cbill

快易通零售B2C单据上传 */
type AlibabaAlihealthDrugKytUploadb2cbillAPIRequest struct {
	model.Params
	// 单据号【20位以内的唯一编码，可以使用16位UUID】
	_billCode string
	// 单据时间【一般情况下写当前时间】
	_billTime string
	// 企业ID
	_refUserId string
	// 操作人
	_operIcCode string
	// 主订单
	_masterOrder string
	// lbx号
	_lbxOrder string
	// 仓号
	_warehouseId string
	// 药品ID
	_drugId string
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
	// 订单来源
	_orderSource string
}

// New
