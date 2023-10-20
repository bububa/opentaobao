package customizemarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallindustrybabyauthprofilebackflowAPIRequest 孕校云回流档案 API请求
// tmall.industry.baby.authprofile.backflow
//
// 孕校云回流档案
type TmallindustrybabyauthprofilebackflowAPIRequest struct {
	model.Params
	// 回流的档案数据
	_yxyBabyProfileCmd *YxyBabyProfileCmd
}

// NewTmallindustrybabyauthprofilebackflowRequest 初始化TmallindustrybabyauthprofilebackflowAPIRequest对象
func NewTmallindustrybabyauthprofilebackflowRequest() *TmallindustrybabyauthprofilebackflowAPIRequest {
	return &TmallindustrybabyauthprofilebackflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallindustrybabyauthprofilebackflowAPIRequest) GetApiMethodName() string {
	return "tmall.industry.baby.authprofile.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallindustrybabyauthprofilebackflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallindustrybabyauthprofilebackflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetYxyBabyProfileCmd is YxyBabyProfileCmd Setter
// 回流的档案数据
func (r *TmallindustrybabyauthprofilebackflowAPIRequest) SetYxyBabyProfileCmd(_yxyBabyProfileCmd *YxyBabyProfileCmd) error {
	r._yxyBabyProfileCmd = _yxyBabyProfileCmd
	r.Set("yxy_baby_profile_cmd", _yxyBabyProfileCmd)
	return nil
}

// GetYxyBabyProfileCmd YxyBabyProfileCmd Getter
func (r TmallindustrybabyauthprofilebackflowAPIRequest) GetYxyBabyProfileCmd() *YxyBabyProfileCmd {
	return r._yxyBabyProfileCmd
}
