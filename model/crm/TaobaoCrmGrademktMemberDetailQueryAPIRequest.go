package crm

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGrademktMemberDetailQueryAPIRequest 会员等级营销-等级营销活动查询 API请求
// taobao.crm.grademkt.member.detail.query
//
// 商家通过该接口查询等级营销活动
type TaobaoCrmGrademktMemberDetailQueryAPIRequest struct {
	model.Params
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
	// 扩展字段
	_feather string
}

// NewTaobaoCrmGrademktMemberDetailQueryRequest 初始化TaobaoCrmGrademktMemberDetailQueryAPIRequest对象
func NewTaobaoCrmGrademktMemberDetailQueryRequest() *TaobaoCrmGrademktMemberDetailQueryAPIRequest {
	return &TaobaoCrmGrademktMemberDetailQueryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoCrmGrademktMemberDetailQueryAPIRequest) Reset() {
	r._parameter = ""
	r._feather = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberDetailQueryAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grademkt.member.detail.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberDetailQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCrmGrademktMemberDetailQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParameter is Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberDetailQueryAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r TaobaoCrmGrademktMemberDetailQueryAPIRequest) GetParameter() string {
	return r._parameter
}

// SetFeather is Feather Setter
// 扩展字段
func (r *TaobaoCrmGrademktMemberDetailQueryAPIRequest) SetFeather(_feather string) error {
	r._feather = _feather
	r.Set("feather", _feather)
	return nil
}

// GetFeather Feather Getter
func (r TaobaoCrmGrademktMemberDetailQueryAPIRequest) GetFeather() string {
	return r._feather
}

var poolTaobaoCrmGrademktMemberDetailQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoCrmGrademktMemberDetailQueryRequest()
	},
}

// GetTaobaoCrmGrademktMemberDetailQueryRequest 从 sync.Pool 获取 TaobaoCrmGrademktMemberDetailQueryAPIRequest
func GetTaobaoCrmGrademktMemberDetailQueryAPIRequest() *TaobaoCrmGrademktMemberDetailQueryAPIRequest {
	return poolTaobaoCrmGrademktMemberDetailQueryAPIRequest.Get().(*TaobaoCrmGrademktMemberDetailQueryAPIRequest)
}

// ReleaseTaobaoCrmGrademktMemberDetailQueryAPIRequest 将 TaobaoCrmGrademktMemberDetailQueryAPIRequest 放入 sync.Pool
func ReleaseTaobaoCrmGrademktMemberDetailQueryAPIRequest(v *TaobaoCrmGrademktMemberDetailQueryAPIRequest) {
	v.Reset()
	poolTaobaoCrmGrademktMemberDetailQueryAPIRequest.Put(v)
}
