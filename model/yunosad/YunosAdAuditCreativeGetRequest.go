package yunosad

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取单个创意审核状态 APIRequest
yunos.ad.audit.creative.get

获取单个创意审核状态
*/
type YunosAdAuditCreativeGetRequest struct {
    model.Params

    // 第三方的dspId
    memberId   int64 

    // 第三方广告创意id
    creativeId   string 

}

func NewYunosAdAuditCreativeGetRequest() *YunosAdAuditCreativeGetRequest{
    return &YunosAdAuditCreativeGetRequest{
        Params: model.NewParams(),
    }
}

func (r YunosAdAuditCreativeGetRequest) GetApiMethodName() string {
    return "yunos.ad.audit.creative.get"
}

func (r YunosAdAuditCreativeGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosAdAuditCreativeGetRequest) SetMemberId(memberId int64) error {
    r.memberId = memberId
    r.Set("member_id", memberId)
    return nil
}

func (r YunosAdAuditCreativeGetRequest) GetMemberId() int64 {
    return r.memberId
}

func (r *YunosAdAuditCreativeGetRequest) SetCreativeId(creativeId string) error {
    r.creativeId = creativeId
    r.Set("creative_id", creativeId)
    return nil
}

func (r YunosAdAuditCreativeGetRequest) GetCreativeId() string {
    return r.creativeId
}

