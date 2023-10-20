package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallnrtmembersynchronizeAPIRequest 新零售会员同步接口 API请求
// tmall.nrt.member.synchronize
//
// 新零售会员上翻接口，商家的会员信息同步至阿里侧
type TmallnrtmembersynchronizeAPIRequest struct {
	model.Params
	// 会员DTO
	_nrtMemberDto *NrtMemberDto
}

// NewTmallnrtmembersynchronizeRequest 初始化TmallnrtmembersynchronizeAPIRequest对象
func NewTmallnrtmembersynchronizeRequest() *TmallnrtmembersynchronizeAPIRequest {
	return &TmallnrtmembersynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallnrtmembersynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.member.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallnrtmembersynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallnrtmembersynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtMemberDto is NrtMemberDto Setter
// 会员DTO
func (r *TmallnrtmembersynchronizeAPIRequest) SetNrtMemberDto(_nrtMemberDto *NrtMemberDto) error {
	r._nrtMemberDto = _nrtMemberDto
	r.Set("nrt_member_dto", _nrtMemberDto)
	return nil
}

// GetNrtMemberDto NrtMemberDto Getter
func (r TmallnrtmembersynchronizeAPIRequest) GetNrtMemberDto() *NrtMemberDto {
	return r._nrtMemberDto
}
