package yunosad

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取创意审核状态 APIRequest
yunos.ad.audit.creative.getlist

批量获取创意审核状态
*/
type YunosAdAuditCreativeGetlistRequest struct {
    model.Params

    // 第三方DSP的id
    memberId   int64 

    // 状态
    status   string 

    // 创意列表
    creativeIds   []string 

}

func NewYunosAdAuditCreativeGetlistRequest() *YunosAdAuditCreativeGetlistRequest{
    return &YunosAdAuditCreativeGetlistRequest{
        Params: model.NewParams(),
    }
}

func (r YunosAdAuditCreativeGetlistRequest) GetApiMethodName() string {
    return "yunos.ad.audit.creative.getlist"
}

func (r YunosAdAuditCreativeGetlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosAdAuditCreativeGetlistRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r YunosAdAuditCreativeGetlistRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *YunosAdAuditCreativeGetlistRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

func (r YunosAdAuditCreativeGetlistRequest) GetStatus() string {
    return r.status
}

func (r *YunosAdAuditCreativeGetlistRequest) SetCreativeIds(creativeIds []string) error {
    r.creativeIds = creativeIds
    r.Set("creative_ids", creativeIds)
    return nil
}

func (r YunosAdAuditCreativeGetlistRequest) GetCreativeIds() []string {
    return r.creativeIds
}

