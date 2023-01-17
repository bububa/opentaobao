package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityUpdateAPIRequest ISV活动修改 API请求
// alibaba.lsy.crm.activity.update
//
// ISV活动修改
type AlibabaLsyCrmActivityUpdateAPIRequest struct {
	model.Params
	// 入参
	_nrtUpdateActivityReq *NrtUpdateActivityReq
}

// NewAlibabaLsyCrmActivityUpdateRequest 初始化AlibabaLsyCrmActivityUpdateAPIRequest对象
func NewAlibabaLsyCrmActivityUpdateRequest() *AlibabaLsyCrmActivityUpdateAPIRequest {
	return &AlibabaLsyCrmActivityUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyCrmActivityUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtUpdateActivityReq is NrtUpdateActivityReq Setter
// 入参
func (r *AlibabaLsyCrmActivityUpdateAPIRequest) SetNrtUpdateActivityReq(_nrtUpdateActivityReq *NrtUpdateActivityReq) error {
	r._nrtUpdateActivityReq = _nrtUpdateActivityReq
	r.Set("nrt_update_activity_req", _nrtUpdateActivityReq)
	return nil
}

// GetNrtUpdateActivityReq NrtUpdateActivityReq Getter
func (r AlibabaLsyCrmActivityUpdateAPIRequest) GetNrtUpdateActivityReq() *NrtUpdateActivityReq {
	return r._nrtUpdateActivityReq
}
