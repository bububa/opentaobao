package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractIsvadminBindAPIRequest
创建及绑定互动实例 API请求
alibaba.interact.isvadmin.bind

创建互动实例，并绑定奖池 */
type AlibabaInteractIsvadminBindAPIRequest struct {
	model.Params
	// 描述信息
	_interactDescription string
	// 互动实例名称
	_instanceName string
	// 互动开始时间
	_startTime string
	// 互动结束时间
	_endTime string
	// 奖池ID
	_lotteryCode string
}

// New
