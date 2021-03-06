package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallNrtEasyhomememberSynAPIRequest 会员信息同 API请求
// tmall.nrt.easyhomemember.syn
//
// 居然之家将会员信息同步到零售中台 包含基本的会员信息
type TmallNrtEasyhomememberSynAPIRequest struct {
	model.Params
	// 入参
	_param *ExternalMemberDto
}

// NewTmallNrtEasyhomememberSynRequest 初始化TmallNrtEasyhomememberSynAPIRequest对象
func NewTmallNrtEasyhomememberSynRequest() *TmallNrtEasyhomememberSynAPIRequest {
	return &TmallNrtEasyhomememberSynAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallNrtEasyhomememberSynAPIRequest) GetApiMethodName() string {
	return "tmall.nrt.easyhomemember.syn"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallNrtEasyhomememberSynAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParam is Param Setter
// 入参
func (r *TmallNrtEasyhomememberSynAPIRequest) SetParam(_param *ExternalMemberDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TmallNrtEasyhomememberSynAPIRequest) GetParam() *ExternalMemberDto {
	return r._param
}
