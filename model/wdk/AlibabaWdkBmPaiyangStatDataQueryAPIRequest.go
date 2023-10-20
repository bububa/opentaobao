package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmPaiyangStatDataQueryAPIRequest 派样统计数据查询 API请求
// alibaba.wdk.bm.paiyang.stat.data.query
//
// 派样统计数据查询
type AlibabaWdkBmPaiyangStatDataQueryAPIRequest struct {
	model.Params
	// 入参对象
	_param *PaiyangStatDataParam
}

// NewAlibabaWdkBmPaiyangStatDataQueryRequest 初始化AlibabaWdkBmPaiyangStatDataQueryAPIRequest对象
func NewAlibabaWdkBmPaiyangStatDataQueryRequest() *AlibabaWdkBmPaiyangStatDataQueryAPIRequest {
	return &AlibabaWdkBmPaiyangStatDataQueryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkBmPaiyangStatDataQueryAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkBmPaiyangStatDataQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.bm.paiyang.stat.data.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkBmPaiyangStatDataQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkBmPaiyangStatDataQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlibabaWdkBmPaiyangStatDataQueryAPIRequest) SetParam(_param *PaiyangStatDataParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaWdkBmPaiyangStatDataQueryAPIRequest) GetParam() *PaiyangStatDataParam {
	return r._param
}

var poolAlibabaWdkBmPaiyangStatDataQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkBmPaiyangStatDataQueryRequest()
	},
}

// GetAlibabaWdkBmPaiyangStatDataQueryRequest 从 sync.Pool 获取 AlibabaWdkBmPaiyangStatDataQueryAPIRequest
func GetAlibabaWdkBmPaiyangStatDataQueryAPIRequest() *AlibabaWdkBmPaiyangStatDataQueryAPIRequest {
	return poolAlibabaWdkBmPaiyangStatDataQueryAPIRequest.Get().(*AlibabaWdkBmPaiyangStatDataQueryAPIRequest)
}

// ReleaseAlibabaWdkBmPaiyangStatDataQueryAPIRequest 将 AlibabaWdkBmPaiyangStatDataQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkBmPaiyangStatDataQueryAPIRequest(v *AlibabaWdkBmPaiyangStatDataQueryAPIRequest) {
	v.Reset()
	poolAlibabaWdkBmPaiyangStatDataQueryAPIRequest.Put(v)
}
