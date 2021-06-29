package cloudgame

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户Avatar body查询 APIRequest
alibaba.cgame.avatar.userbody.query

Avatar用户body数据查询
*/
type AlibabaCgameAvatarUserbodyQueryRequest struct {
    model.Params

    // 查询数据所属用户的mixUserId
    mixUserId   string 

}

func NewAlibabaCgameAvatarUserbodyQueryRequest() *AlibabaCgameAvatarUserbodyQueryRequest{
    return &AlibabaCgameAvatarUserbodyQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaCgameAvatarUserbodyQueryRequest) GetApiMethodName() string {
    return "alibaba.cgame.avatar.userbody.query"
}

func (r AlibabaCgameAvatarUserbodyQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaCgameAvatarUserbodyQueryRequest) SetMixUserId(mixUserId string) error {
    r.mixUserId = mixUserId
    r.Set("mix_user_id", mixUserId)
    return nil
}

func (r AlibabaCgameAvatarUserbodyQueryRequest) GetMixUserId() string {
    return r.mixUserId
}

