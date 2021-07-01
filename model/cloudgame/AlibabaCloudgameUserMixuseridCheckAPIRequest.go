package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCloudgameUserMixuseridCheckAPIRequest
云游戏混淆用户ID校验 API请求
alibaba.cloudgame.user.mixuserid.check

验证混淆用户ID是否合法 */
type AlibabaCloudgameUserMixuseridCheckAPIRequest struct {
	model.Params
	// 云游戏混淆用户ID
	_mixUserId string
}

// New
