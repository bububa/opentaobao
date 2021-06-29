package alimember

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
会员信息同步 APIRequest
alibaba.member.sync

会员信息同步
*/
type AlibabaMemberSyncRequest struct {
    model.Params

    // 会员同步信息
    syncMember   *SyncMemberDto 

}

func NewAlibabaMemberSyncRequest() *AlibabaMemberSyncRequest{
    return &AlibabaMemberSyncRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMemberSyncRequest) GetApiMethodName() string {
    return "alibaba.member.sync"
}

func (r AlibabaMemberSyncRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMemberSyncRequest) SetSyncMember(syncMember *SyncMemberDto) error {
    r.syncMember = syncMember
    r.Set("sync_member", syncMember)
    return nil
}

func (r AlibabaMemberSyncRequest) GetSyncMember() *SyncMemberDto {
    return r.syncMember
}

