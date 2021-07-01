package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbImportThreeplOfflineConsignAPIRequest
3PL直邮线下发货 API请求
taobao.wlb.import.threepl.offline.consign

菜鸟认证直邮线下发货 */
type TaobaoWlbImportThreeplOfflineConsignAPIRequest struct {
	model.Params
	// 交易单号
	_tradeId int64
	// 资源id
	_resId int64
	// 资源code
	_resCode string
	// 运单号
	_waybillNo string
	// 发件人地址库id
	_fromId int64
}

// New
