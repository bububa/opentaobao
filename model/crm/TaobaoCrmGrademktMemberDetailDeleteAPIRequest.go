package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGrademktMemberDetailDeleteAPIRequest 会员等级营销-删除商品等级营销明细 API请求
// taobao.crm.grademkt.member.detail.delete
//
// 删除商品等级营销明细
type TaobaoCrmGrademktMemberDetailDeleteAPIRequest struct {
	model.Params
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
	// 扩展字段
	_feather string
}

// NewTaobaoCrmGrademktMemberDetailDeleteRequest 初始化TaobaoCrmGrademktMemberDetailDeleteAPIRequest对象
func NewTaobaoCrmGrademktMemberDetailDeleteRequest() *TaobaoCrmGrademktMemberDetailDeleteAPIRequest {
	return &TaobaoCrmGrademktMemberDetailDeleteAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmGrademktMemberDetailDeleteAPIRequest) Reset() {
	r._parameter = ""
	r._feather = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberDetailDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grademkt.member.detail.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberDetailDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmGrademktMemberDetailDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameter is Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberDetailDeleteAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r TaobaoCrmGrademktMemberDetailDeleteAPIRequest) GetParameter() string {
	return r._parameter
}

// SetFeather is Feather Setter
// 扩展字段
func (r *TaobaoCrmGrademktMemberDetailDeleteAPIRequest) SetFeather(_feather string) error {
	r._feather = _feather
	r.Set("feather", _feather)
	return nil
}

// GetFeather Feather Getter
func (r TaobaoCrmGrademktMemberDetailDeleteAPIRequest) GetFeather() string {
	return r._feather
}

var poolTaobaoCrmGrademktMemberDetailDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmGrademktMemberDetailDeleteRequest()
	},
}

// GetTaobaoCrmGrademktMemberDetailDeleteRequest 从 sync.Pool 获取 TaobaoCrmGrademktMemberDetailDeleteAPIRequest
func GetTaobaoCrmGrademktMemberDetailDeleteAPIRequest() *TaobaoCrmGrademktMemberDetailDeleteAPIRequest {
	return poolTaobaoCrmGrademktMemberDetailDeleteAPIRequest.Get().(*TaobaoCrmGrademktMemberDetailDeleteAPIRequest)
}

// ReleaseTaobaoCrmGrademktMemberDetailDeleteAPIRequest 将 TaobaoCrmGrademktMemberDetailDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmGrademktMemberDetailDeleteAPIRequest(v *TaobaoCrmGrademktMemberDetailDeleteAPIRequest) {
	v.Reset()
	poolTaobaoCrmGrademktMemberDetailDeleteAPIRequest.Put(v)
}
