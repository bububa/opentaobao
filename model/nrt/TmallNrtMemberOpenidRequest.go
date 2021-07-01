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
type TmallNrtMemberOpenidAPIRequest struct {
    model.Params
    // 会员DTO
    _nrtMemberDto   *NrtMemberDTO
}

// 初始化TmallNrtMemberOpenidAPIRequest对象
func NewTmallNrtMemberOpenidRequest() *TmallNrtMemberOpenidAPIRequest{
    return &TmallNrtMemberOpenidAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtMemberOpenidAPIRequest) GetApiMethodName() string {
    return "tmall.nrt.member.openid"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtMemberOpenidAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtMemberDto Setter
// 会员DTO
func (r *TmallNrtMemberOpenidAPIRequest) SetNrtMemberDto(_nrtMemberDto *NrtMemberDTO) error {
    r._nrtMemberDto = _nrtMemberDto
    r.Set("nrt_member_dto", _nrtMemberDto)
    return nil
}

// NrtMemberDto Getter
func (r TmallNrtMemberOpenidAPIRequest) GetNrtMemberDto() *NrtMemberDTO {
    return r._nrtMemberDto
}
