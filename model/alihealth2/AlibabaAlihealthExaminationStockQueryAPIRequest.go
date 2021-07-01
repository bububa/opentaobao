package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthExaminationStockQueryAPIRequest
体检机构对接_体检套餐库存查询 API请求
alibaba.alihealth.examination.stock.query

体检机构对接_体检套餐库存查询 */
type AlibabaAlihealthExaminationStockQueryAPIRequest struct {
	model.Params
	// 商户唯一码
	_merchantCode string
	// 分院ID
	_hospitalId string
	// 套餐id
	_packageId string
	// 开始日期
	_timeFrom string
	// 结束日期
	_timeTo string
}

// New
