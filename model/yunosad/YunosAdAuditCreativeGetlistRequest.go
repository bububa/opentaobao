package yunosad

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
批量获取创意审核状态 API请求
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

// 初始化YunosAdAuditCreativeGetlistRequest对象
func NewYunosAdAuditCreativeGetlistRequest() *YunosAdAuditCreativeGetlistRequest{
    return &YunosAdAuditCreativeGetlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosAdAuditCreativeGetlistRequest) GetApiMethodName() string {
    return "yunos.ad.audit.creative.getlist"
}

// IRequest interface 方法, 获取API参数
func (r YunosAdAuditCreativeGetlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// 第三方DSP的id
func (r *YunosAdAuditCreativeGetlistRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r YunosAdAuditCreativeGetlistRequest) GetMemberId() int64 {
    return r.memberId
}
// Status Setter
// 状态
func (r *YunosAdAuditCreativeGetlistRequest) SetStatus(status string) error {
    r.status = status
    r.Set("status", status)
    return nil
}

// Status Getter
func (r YunosAdAuditCreativeGetlistRequest) GetStatus() string {
    return r.status
}
// CreativeIds Setter
// 创意列表
func (r *YunosAdAuditCreativeGetlistRequest) SetCreativeIds(creativeIds []string) error {
    r.creativeIds = creativeIds
    r.Set("creative_ids", creativeIds)
    return nil
}

// CreativeIds Getter
func (r YunosAdAuditCreativeGetlistRequest) GetCreativeIds() []string {
    return r.creativeIds
}
