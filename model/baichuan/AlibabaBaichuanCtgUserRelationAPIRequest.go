package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanCtgUserRelationAPIRequest
用户 API请求
alibaba.baichuan.ctg.user.relation

提供给优酷查询道长和淘宝账户的绑定关系 */
type AlibabaBaichuanCtgUserRelationAPIRequest struct {
	model.Params
	// 调用的业务方
	_app string
	// 业务方的用户ID
	_uid string
	// 淘宝的用户ID
	_tbUid string
}

// New
