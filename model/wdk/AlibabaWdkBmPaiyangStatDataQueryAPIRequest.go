package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkbmpaiyangstatdataqueryAPIRequest 派样统计数据查询 API请求
// alibaba.wdk.bm.paiyang.stat.data.query
//
// 派样统计数据查询
type AlibabawdkbmpaiyangstatdataqueryAPIRequest struct {
	model.Params
	// 入参对象
	_param *PaiyangStatDataParam
}

// NewAlibabawdkbmpaiyangstatdataqueryRequest 初始化AlibabawdkbmpaiyangstatdataqueryAPIRequest对象
func NewAlibabawdkbmpaiyangstatdataqueryRequest() *AlibabawdkbmpaiyangstatdataqueryAPIRequest {
	return &AlibabawdkbmpaiyangstatdataqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkbmpaiyangstatdataqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.paiyang.stat.data.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkbmpaiyangstatdataqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkbmpaiyangstatdataqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlibabawdkbmpaiyangstatdataqueryAPIRequest) SetParam(_param *PaiyangStatDataParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkbmpaiyangstatdataqueryAPIRequest) GetParam() *PaiyangStatDataParam {
	return r._param
}
