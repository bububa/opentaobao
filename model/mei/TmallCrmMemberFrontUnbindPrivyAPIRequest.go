package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCrmMemberFrontUnbindPrivyAPIRequest 品牌会员解绑(隐私号版） API请求
// tmall.crm.member.front.unbind.privy
//
// 品牌会员解绑功能
type TmallCrmMemberFrontUnbindPrivyAPIRequest struct {
	model.Params
	// ouid
	_ouid string
}

// NewTmallCrmMemberFrontUnbindPrivyRequest 初始化TmallCrmMemberFrontUnbindPrivyAPIRequest对象
func NewTmallCrmMemberFrontUnbindPrivyRequest() *TmallCrmMemberFrontUnbindPrivyAPIRequest {
	return &TmallCrmMemberFrontUnbindPrivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCrmMemberFrontUnbindPrivyAPIRequest) GetApiMethodName() string {
	return "tmall.crm.member.front.unbind.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCrmMemberFrontUnbindPrivyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetOuid is Ouid Setter
// ouid
func (r *TmallCrmMemberFrontUnbindPrivyAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TmallCrmMemberFrontUnbindPrivyAPIRequest) GetOuid() string {
	return r._ouid
}
