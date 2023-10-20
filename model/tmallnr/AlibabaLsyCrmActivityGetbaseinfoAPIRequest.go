package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityGetbaseinfoAPIRequest ISV查询活动 API请求
// alibaba.lsy.crm.activity.getbaseinfo
//
// ISV查询活动
type AlibabaLsyCrmActivityGetbaseinfoAPIRequest struct {
	model.Params
	// 入参
	_nrtQueryActivityReq *NrtQueryActivityReq
}

// NewAlibabaLsyCrmActivityGetbaseinfoRequest 初始化AlibabaLsyCrmActivityGetbaseinfoAPIRequest对象
func NewAlibabaLsyCrmActivityGetbaseinfoRequest() *AlibabaLsyCrmActivityGetbaseinfoAPIRequest {
	return &AlibabaLsyCrmActivityGetbaseinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityGetbaseinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.getbaseinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityGetbaseinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLsyCrmActivityGetbaseinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtQueryActivityReq is NrtQueryActivityReq Setter
// 入参
func (r *AlibabaLsyCrmActivityGetbaseinfoAPIRequest) SetNrtQueryActivityReq(_nrtQueryActivityReq *NrtQueryActivityReq) error {
	r._nrtQueryActivityReq = _nrtQueryActivityReq
	r.Set("nrt_query_activity_req", _nrtQueryActivityReq)
	return nil
}

// GetNrtQueryActivityReq NrtQueryActivityReq Getter
func (r AlibabaLsyCrmActivityGetbaseinfoAPIRequest) GetNrtQueryActivityReq() *NrtQueryActivityReq {
	return r._nrtQueryActivityReq
}
