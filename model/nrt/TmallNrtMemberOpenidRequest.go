package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据会员手机查询openId API请求
tmall.nrt.member.openid

根据会员手机查询openId
*/
type TmallNrtMemberOpenidRequest struct {
    model.Params
    // 会员DTO
    _nrtMemberDto   *NrtMemberDTO
}

// 初始化TmallNrtMemberOpenidRequest对象
func NewTmallNrtMemberOpenidRequest() *TmallNrtMemberOpenidRequest{
    return &TmallNrtMemberOpenidRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtMemberOpenidRequest) GetApiMethodName() string {
    return "tmall.nrt.member.openid"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtMemberOpenidRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtMemberDto Setter
// 会员DTO
func (r *TmallNrtMemberOpenidRequest) SetNrtMemberDto(_nrtMemberDto *NrtMemberDTO) error {
    r._nrtMemberDto = _nrtMemberDto
    r.Set("nrt_member_dto", _nrtMemberDto)
    return nil
}

// NrtMemberDto Getter
func (r TmallNrtMemberOpenidRequest) GetNrtMemberDto() *NrtMemberDTO {
    return r._nrtMemberDto
}
