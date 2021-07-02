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
func (r AlibabaLsyCrmActivityUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NrtUpdateActivityReq Setter
// 入参
func (r *AlibabaLsyCrmActivityUpdateAPIRequest) SetNrtUpdateActivityReq(_nrtUpdateActivityReq *NrtUpdateActivityReq) error {
	r._nrtUpdateActivityReq = _nrtUpdateActivityReq
	r.Set("nrt_update_activity_req", _nrtUpdateActivityReq)
	return nil
}

// Get NrtUpdateActivityReq Getter
func (r AlibabaLsyCrmActivityUpdateAPIRequest) GetNrtUpdateActivityReq() *NrtUpdateActivityReq {
	return r._nrtUpdateActivityReq
}
