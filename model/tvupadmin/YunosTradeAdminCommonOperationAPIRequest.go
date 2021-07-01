package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosTradeAdminCommonOperationAPIRequest
交易迎客松通用服务接口 API请求
yunos.trade.admin.common.operation

迎客松交易相关通用接口 */
type YunosTradeAdminCommonOperationAPIRequest struct {
	model.Params
	// 入参数据，json格式
	_parameter string
	// 调用方法
	_methodName string
	// 调用接口
	_interfaceName string
}

// New
