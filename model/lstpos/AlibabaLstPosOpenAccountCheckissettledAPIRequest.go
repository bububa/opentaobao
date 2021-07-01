package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaLstPosOpenAccountCheckissettledAPIRequest
校验当前用户是否入驻了零售通门店接口 API请求
alibaba.lst.pos.open.account.checkissettled

校验当前用户是否入驻了零售通门店接口 */
type AlibabaLstPosOpenAccountCheckissettledAPIRequest struct {
	model.Params
	// 当前登录主账号userId
	_userId int64
}

// New
