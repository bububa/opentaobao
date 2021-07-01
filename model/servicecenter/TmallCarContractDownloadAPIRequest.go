package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCarContractDownloadAPIRequest
合同下载 API请求
tmall.car.contract.download

目前天猫开新车会在线上签署一份合同，协议，需要一个个在已卖出打开，另存为pdf，人工一个个下载比较麻烦，期望通过接口直接读取pdf；
因为比较耗时，建议一个个下载，假设并发下载，很可能限流，每天的调用量有限； */
type TmallCarContractDownloadAPIRequest struct {
	model.Params
	// 天猫订单号
	_orderId int64
	// 是否下载html，true是html，false是pdf， html速度会快一点
	_html bool
}

// New
