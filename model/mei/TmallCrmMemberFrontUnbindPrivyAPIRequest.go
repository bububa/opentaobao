package mei

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCrmMemberFrontUnbindPrivyAPIRequest) Reset() {
	r._ouid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCrmMemberFrontUnbindPrivyAPIRequest) GetApiMethodName() string {
	return "tmall.crm.member.front.unbind.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCrmMemberFrontUnbindPrivyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCrmMemberFrontUnbindPrivyAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallCrmMemberFrontUnbindPrivyAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCrmMemberFrontUnbindPrivyRequest()
	},
}

// GetTmallCrmMemberFrontUnbindPrivyRequest 从 sync.Pool 获取 TmallCrmMemberFrontUnbindPrivyAPIRequest
func GetTmallCrmMemberFrontUnbindPrivyAPIRequest() *TmallCrmMemberFrontUnbindPrivyAPIRequest {
	return poolTmallCrmMemberFrontUnbindPrivyAPIRequest.Get().(*TmallCrmMemberFrontUnbindPrivyAPIRequest)
}

// ReleaseTmallCrmMemberFrontUnbindPrivyAPIRequest 将 TmallCrmMemberFrontUnbindPrivyAPIRequest 放入 sync.Pool
func ReleaseTmallCrmMemberFrontUnbindPrivyAPIRequest(v *TmallCrmMemberFrontUnbindPrivyAPIRequest) {
	v.Reset()
	poolTmallCrmMemberFrontUnbindPrivyAPIRequest.Put(v)
}
