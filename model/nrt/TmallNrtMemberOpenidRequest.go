package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据会员手机查询openId APIRequest
tmall.nrt.member.openid

根据会员手机查询openId
*/
type TmallNrtMemberOpenidRequest struct {
    model.Params

    // 会员DTO
    nrtMemberDto   *NrtMemberDTO 

}

func NewTmallNrtMemberOpenidRequest() *TmallNrtMemberOpenidRequest{
    return &TmallNrtMemberOpenidRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtMemberOpenidRequest) GetApiMethodName() string {
    return "tmall.nrt.member.openid"
}

func (r TmallNrtMemberOpenidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtMemberOpenidRequest) SetNrtMemberDto(nrtMemberDto *NrtMemberDTO) error {
    r.nrtMemberDto = nrtMemberDto
    r.Set("nrt_member_dto", nrtMemberDto)
    return nil
}

func (r TmallNrtMemberOpenidRequest) GetNrtMemberDto() *NrtMemberDTO {
    return r.nrtMemberDto
}

