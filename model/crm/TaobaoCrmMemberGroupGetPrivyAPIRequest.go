package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberGroupGetPrivyAPIRequest 获取买家身上的标签（隐私号版） API请求
// taobao.crm.member.group.get.privy
//
// 获取买家身上的标签，不返回标签的总人数
type TaobaoCrmMemberGroupGetPrivyAPIRequest struct {
	model.Params
	// ouid
	_ouid string
}

// NewTaobaoCrmMemberGroupGetPrivyRequest 初始化TaobaoCrmMemberGroupGetPrivyAPIRequest对象
func NewTaobaoCrmMemberGroupGetPrivyRequest() *TaobaoCrmMemberGroupGetPrivyAPIRequest {
	return &TaobaoCrmMemberGroupGetPrivyAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmMemberGroupGetPrivyAPIRequest) Reset() {
	r._ouid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMemberGroupGetPrivyAPIRequest) GetApiMethodName() string {
	return "taobao.crm.member.group.get.privy"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMemberGroupGetPrivyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmMemberGroupGetPrivyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuid is Ouid Setter
// ouid
func (r *TaobaoCrmMemberGroupGetPrivyAPIRequest) SetOuid(_ouid string) error {
	r._ouid = _ouid
	r.Set("ouid", _ouid)
	return nil
}

// GetOuid Ouid Getter
func (r TaobaoCrmMemberGroupGetPrivyAPIRequest) GetOuid() string {
	return r._ouid
}

var poolTaobaoCrmMemberGroupGetPrivyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmMemberGroupGetPrivyRequest()
	},
}

// GetTaobaoCrmMemberGroupGetPrivyRequest 从 sync.Pool 获取 TaobaoCrmMemberGroupGetPrivyAPIRequest
func GetTaobaoCrmMemberGroupGetPrivyAPIRequest() *TaobaoCrmMemberGroupGetPrivyAPIRequest {
	return poolTaobaoCrmMemberGroupGetPrivyAPIRequest.Get().(*TaobaoCrmMemberGroupGetPrivyAPIRequest)
}

// ReleaseTaobaoCrmMemberGroupGetPrivyAPIRequest 将 TaobaoCrmMemberGroupGetPrivyAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmMemberGroupGetPrivyAPIRequest(v *TaobaoCrmMemberGroupGetPrivyAPIRequest) {
	v.Reset()
	poolTaobaoCrmMemberGroupGetPrivyAPIRequest.Put(v)
}
