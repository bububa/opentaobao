package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstAstrolabeOrderstatusSyncAPIRequest
线下门店派单以及单据相关操作接口 API请求
taobao.jst.astrolabe.orderstatus.sync

针对ERP系统部署在门店的商家，将派单状态回流到星盘 */
type TaobaoJstAstrolabeOrderstatusSyncAPIRequest struct {
	model.Params
	// 子订单Id
	_subOrderIds []int64
	// 事件发生时间
	_actionTime string
	// 操作人
	_operator string
	// 业务类型
	_type string
	// 订单状态
	_status string
	// 目标门店的商户中心门店编码
	_storeId int64
	// 交易订单
	_parentOrderCode int64
}

// New
