package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest
零头出入库单据上传 API请求
alibaba.alihealth.drug.kyt.remnantbill.upload

零头出入库单据上传 */
type AlibabaAlihealthDrugKytRemnantbillUploadAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 零头入库：106；零头出库：210
	_billType string
	// 单据编号
	_billCode string
	// 单据时间:yyyy-MM-dd HH:mm:ss
	_billTime string
	// 发货企业ID
	_fromRefUserId string
	// 收货企业ID
	_toRefUserId string
	// 委托企业ID
	_assRefEntId string
	// 配送企业ID
	_disRefEntId string
	// 药品ID
	_drugEntBaseInfoId string
	// 生产日期：yyyy-MM-dd
	_produceDate string
	// 有效期:yyyyMMdd
	_expireDate string
	// 生产批次
	_produceBatchNo string
	// 药品数量
	_inputAmount string
}

// New
