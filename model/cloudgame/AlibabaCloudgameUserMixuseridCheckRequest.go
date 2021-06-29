package cloudgame

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云游戏混淆用户ID校验 API请求
alibaba.cloudgame.user.mixuserid.check

验证混淆用户ID是否合法
*/
type AlibabaCloudgameUserMixuseridCheckRequest struct {
    model.Params
    // 云游戏混淆用户ID
    mixUserId   string
}

// 初始化AlibabaCloudgameUserMixuseridCheckRequest对象
func NewAlibabaCloudgameUserMixuseridCheckRequest() *AlibabaCloudgameUserMixuseridCheckRequest{
    return &AlibabaCloudgameUserMixuseridCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCloudgameUserMixuseridCheckRequest) GetApiMethodName() string {
    return "alibaba.cloudgame.user.mixuserid.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCloudgameUserMixuseridCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MixUserId Setter
// 云游戏混淆用户ID
func (r *AlibabaCloudgameUserMixuseridCheckRequest) SetMixUserId(mixUserId string) error {
    r.mixUserId = mixUserId
    r.Set("mix_user_id", mixUserId)
    return nil
}

// MixUserId Getter
func (r AlibabaCloudgameUserMixuseridCheckRequest) GetMixUserId() string {
    return r.mixUserId
}
