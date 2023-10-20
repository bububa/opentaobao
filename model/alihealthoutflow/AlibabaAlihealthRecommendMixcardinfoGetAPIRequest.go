package alihealthoutflow

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthRecommendMixcardinfoGetAPIRequest 快应用混合卡片信息 API请求
// alibaba.alihealth.recommend.mixcardinfo.get
//
// 快应用混合卡片信息
type AlibabaAlihealthRecommendMixcardinfoGetAPIRequest struct {
	model.Params
	// 请求入参
	_quickAppRequest *QuickAppRequest
}

// NewAlibabaAlihealthRecommendMixcardinfoGetRequest 初始化AlibabaAlihealthRecommendMixcardinfoGetAPIRequest对象
func NewAlibabaAlihealthRecommendMixcardinfoGetRequest() *AlibabaAlihealthRecommendMixcardinfoGetAPIRequest {
	return &AlibabaAlihealthRecommendMixcardinfoGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) Reset() {
	r._quickAppRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.recommend.mixcardinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuickAppRequest is QuickAppRequest Setter
// 请求入参
func (r *AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) SetQuickAppRequest(_quickAppRequest *QuickAppRequest) error {
	r._quickAppRequest = _quickAppRequest
	r.Set("quick_app_request", _quickAppRequest)
	return nil
}

// GetQuickAppRequest QuickAppRequest Getter
func (r AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) GetQuickAppRequest() *QuickAppRequest {
	return r._quickAppRequest
}

var poolAlibabaAlihealthRecommendMixcardinfoGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthRecommendMixcardinfoGetRequest()
	},
}

// GetAlibabaAlihealthRecommendMixcardinfoGetRequest 从 sync.Pool 获取 AlibabaAlihealthRecommendMixcardinfoGetAPIRequest
func GetAlibabaAlihealthRecommendMixcardinfoGetAPIRequest() *AlibabaAlihealthRecommendMixcardinfoGetAPIRequest {
	return poolAlibabaAlihealthRecommendMixcardinfoGetAPIRequest.Get().(*AlibabaAlihealthRecommendMixcardinfoGetAPIRequest)
}

// ReleaseAlibabaAlihealthRecommendMixcardinfoGetAPIRequest 将 AlibabaAlihealthRecommendMixcardinfoGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthRecommendMixcardinfoGetAPIRequest(v *AlibabaAlihealthRecommendMixcardinfoGetAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthRecommendMixcardinfoGetAPIRequest.Put(v)
}
