package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtMemberSynchronizeAPIRequest 新零售会员同步接口 API请求
// tmall.nrt.member.synchronize
//
// 新零售会员上翻接口，商家的会员信息同步至阿里侧
type TmallNrtMemberSynchronizeAPIRequest struct {
	model.Params
	// 会员DTO
	_nrtMemberDto *NrtMemberDto
}

// NewTmallNrtMemberSynchronizeRequest 初始化TmallNrtMemberSynchronizeAPIRequest对象
func NewTmallNrtMemberSynchronizeRequest() *TmallNrtMemberSynchronizeAPIRequest {
	return &TmallNrtMemberSynchronizeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtMemberSynchronizeAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.member.synchronize"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtMemberSynchronizeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtMemberSynchronizeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtMemberDto is NrtMemberDto Setter
// 会员DTO
func (r *TmallNrtMemberSynchronizeAPIRequest) SetNrtMemberDto(_nrtMemberDto *NrtMemberDto) error {
	r._nrtMemberDto = _nrtMemberDto
	r.Set("nrt_member_dto", _nrtMemberDto)
	return nil
}

// GetNrtMemberDto NrtMemberDto Getter
func (r TmallNrtMemberSynchronizeAPIRequest) GetNrtMemberDto() *NrtMemberDto {
	return r._nrtMemberDto
}
