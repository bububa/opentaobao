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
    _memberId   int64
    // 创意审核入参
    _creative   *CreativeParamDTO
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
func (r *YunosAdAuditCreativeAddRequest) SetMemberId(_memberId int64) error {
    r._memberId = _memberId
    r.Set("member_id", _memberId)
    return nil
}

// MemberId Getter
func (r YunosAdAuditCreativeAddRequest) GetMemberId() int64 {
    return r._memberId
}
// Creative Setter
// 创意审核入参
func (r *YunosAdAuditCreativeAddRequest) SetCreative(_creative *CreativeParamDTO) error {
    r._creative = _creative
    r.Set("creative", _creative)
    return nil
}

// Creative Getter
func (r YunosAdAuditCreativeAddRequest) GetCreative() *CreativeParamDTO {
    return r._creative
}
