package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmMemberinfoUpdateAPIRequest 编辑会员资料 API请求
// taobao.crm.memberinfo.update
//
// 编辑会员的基本资料，接口返回会员信息修改是否成功
type TaobaoCrmMemberinfoUpdateAPIRequest struct {
	model.Params
}

// NewTaobaoCrmMemberinfoUpdateRequest 初始化TaobaoCrmMemberinfoUpdateAPIRequest对象
func NewTaobaoCrmMemberinfoUpdateRequest() *TaobaoCrmMemberinfoUpdateAPIRequest {
	return &TaobaoCrmMemberinfoUpdateAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmMemberinfoUpdateAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmMemberinfoUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.memberinfo.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmMemberinfoUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmMemberinfoUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoCrmMemberinfoUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmMemberinfoUpdateRequest()
	},
}

// GetTaobaoCrmMemberinfoUpdateRequest 从 sync.Pool 获取 TaobaoCrmMemberinfoUpdateAPIRequest
func GetTaobaoCrmMemberinfoUpdateAPIRequest() *TaobaoCrmMemberinfoUpdateAPIRequest {
	return poolTaobaoCrmMemberinfoUpdateAPIRequest.Get().(*TaobaoCrmMemberinfoUpdateAPIRequest)
}

// ReleaseTaobaoCrmMemberinfoUpdateAPIRequest 将 TaobaoCrmMemberinfoUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmMemberinfoUpdateAPIRequest(v *TaobaoCrmMemberinfoUpdateAPIRequest) {
	v.Reset()
	poolTaobaoCrmMemberinfoUpdateAPIRequest.Put(v)
}
