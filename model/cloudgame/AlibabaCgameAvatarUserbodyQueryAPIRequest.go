package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCgameAvatarUserbodyQueryAPIRequest
用户Avatar body查询 API请求
alibaba.cgame.avatar.userbody.query

Avatar用户body数据查询 */
type AlibabaCgameAvatarUserbodyQueryAPIRequest struct {
	model.Params
	// 查询数据所属用户的mixUserId
	_mixUserId string
}

// New
