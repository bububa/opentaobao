package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractIsvadminGetpondbyinteractAPIRequest
根据互动实例ID查询奖池信息 API请求
alibaba.interact.isvadmin.getpondbyinteract

根据互动实例ID查询奖池信息 */
type AlibabaInteractIsvadminGetpondbyinteractAPIRequest struct {
	model.Params
	// 互动实例ID
	_interactId string
}

// New
