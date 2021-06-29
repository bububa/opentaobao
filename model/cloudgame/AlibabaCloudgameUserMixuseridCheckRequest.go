package cloudgame

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
云游戏混淆用户ID校验 APIRequest
alibaba.cloudgame.user.mixuserid.check

验证混淆用户ID是否合法
*/
type AlibabaCloudgameUserMixuseridCheckRequest struct {
    model.Params

    // 云游戏混淆用户ID
    mixUserId   string 

}

func NewAlibabaCloudgameUserMixuseridCheckRequest() *AlibabaCloudgameUserMixuseridCheckRequest{
    return &AlibabaCloudgameUserMixuseridCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCloudgameUserMixuseridCheckRequest) GetApiMethodName() string {
    return "alibaba.cloudgame.user.mixuserid.check"
}

func (r AlibabaCloudgameUserMixuseridCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCloudgameUserMixuseridCheckRequest) SetMixUserId(mixUserId string) error {
    r.mixUserId = mixUserId
    r.Set("mix_user_id", mixUserId)
    return nil
}

func (r AlibabaCloudgameUserMixuseridCheckRequest) GetMixUserId() string {
    return r.mixUserId
}

