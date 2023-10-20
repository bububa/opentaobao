package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempoolactivitycreateAPIRequest 创建活动新接口 API请求
// alibaba.hm.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
type AlibabahmmarketingitempoolactivitycreateAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// NewAlibabahmmarketingitempoolactivitycreateRequest 初始化AlibabahmmarketingitempoolactivitycreateAPIRequest对象
func NewAlibabahmmarketingitempoolactivitycreateRequest() *AlibabahmmarketingitempoolactivitycreateAPIRequest {
	return &AlibabahmmarketingitempoolactivitycreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempoolactivitycreateAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempoolactivitycreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempoolactivitycreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabahmmarketingitempoolactivitycreateAPIRequest) SetParam(_param *ItemPoolActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitempoolactivitycreateAPIRequest) GetParam() *ItemPoolActivity {
	return r._param
}
