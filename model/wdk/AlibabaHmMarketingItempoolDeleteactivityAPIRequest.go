package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempooldeleteactivityAPIRequest 删除商品池活动 API请求
// alibaba.hm.marketing.itempool.deleteactivity
//
// 删除商品池活动
type AlibabahmmarketingitempooldeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityParam
}

// NewAlibabahmmarketingitempooldeleteactivityRequest 初始化AlibabahmmarketingitempooldeleteactivityAPIRequest对象
func NewAlibabahmmarketingitempooldeleteactivityRequest() *AlibabahmmarketingitempooldeleteactivityAPIRequest {
	return &AlibabahmmarketingitempooldeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempooldeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempooldeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempooldeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabahmmarketingitempooldeleteactivityAPIRequest) SetParam(_param *CommonActivityParam) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitempooldeleteactivityAPIRequest) GetParam() *CommonActivityParam {
	return r._param
}
