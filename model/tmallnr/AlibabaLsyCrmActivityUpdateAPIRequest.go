package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivityupdateAPIRequest ISV活动修改 API请求
// alibaba.lsy.crm.activity.update
//
// ISV活动修改
type AlibabalsycrmactivityupdateAPIRequest struct {
	model.Params
	// 入参
	_nrtUpdateActivityReq *NrtUpdateActivityReq
}

// NewAlibabalsycrmactivityupdateRequest 初始化AlibabalsycrmactivityupdateAPIRequest对象
func NewAlibabalsycrmactivityupdateRequest() *AlibabalsycrmactivityupdateAPIRequest {
	return &AlibabalsycrmactivityupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmactivityupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmactivityupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmactivityupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtUpdateActivityReq is NrtUpdateActivityReq Setter
// 入参
func (r *AlibabalsycrmactivityupdateAPIRequest) SetNrtUpdateActivityReq(_nrtUpdateActivityReq *NrtUpdateActivityReq) error {
	r._nrtUpdateActivityReq = _nrtUpdateActivityReq
	r.Set("nrt_update_activity_req", _nrtUpdateActivityReq)
	return nil
}

// GetNrtUpdateActivityReq NrtUpdateActivityReq Getter
func (r AlibabalsycrmactivityupdateAPIRequest) GetNrtUpdateActivityReq() *NrtUpdateActivityReq {
	return r._nrtUpdateActivityReq
}
