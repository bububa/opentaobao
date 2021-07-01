package crm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoCrmGrademktMemberDetailDeleteAPIRequest
会员等级营销-删除商品等级营销明细 API请求
taobao.crm.grademkt.member.detail.delete

删除商品等级营销明细 */
type TaobaoCrmGrademktMemberDetailDeleteAPIRequest struct {
	model.Params
	// 扩展字段
	_feather string
	// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
	_parameter string
}

// NewTaobaoCrmGrademktMemberDetailDeleteRequest 初始化TaobaoCrmGrademktMemberDetailDeleteAPIRequest对象
func NewTaobaoCrmGrademktMemberDetailDeleteRequest() *TaobaoCrmGrademktMemberDetailDeleteAPIRequest {
	return &TaobaoCrmGrademktMemberDetailDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCrmGrademktMemberDetailDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.crm.grademkt.member.detail.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCrmGrademktMemberDetailDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Feather Setter
// 扩展字段
func (r *TaobaoCrmGrademktMemberDetailDeleteAPIRequest) SetFeather(_feather string) error {
	r._feather = _feather
	r.Set("feather", _feather)
	return nil
}

// Get Feather Getter
func (r TaobaoCrmGrademktMemberDetailDeleteAPIRequest) GetFeather() string {
	return r._feather
}

// Set is Parameter Setter
// 创建营销详情，生成方法见http://open.taobao.com/doc/detail.htm?id=101281
func (r *TaobaoCrmGrademktMemberDetailDeleteAPIRequest) SetParameter(_parameter string) error {
	r._parameter = _parameter
	r.Set("parameter", _parameter)
	return nil
}

// Get Parameter Getter
func (r TaobaoCrmGrademktMemberDetailDeleteAPIRequest) GetParameter() string {
	return r._parameter
}
