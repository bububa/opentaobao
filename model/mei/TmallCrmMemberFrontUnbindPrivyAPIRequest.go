package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcrmmemberfrontunbindprivyAPIRequest 品牌会员解绑(隐私号版） API请求
// tmall.crm.member.front.unbind.privy
//
// 品牌会员解绑功能
type TmallcrmmemberfrontunbindprivyAPIRequest struct {
	model.Params
	// ouid
	_ouid string
}

// NewTmallcrmmemberfrontunbindprivyRequest 初始化TmallcrmmemberfrontunbindprivyAPIRequest对象
func NewTmallcrmmemberfrontunbindprivyRequest() *TmallcrmmemberfrontunbindprivyAPIRequest {
	return &TmallcrmmemberfrontunbindprivyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcrmmemberfrontunbindprivyAPIRequest) GetApiMethodName() string {
	return "tmall.crm.member.front.unbind.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcrmmemberfrontunbindprivyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcrmmemberfrontunbindprivyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuid is Ouid Setter
// ouid
func (r *TmallcrmmemberfrontunbindprivyAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TmallcrmmemberfrontunbindprivyAPIRequest) GetOuid() string {
	return r._ouid
}
