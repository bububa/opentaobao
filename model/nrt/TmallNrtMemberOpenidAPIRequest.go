package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtMemberOpenidAPIRequest 根据会员手机查询openId API请求
// tmall.nrt.member.openid
//
// 根据会员手机查询openId
type TmallNrtMemberOpenidAPIRequest struct {
	model.Params
	// 会员DTO
	_nrtMemberDto *NrtMemberDto
}

// NewTmallNrtMemberOpenidRequest 初始化TmallNrtMemberOpenidAPIRequest对象
func NewTmallNrtMemberOpenidRequest() *TmallNrtMemberOpenidAPIRequest {
	return &TmallNrtMemberOpenidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtMemberOpenidAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.member.openid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtMemberOpenidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNrtMemberDto is NrtMemberDto Setter
// 会员DTO
func (r *TmallNrtMemberOpenidAPIRequest) SetNrtMemberDto(_nrtMemberDto *NrtMemberDto) error {
	r._nrtMemberDto = _nrtMemberDto
	r.Set("nrt_member_dto", _nrtMemberDto)
	return nil
}

// GetNrtMemberDto NrtMemberDto Getter
func (r TmallNrtMemberOpenidAPIRequest) GetNrtMemberDto() *NrtMemberDto {
	return r._nrtMemberDto
}
