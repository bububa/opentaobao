package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitembuygiftcreateactivityAPIRequest 创建买赠活动 API请求
// alibaba.wdk.marketing.itembuygift.createactivity
//
// 创建买赠活动
type AlibabawdkmarketingitembuygiftcreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemBuyGiftActivity
}

// NewAlibabawdkmarketingitembuygiftcreateactivityRequest 初始化AlibabawdkmarketingitembuygiftcreateactivityAPIRequest对象
func NewAlibabawdkmarketingitembuygiftcreateactivityRequest() *AlibabawdkmarketingitembuygiftcreateactivityAPIRequest {
	return &AlibabawdkmarketingitembuygiftcreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitembuygiftcreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itembuygift.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitembuygiftcreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitembuygiftcreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabawdkmarketingitembuygiftcreateactivityAPIRequest) SetParam(_param *ItemBuyGiftActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitembuygiftcreateactivityAPIRequest) GetParam() *ItemBuyGiftActivity {
	return r._param
}
