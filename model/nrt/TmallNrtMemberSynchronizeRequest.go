package nrt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新零售会员同步接口 API请求
tmall.nrt.member.synchronize

新零售会员上翻接口，商家的会员信息同步至阿里侧
*/
type TmallNrtMemberSynchronizeRequest struct {
    model.Params
    // 会员DTO
    _nrtMemberDto   *NrtMemberDTO
}

// 初始化TmallNrtMemberSynchronizeRequest对象
func NewTmallNrtMemberSynchronizeRequest() *TmallNrtMemberSynchronizeRequest{
    return &TmallNrtMemberSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallNrtMemberSynchronizeRequest) GetApiMethodName() string {
    return "tmall.nrt.member.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TmallNrtMemberSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// NrtMemberDto Setter
// 会员DTO
func (r *TmallNrtMemberSynchronizeRequest) SetNrtMemberDto(_nrtMemberDto *NrtMemberDTO) error {
    r._nrtMemberDto = _nrtMemberDto
    r.Set("nrt_member_dto", _nrtMemberDto)
    return nil
}

// NrtMemberDto Getter
func (r TmallNrtMemberSynchronizeRequest) GetNrtMemberDto() *NrtMemberDTO {
    return r._nrtMemberDto
}
