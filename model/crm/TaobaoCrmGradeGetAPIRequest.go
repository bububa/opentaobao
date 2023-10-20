package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGradeGetAPIRequest 卖家查询等级规则 API请求
// taobao.crm.grade.get
//
// 卖家查询等级规则，包括店铺客户、普通会员、高级会员、VIP会员、至尊VIP会员四个等级的信息
type TaobaoCrmGradeGetAPIRequest struct {
	model.Params
}

// NewTaobaoCrmGradeGetRequest 初始化TaobaoCrmGradeGetAPIRequest对象
func NewTaobaoCrmGradeGetRequest() *TaobaoCrmGradeGetAPIRequest {
	return &TaobaoCrmGradeGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmGradeGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGradeGetAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGradeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmGradeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolTaobaoCrmGradeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmGradeGetRequest()
	},
}

// GetTaobaoCrmGradeGetRequest 从 sync.Pool 获取 TaobaoCrmGradeGetAPIRequest
func GetTaobaoCrmGradeGetAPIRequest() *TaobaoCrmGradeGetAPIRequest {
	return poolTaobaoCrmGradeGetAPIRequest.Get().(*TaobaoCrmGradeGetAPIRequest)
}

// ReleaseTaobaoCrmGradeGetAPIRequest 将 TaobaoCrmGradeGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmGradeGetAPIRequest(v *TaobaoCrmGradeGetAPIRequest) {
	v.Reset()
	poolTaobaoCrmGradeGetAPIRequest.Put(v)
}
