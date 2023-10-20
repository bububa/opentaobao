package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolActivityCreateAPIRequest 创建活动新接口 API请求
// alibaba.hm.marketing.itempool.activity.create
//
// 创建活动新接口，支持新工具玩法
type AlibabaHmMarketingItempoolActivityCreateAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemPoolActivity
}

// NewAlibabaHmMarketingItempoolActivityCreateRequest 初始化AlibabaHmMarketingItempoolActivityCreateAPIRequest对象
func NewAlibabaHmMarketingItempoolActivityCreateRequest() *AlibabaHmMarketingItempoolActivityCreateAPIRequest {
	return &AlibabaHmMarketingItempoolActivityCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolActivityCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.activity.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolActivityCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolActivityCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabaHmMarketingItempoolActivityCreateAPIRequest) SetParam(_param *ItemPoolActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaHmMarketingItempoolActivityCreateAPIRequest) GetParam() *ItemPoolActivity {
	return r._param
}
