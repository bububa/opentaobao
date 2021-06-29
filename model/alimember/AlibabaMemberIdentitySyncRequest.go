package alimember

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员身份信息同步 API请求
alibaba.member.identity.sync

会员身份信息同步
*/
type AlibabaMemberIdentitySyncRequest struct {
    model.Params
    // 会员身份同步信息
    _syncDto   *SyncMemberIdentityDto
}

// 初始化AlibabaMemberIdentitySyncRequest对象
func NewAlibabaMemberIdentitySyncRequest() *AlibabaMemberIdentitySyncRequest{
    return &AlibabaMemberIdentitySyncRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMemberIdentitySyncRequest) GetApiMethodName() string {
    return "alibaba.member.identity.sync"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMemberIdentitySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SyncDto Setter
// 会员身份同步信息
func (r *AlibabaMemberIdentitySyncRequest) SetSyncDto(_syncDto *SyncMemberIdentityDto) error {
    r._syncDto = _syncDto
    r.Set("sync_dto", _syncDto)
    return nil
}

// SyncDto Getter
func (r AlibabaMemberIdentitySyncRequest) GetSyncDto() *SyncMemberIdentityDto {
    return r._syncDto
}
