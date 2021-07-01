package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrDrugrecalAPIRequest
疫苗药品召回 API请求
alibaba.alihealth.drug.kyt.dr.drugrecal

生产企业发布的召回信息，按照批次进行召回，收货和发货环节的单据处理中调用接口进行查询； */
type AlibabaAlihealthDrugKytDrDrugrecalAPIRequest struct {
	model.Params
	// 调用企业ID
	_refEntId string
	// 召回开始时间
	_recallBeginTime string
	// 召回结束时间
	_recallEndTime string
}

// New
