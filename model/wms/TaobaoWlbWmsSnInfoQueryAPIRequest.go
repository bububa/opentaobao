package wms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWlbWmsSnInfoQueryAPIRequest
查询单据序列号信息 API请求
taobao.wlb.wms.sn.info.query

查询仓库作业的各类单据记录的Sn信息 */
type TaobaoWlbWmsSnInfoQueryAPIRequest struct {
	model.Params
	// 订单编码
	_orderCode string
	// 订单类型（1:仓配订单 10：配送扫码 20：增值扫码 40:ERP单号; 50：交易订单 ）
	_orderCodeType int64
	// 页数，默认每页50条
	_pageIndex int64
}

// New
