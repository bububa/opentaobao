package yunosad

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单个创意预审接口 APIRequest
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

func NewYunosAdAuditCreativeAddRequest() *YunosAdAuditCreativeAddRequest{
    return &YunosAdAuditCreativeAddRequest{
        Params: model.NewParams(),
    }
}

func (r YunosAdAuditCreativeAddRequest) GetApiMethodName() string {
    return "yunos.ad.audit.creative.add"
}

func (r YunosAdAuditCreativeAddRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosAdAuditCreativeAddRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r YunosAdAuditCreativeAddRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *YunosAdAuditCreativeAddRequest) SetCreative(creative *CreativeParamDto) error {
    r.creative = creative
    r.Set("creative", creative)
    return nil
}

func (r YunosAdAuditCreativeAddRequest) GetCreative() *CreativeParamDto {
    return r.creative
}

