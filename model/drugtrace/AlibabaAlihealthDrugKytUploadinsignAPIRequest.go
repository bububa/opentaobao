package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytUploadinsignAPIRequest
仓库批量扫码回传接口 API请求
alibaba.alihealth.drug.kyt.uploadinsign

连锁总部仓库在采购入库或者销售出库环节，批量采集追溯码之后回传到码上放心平台。 */
type AlibabaAlihealthDrugKytUploadinsignAPIRequest struct {
	model.Params
	// 单据编号（小于20位字符串，唯一）
	_billCode string
	// 单据生成时间
	_billTime string
	// 码上放心平台企业编码（仓库所有者）
	_refUserId string
	// 仓库名称（企业自定义）
	_warehouseId string
	// 药品ID
	_drugId string
	// 追溯码[多个时用逗号分开]
	_traceCodes []string
}

// New
