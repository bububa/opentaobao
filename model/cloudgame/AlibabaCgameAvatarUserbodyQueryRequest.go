package cloudgame

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户Avatar body查询 API请求
alibaba.cgame.avatar.userbody.query

Avatar用户body数据查询
*/
type AlibabaCgameAvatarUserbodyQueryRequest struct {
    model.Params
    // 查询数据所属用户的mixUserId
    _mixUserId   string
}

// 初始化AlibabaCgameAvatarUserbodyQueryRequest对象
func NewAlibabaCgameAvatarUserbodyQueryRequest() *AlibabaCgameAvatarUserbodyQueryRequest{
    return &AlibabaCgameAvatarUserbodyQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaCgameAvatarUserbodyQueryRequest) GetApiMethodName() string {
    return "alibaba.cgame.avatar.userbody.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaCgameAvatarUserbodyQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MixUserId Setter
// 查询数据所属用户的mixUserId
func (r *AlibabaCgameAvatarUserbodyQueryRequest) SetMixUserId(_mixUserId string) error {
    r._mixUserId = _mixUserId
    r.Set("mix_user_id", _mixUserId)
    return nil
}

// MixUserId Getter
func (r AlibabaCgameAvatarUserbodyQueryRequest) GetMixUserId() string {
    return r._mixUserId
}
