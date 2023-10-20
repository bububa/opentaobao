package nrt

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallNrtMemberOpenidAPIRequest) Reset() {
	r._nrtMemberDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtMemberOpenidAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.member.openid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtMemberOpenidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallNrtMemberOpenidAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallNrtMemberOpenidAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallNrtMemberOpenidRequest()
	},
}

// GetTmallNrtMemberOpenidRequest 从 sync.Pool 获取 TmallNrtMemberOpenidAPIRequest
func GetTmallNrtMemberOpenidAPIRequest() *TmallNrtMemberOpenidAPIRequest {
	return poolTmallNrtMemberOpenidAPIRequest.Get().(*TmallNrtMemberOpenidAPIRequest)
}

// ReleaseTmallNrtMemberOpenidAPIRequest 将 TmallNrtMemberOpenidAPIRequest 放入 sync.Pool
func ReleaseTmallNrtMemberOpenidAPIRequest(v *TmallNrtMemberOpenidAPIRequest) {
	v.Reset()
	poolTmallNrtMemberOpenidAPIRequest.Put(v)
}
