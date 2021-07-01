package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoReachableBatchjudgeAPIRequest
是否派送可达判定批量查询接口 API请求
cainiao.reachable.batchjudge

提供给商家在发货之前做截单处理，输入物流商编码和收发货地址进行可达判定，目前支持国内主流的物流服务商, 支持快运和快递两种类型 */
type CainiaoReachableBatchjudgeAPIRequest struct {
	model.Params
	// 调用方对象
	_clientInfo *ClientInfoDto
	// 1:快递 2:快运
	_addressType int64
	// 收发信息
	_data *RoutingReachableBatchRequestDto
}

// New
