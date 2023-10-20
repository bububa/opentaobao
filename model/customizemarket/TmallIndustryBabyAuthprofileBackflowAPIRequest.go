package customizemarket

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallIndustryBabyAuthprofileBackflowAPIRequest 孕校云回流档案 API请求
// tmall.industry.baby.authprofile.backflow
//
// 孕校云回流档案
type TmallIndustryBabyAuthprofileBackflowAPIRequest struct {
	model.Params
	// 回流的档案数据
	_yxyBabyProfileCmd *YxyBabyProfileCmd
}

// NewTmallIndustryBabyAuthprofileBackflowRequest 初始化TmallIndustryBabyAuthprofileBackflowAPIRequest对象
func NewTmallIndustryBabyAuthprofileBackflowRequest() *TmallIndustryBabyAuthprofileBackflowAPIRequest {
	return &TmallIndustryBabyAuthprofileBackflowAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallIndustryBabyAuthprofileBackflowAPIRequest) Reset() {
	r._yxyBabyProfileCmd = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallIndustryBabyAuthprofileBackflowAPIRequest) GetApiMethodName() string {
	return "tmall.industry.baby.authprofile.backflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallIndustryBabyAuthprofileBackflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallIndustryBabyAuthprofileBackflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetYxyBabyProfileCmd is YxyBabyProfileCmd Setter
// 回流的档案数据
func (r *TmallIndustryBabyAuthprofileBackflowAPIRequest) SetYxyBabyProfileCmd(_yxyBabyProfileCmd *YxyBabyProfileCmd) error {
	r._yxyBabyProfileCmd = _yxyBabyProfileCmd
	r.Set("yxy_baby_profile_cmd", _yxyBabyProfileCmd)
	return nil
}

// GetYxyBabyProfileCmd YxyBabyProfileCmd Getter
func (r TmallIndustryBabyAuthprofileBackflowAPIRequest) GetYxyBabyProfileCmd() *YxyBabyProfileCmd {
	return r._yxyBabyProfileCmd
}

var poolTmallIndustryBabyAuthprofileBackflowAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallIndustryBabyAuthprofileBackflowRequest()
	},
}

// GetTmallIndustryBabyAuthprofileBackflowRequest 从 sync.Pool 获取 TmallIndustryBabyAuthprofileBackflowAPIRequest
func GetTmallIndustryBabyAuthprofileBackflowAPIRequest() *TmallIndustryBabyAuthprofileBackflowAPIRequest {
	return poolTmallIndustryBabyAuthprofileBackflowAPIRequest.Get().(*TmallIndustryBabyAuthprofileBackflowAPIRequest)
}

// ReleaseTmallIndustryBabyAuthprofileBackflowAPIRequest 将 TmallIndustryBabyAuthprofileBackflowAPIRequest 放入 sync.Pool
func ReleaseTmallIndustryBabyAuthprofileBackflowAPIRequest(v *TmallIndustryBabyAuthprofileBackflowAPIRequest) {
	v.Reset()
	poolTmallIndustryBabyAuthprofileBackflowAPIRequest.Put(v)
}
