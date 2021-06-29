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
    _memberId   int64
    // 状态
    _status   string
    // 创意列表
    _creativeIds   []string
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
func (r *YunosAdAuditCreativeGetlistRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r YunosAdAuditCreativeGetlistRequest) GetMemberId() int64 {
    return r._memberId
}
// Status Setter
// 状态
func (r *YunosAdAuditCreativeGetlistRequest) SetStatus(_status string) error {
    r._status = _status
    r.Set("status", _status)
    return nil
}

// Status Getter
func (r YunosAdAuditCreativeGetlistRequest) GetStatus() string {
    return r._status
}
// CreativeIds Setter
// 创意列表
func (r *YunosAdAuditCreativeGetlistRequest) SetCreativeIds(_creativeIds []string) error {
    r._creativeIds = _creativeIds
    r.Set("creative_ids", _creativeIds)
    return nil
}

// CreativeIds Getter
func (r YunosAdAuditCreativeGetlistRequest) GetCreativeIds() []string {
    return r._creativeIds
}
