package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCrmGrademktMemberDetailCreateAPIRequest 会员等级营销-创建商品等级营销明细 API请求
// taobao.crm.grademkt.member.detail.create
//
// 创建商品等级营销明细
type TaobaoCrmGrademktMemberDetailCreateAPIRequest struct {
	model.Params
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
	// 扩展字段
	_feather string
}

// NewTaobaoCrmGrademktMemberDetailCreateRequest 初始化TaobaoCrmGrademktMemberDetailCreateAPIRequest对象
func NewTaobaoCrmGrademktMemberDetailCreateRequest() *TaobaoCrmGrademktMemberDetailCreateAPIRequest {
	return &TaobaoCrmGrademktMemberDetailCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberDetailCreateAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grademkt.member.detail.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberDetailCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParameter is Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberDetailCreateAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// GetParameter Parameter Getter
func (r TaobaoCrmGrademktMemberDetailCreateAPIRequest) GetParameter() string {
	return r._parameter
}

// SetFeather is Feather Setter
// 扩展字段
func (r *TaobaoCrmGrademktMemberDetailCreateAPIRequest) SetFeather(_feather string) error {
	r._feather = _feather
	r.Set("feather", _feather)
	return nil
}

// GetFeather Feather Getter
func (r TaobaoCrmGrademktMemberDetailCreateAPIRequest) GetFeather() string {
	return r._feather
}
