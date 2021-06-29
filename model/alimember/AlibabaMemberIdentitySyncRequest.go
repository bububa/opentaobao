package alimember

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员身份信息同步 APIRequest
alibaba.member.identity.sync

会员身份信息同步
*/
type AlibabaMemberIdentitySyncRequest struct {
    model.Params

    // 会员身份同步信息
    syncDto   *SyncMemberIdentityDto 

}

func NewAlibabaMemberIdentitySyncRequest() *AlibabaMemberIdentitySyncRequest{
    return &AlibabaMemberIdentitySyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMemberIdentitySyncRequest) GetApiMethodName() string {
    return "alibaba.member.identity.sync"
}

func (r AlibabaMemberIdentitySyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMemberIdentitySyncRequest) SetSyncDto(syncDto *SyncMemberIdentityDto) error {
    r.syncDto = syncDto
    r.Set("sync_dto", syncDto)
    return nil
}

func (r AlibabaMemberIdentitySyncRequest) GetSyncDto() *SyncMemberIdentityDto {
    return r.syncDto
}

