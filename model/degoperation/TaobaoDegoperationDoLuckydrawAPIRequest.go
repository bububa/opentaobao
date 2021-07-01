package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegoperationDoLuckydrawAPIRequest
激励抽奖 API请求
taobao.degoperation.do.luckydraw

激励平台抽奖接口。用户可以通过接口完成抽奖功能 */
type TaobaoDegoperationDoLuckydrawAPIRequest struct {
	model.Params
	// 后台活动配置appkey
	_degAppKey string
	// 后台活动配置eventkey
	_degEventKey string
	// 前端标识
	_source string
	// 设备uuid
	_uuid string
	// 参数校验
	_paramSign string
	// 传参信息
	_degAccessToken string
}

// New
