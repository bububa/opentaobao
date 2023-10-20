package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivitypageupdateAPIRequest ISV活动页面创建修改 API请求
// alibaba.lsy.crm.activity.page.update
//
// ISV活动页面创建修改
type AlibabalsycrmactivitypageupdateAPIRequest struct {
	model.Params
	// 入参
	_nrtCrmActivityPageCreateReq *NrtCrmActivityPageCreateReq
}

// NewAlibabalsycrmactivitypageupdateRequest 初始化AlibabalsycrmactivitypageupdateAPIRequest对象
func NewAlibabalsycrmactivitypageupdateRequest() *AlibabalsycrmactivitypageupdateAPIRequest {
	return &AlibabalsycrmactivitypageupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmactivitypageupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.page.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmactivitypageupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmactivitypageupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtCrmActivityPageCreateReq is NrtCrmActivityPageCreateReq Setter
// 入参
func (r *AlibabalsycrmactivitypageupdateAPIRequest) SetNrtCrmActivityPageCreateReq(_nrtCrmActivityPageCreateReq *NrtCrmActivityPageCreateReq) error {
	r._nrtCrmActivityPageCreateReq = _nrtCrmActivityPageCreateReq
	r.Set("nrt_crm_activity_page_create_req", _nrtCrmActivityPageCreateReq)
	return nil
}

// GetNrtCrmActivityPageCreateReq NrtCrmActivityPageCreateReq Getter
func (r AlibabalsycrmactivitypageupdateAPIRequest) GetNrtCrmActivityPageCreateReq() *NrtCrmActivityPageCreateReq {
	return r._nrtCrmActivityPageCreateReq
}
