package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivitygetbaseinfoAPIRequest ISV查询活动 API请求
// alibaba.lsy.crm.activity.getbaseinfo
//
// ISV查询活动
type AlibabalsycrmactivitygetbaseinfoAPIRequest struct {
	model.Params
	// 入参
	_nrtQueryActivityReq *NrtQueryActivityReq
}

// NewAlibabalsycrmactivitygetbaseinfoRequest 初始化AlibabalsycrmactivitygetbaseinfoAPIRequest对象
func NewAlibabalsycrmactivitygetbaseinfoRequest() *AlibabalsycrmactivitygetbaseinfoAPIRequest {
	return &AlibabalsycrmactivitygetbaseinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmactivitygetbaseinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.getbaseinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmactivitygetbaseinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmactivitygetbaseinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNrtQueryActivityReq is NrtQueryActivityReq Setter
// 入参
func (r *AlibabalsycrmactivitygetbaseinfoAPIRequest) SetNrtQueryActivityReq(_nrtQueryActivityReq *NrtQueryActivityReq) error {
	r._nrtQueryActivityReq = _nrtQueryActivityReq
	r.Set("nrt_query_activity_req", _nrtQueryActivityReq)
	return nil
}

// GetNrtQueryActivityReq NrtQueryActivityReq Getter
func (r AlibabalsycrmactivitygetbaseinfoAPIRequest) GetNrtQueryActivityReq() *NrtQueryActivityReq {
	return r._nrtQueryActivityReq
}
