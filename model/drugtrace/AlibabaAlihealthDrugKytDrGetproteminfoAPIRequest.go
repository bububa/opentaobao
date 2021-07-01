package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrGetproteminfoAPIRequest
疫苗，获取生产企业的存储和运输温度 API请求
alibaba.alihealth.drug.kyt.dr.getproteminfo

疫苗，获取生产企业的存储和运输温度 */
type AlibabaAlihealthDrugKytDrGetproteminfoAPIRequest struct {
	model.Params
	// 调用企业ID
	_refEntId string
	// 药品ID
	_drugEntBaseInfoId string
	// 批次编号
	_batchNo string
	// 出库单号
	_billOutCode string
}

// New
