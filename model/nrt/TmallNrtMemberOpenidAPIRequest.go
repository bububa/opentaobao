package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtmemberopenidAPIRequest 根据会员手机查询openId API请求
// tmall.nrt.member.openid
//
// 根据会员手机查询openId
type TmallnrtmemberopenidAPIRequest struct {
	model.Params
	// 会员DTO
	_nrtMemberDto *NrtMemberDto
}

// NewTmallnrtmemberopenidRequest 初始化TmallnrtmemberopenidAPIRequest对象
func NewTmallnrtmemberopenidRequest() *TmallnrtmemberopenidAPIRequest {
	return &TmallnrtmemberopenidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtmemberopenidAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.member.openid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtmemberopenidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtmemberopenidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtMemberDto is NrtMemberDto Setter
// 会员DTO
func (r *TmallnrtmemberopenidAPIRequest) SetNrtMemberDto(_nrtMemberDto *NrtMemberDto) error {
	r._nrtMemberDto = _nrtMemberDto
	r.Set("nrt_member_dto", _nrtMemberDto)
	return nil
}

// GetNrtMemberDto NrtMemberDto Getter
func (r TmallnrtmemberopenidAPIRequest) GetNrtMemberDto() *NrtMemberDto {
	return r._nrtMemberDto
}
