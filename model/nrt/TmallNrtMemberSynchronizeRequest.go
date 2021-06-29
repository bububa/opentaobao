package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新零售会员同步接口 APIRequest
tmall.nrt.member.synchronize

新零售会员上翻接口，商家的会员信息同步至阿里侧
*/
type TmallNrtMemberSynchronizeRequest struct {
    model.Params

    // 会员DTO
    nrtMemberDto   *NrtMemberDTO 

}

func NewTmallNrtMemberSynchronizeRequest() *TmallNrtMemberSynchronizeRequest{
    return &TmallNrtMemberSynchronizeRequest{
        Params: model.NewParams(),
    }
}

func (r TmallNrtMemberSynchronizeRequest) GetApiMethodName() string {
    return "tmall.nrt.member.synchronize"
}

func (r TmallNrtMemberSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallNrtMemberSynchronizeRequest) SetNrtMemberDto(nrtMemberDto *NrtMemberDTO) error {
    r.nrtMemberDto = nrtMemberDto
    r.Set("nrt_member_dto", nrtMemberDto)
    return nil
}

func (r TmallNrtMemberSynchronizeRequest) GetNrtMemberDto() *NrtMemberDTO {
    return r.nrtMemberDto
}

