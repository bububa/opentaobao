package cloudgame

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabacloudgameusermixuseridcheckAPIRequest 云游戏混淆用户ID校验 API请求
// alibaba.cloudgame.user.mixuserid.check
//
// 验证混淆用户ID是否合法
type AlibabacloudgameusermixuseridcheckAPIRequest struct {
	model.Params
	// 云游戏混淆用户ID
	_mixUserId string
}

// NewAlibabacloudgameusermixuseridcheckRequest 初始化AlibabacloudgameusermixuseridcheckAPIRequest对象
func NewAlibabacloudgameusermixuseridcheckRequest() *AlibabacloudgameusermixuseridcheckAPIRequest {
	return &AlibabacloudgameusermixuseridcheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabacloudgameusermixuseridcheckAPIRequest) GetApiMethodName() string {
	return "alibaba.cloudgame.user.mixuserid.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabacloudgameusermixuseridcheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabacloudgameusermixuseridcheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetMixUserId is MixUserId Setter
// 云游戏混淆用户ID
func (r *AlibabacloudgameusermixuseridcheckAPIRequest) SetMixUserId(_mixUserId string) error {
	r._mixUserId = _mixUserId
	r.Set("mix_user_id", _mixUserId)
	return nil
}

// GetMixUserId MixUserId Getter
func (r AlibabacloudgameusermixuseridcheckAPIRequest) GetMixUserId() string {
	return r._mixUserId
}
