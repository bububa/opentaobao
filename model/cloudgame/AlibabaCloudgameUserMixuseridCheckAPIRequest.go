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

// NewAlibabaCloudgameUserMixuseridCheckRequest 初始化AlibabaCloudgameUserMixuseridCheckAPIRequest对象
func NewAlibabaCloudgameUserMixuseridCheckRequest() *AlibabaCloudgameUserMixuseridCheckAPIRequest {
	return &AlibabaCloudgameUserMixuseridCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameUserMixuseridCheckAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.user.mixuserid.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameUserMixuseridCheckAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MixUserId Setter
// 云游戏混淆用户ID
func (r *AlibabaCloudgameUserMixuseridCheckAPIRequest) SetMixUserId(_mixUserId string) error {
	r._mixUserId = _mixUserId
	r.Set("mix_user_id", _mixUserId)
	return nil
}

// Get MixUserId Getter
func (r AlibabaCloudgameUserMixuseridCheckAPIRequest) GetMixUserId() string {
	return r._mixUserId
}
