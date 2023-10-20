package tmallnr

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityPageUpdateAPIRequest ISV活动页面创建修改 API请求
// alibaba.lsy.crm.activity.page.update
//
// ISV活动页面创建修改
type AlibabaLsyCrmActivityPageUpdateAPIRequest struct {
	model.Params
	// 入参
	_nrtCrmActivityPageCreateReq *NrtCrmActivityPageCreateReq
}

// NewAlibabaLsyCrmActivityPageUpdateRequest 初始化AlibabaLsyCrmActivityPageUpdateAPIRequest对象
func NewAlibabaLsyCrmActivityPageUpdateRequest() *AlibabaLsyCrmActivityPageUpdateAPIRequest {
	return &AlibabaLsyCrmActivityPageUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLsyCrmActivityPageUpdateAPIRequest) Reset() {
	r._nrtCrmActivityPageCreateReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityPageUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.page.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityPageUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyCrmActivityPageUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtCrmActivityPageCreateReq is NrtCrmActivityPageCreateReq Setter
// 入参
func (r *AlibabaLsyCrmActivityPageUpdateAPIRequest) SetNrtCrmActivityPageCreateReq(_nrtCrmActivityPageCreateReq *NrtCrmActivityPageCreateReq) error {
	r._nrtCrmActivityPageCreateReq = _nrtCrmActivityPageCreateReq
	r.Set("nrt_crm_activity_page_create_req", _nrtCrmActivityPageCreateReq)
	return nil
}

// GetNrtCrmActivityPageCreateReq NrtCrmActivityPageCreateReq Getter
func (r AlibabaLsyCrmActivityPageUpdateAPIRequest) GetNrtCrmActivityPageCreateReq() *NrtCrmActivityPageCreateReq {
	return r._nrtCrmActivityPageCreateReq
}

var poolAlibabaLsyCrmActivityPageUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLsyCrmActivityPageUpdateRequest()
	},
}

// GetAlibabaLsyCrmActivityPageUpdateRequest 从 sync.Pool 获取 AlibabaLsyCrmActivityPageUpdateAPIRequest
func GetAlibabaLsyCrmActivityPageUpdateAPIRequest() *AlibabaLsyCrmActivityPageUpdateAPIRequest {
	return poolAlibabaLsyCrmActivityPageUpdateAPIRequest.Get().(*AlibabaLsyCrmActivityPageUpdateAPIRequest)
}

// ReleaseAlibabaLsyCrmActivityPageUpdateAPIRequest 将 AlibabaLsyCrmActivityPageUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaLsyCrmActivityPageUpdateAPIRequest(v *AlibabaLsyCrmActivityPageUpdateAPIRequest) {
	v.Reset()
	poolAlibabaLsyCrmActivityPageUpdateAPIRequest.Put(v)
}
