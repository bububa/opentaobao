package yunosad

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个创意预审接口 API请求
yunos.ad.audit.creative.add

YunOS广告业务第三方DSP单个创意预审接口
*/
type YunosAdAuditCreativeAddRequest struct {
    model.Params
    // 外部dsp的id
    memberId   int64
    // 创意审核入参
    creative   *CreativeParamDto
}

// 初始化YunosAdAuditCreativeAddRequest对象
func NewYunosAdAuditCreativeAddRequest() *YunosAdAuditCreativeAddRequest{
    return &YunosAdAuditCreativeAddRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosAdAuditCreativeAddRequest) GetApiMethodName() string {
    return "yunos.ad.audit.creative.add"
}

// IRequest interface 方法, 获取API参数
func (r YunosAdAuditCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MemberId Setter
// 外部dsp的id
func (r *YunosAdAuditCreativeAddRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

// MemberId Getter
func (r YunosAdAuditCreativeAddRequest) GetMemberId() int64 {
    return r.memberId
}
// Creative Setter
// 创意审核入参
func (r *YunosAdAuditCreativeAddRequest) SetCreative(creative *CreativeParamDto) error {
    r.creative = creative
    r.Set("creative", creative)
    return nil
}

// Creative Getter
func (r YunosAdAuditCreativeAddRequest) GetCreative() *CreativeParamDto {
    return r.creative
}
