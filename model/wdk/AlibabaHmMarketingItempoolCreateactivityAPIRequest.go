package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolcreateactivityAPIRequest 添加商品池活动 API请求
// alibaba.hm.marketing.itempool.createactivity
//
// 添加商品池活动
type AlibabahmmarketingitempoolcreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// NewAlibabahmmarketingitempoolcreateactivityRequest 初始化AlibabahmmarketingitempoolcreateactivityAPIRequest对象
func NewAlibabahmmarketingitempoolcreateactivityRequest() *AlibabahmmarketingitempoolcreateactivityAPIRequest {
	return &AlibabahmmarketingitempoolcreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempoolcreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempoolcreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempoolcreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabahmmarketingitempoolcreateactivityAPIRequest) SetParam(_param *ItemPoolActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitempoolcreateactivityAPIRequest) GetParam() *ItemPoolActivity {
	return r._param
}
