package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMsfIdentifyStatusQueryAPIRequest
喵师傅定案核销状态查询 API请求
tmall.msf.identify.status.query

喵师傅定案核销状态查询，供服务商erp系统调用 */
type TmallMsfIdentifyStatusQueryAPIRequest struct {
	model.Params
	// 天猫订单号
	_orderId int64
	// 服务类型，0 家装的送货上门并安装 1 单向安装 2 建材的送货上门 3 建材的安装
	_serviceType int64
}

// New
