package mei

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallMeiCrmMemberGetbypaycodeAPIRequest 支付码获取会员信息 API请求
// tmall.mei.crm.member.getbypaycode
//
// 通过支付码获取会员信息
type TmallMeiCrmMemberGetbypaycodeAPIRequest struct {
	model.Params
	// 会员码
	_payCode string
}

// NewTmallMeiCrmMemberGetbypaycodeRequest 初始化TmallMeiCrmMemberGetbypaycodeAPIRequest对象
func NewTmallMeiCrmMemberGetbypaycodeRequest() *TmallMeiCrmMemberGetbypaycodeAPIRequest {
	return &TmallMeiCrmMemberGetbypaycodeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMeiCrmMemberGetbypaycodeAPIRequest) Reset() {
	r._payCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMeiCrmMemberGetbypaycodeAPIRequest) GetApiMethodName() string {
	return "tmall.mei.crm.member.getbypaycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMeiCrmMemberGetbypaycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMeiCrmMemberGetbypaycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayCode is PayCode Setter
// 会员码
func (r *TmallMeiCrmMemberGetbypaycodeAPIRequest) SetPayCode(_payCode string) error {
	r._payCode = _payCode
	r.Set("pay_code", _payCode)
	return nil
}

// GetPayCode PayCode Getter
func (r TmallMeiCrmMemberGetbypaycodeAPIRequest) GetPayCode() string {
	return r._payCode
}

var poolTmallMeiCrmMemberGetbypaycodeAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMeiCrmMemberGetbypaycodeRequest()
	},
}

// GetTmallMeiCrmMemberGetbypaycodeRequest 从 sync.Pool 获取 TmallMeiCrmMemberGetbypaycodeAPIRequest
func GetTmallMeiCrmMemberGetbypaycodeAPIRequest() *TmallMeiCrmMemberGetbypaycodeAPIRequest {
	return poolTmallMeiCrmMemberGetbypaycodeAPIRequest.Get().(*TmallMeiCrmMemberGetbypaycodeAPIRequest)
}

// ReleaseTmallMeiCrmMemberGetbypaycodeAPIRequest 将 TmallMeiCrmMemberGetbypaycodeAPIRequest 放入 sync.Pool
func ReleaseTmallMeiCrmMemberGetbypaycodeAPIRequest(v *TmallMeiCrmMemberGetbypaycodeAPIRequest) {
	v.Reset()
	poolTmallMeiCrmMemberGetbypaycodeAPIRequest.Put(v)
}
