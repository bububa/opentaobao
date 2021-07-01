package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMeiCrmCallbackPointChangeAPIRequest
品牌积分变更回调API API请求
tmall.mei.crm.callback.point.change

线下品牌积分变更消息回调API，告诉积分扣减或者累加是否成功。 */
type TmallMeiCrmCallbackPointChangeAPIRequest struct {
	model.Params
	// 混淆会员手机号码
	_mixMobile string
	// 变更记录ID
	_recordId int64
	// 0:成功。1：失败
	_result int64
	// 处理失败的错误码.
	_errorCode string
	// 拓展信息
	_extInfo string
	// 积分信息
	_point int64
}

// New
