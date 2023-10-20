package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingVersionGenerateAPIRequest 生成发布使用的版本号 API请求
// alibaba.hm.marketing.version.generate
//
// 生成发布使用的版本号
type AlibabaHmMarketingVersionGenerateAPIRequest struct {
	model.Params
	// 档期版本号参数信息
	_param *SeasonVersionParam
}

// NewAlibabaHmMarketingVersionGenerateRequest 初始化AlibabaHmMarketingVersionGenerateAPIRequest对象
func NewAlibabaHmMarketingVersionGenerateRequest() *AlibabaHmMarketingVersionGenerateAPIRequest {
	return &AlibabaHmMarketingVersionGenerateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingVersionGenerateAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingVersionGenerateAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.version.generate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingVersionGenerateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingVersionGenerateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 档期版本号参数信息
func (r *AlibabaHmMarketingVersionGenerateAPIRequest) SetParam(_param *SeasonVersionParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingVersionGenerateAPIRequest) GetParam() *SeasonVersionParam {
	return r._param
}

var poolAlibabaHmMarketingVersionGenerateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingVersionGenerateRequest()
	},
}

// GetAlibabaHmMarketingVersionGenerateRequest 从 sync.Pool 获取 AlibabaHmMarketingVersionGenerateAPIRequest
func GetAlibabaHmMarketingVersionGenerateAPIRequest() *AlibabaHmMarketingVersionGenerateAPIRequest {
	return poolAlibabaHmMarketingVersionGenerateAPIRequest.Get().(*AlibabaHmMarketingVersionGenerateAPIRequest)
}

// ReleaseAlibabaHmMarketingVersionGenerateAPIRequest 将 AlibabaHmMarketingVersionGenerateAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingVersionGenerateAPIRequest(v *AlibabaHmMarketingVersionGenerateAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingVersionGenerateAPIRequest.Put(v)
}
