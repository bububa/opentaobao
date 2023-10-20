package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitembuygiftcreateactivityAPIRequest 创建买赠活动 API请求
// alibaba.hm.marketing.itembuygift.createactivity
//
// 创建买赠活动
type AlibabahmmarketingitembuygiftcreateactivityAPIRequest struct {
	model.Params
	// 创建活动请求入参
	_param *ItemBuyGiftActivity
}

// NewAlibabahmmarketingitembuygiftcreateactivityRequest 初始化AlibabahmmarketingitembuygiftcreateactivityAPIRequest对象
func NewAlibabahmmarketingitembuygiftcreateactivityRequest() *AlibabahmmarketingitembuygiftcreateactivityAPIRequest {
	return &AlibabahmmarketingitembuygiftcreateactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitembuygiftcreateactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itembuygift.createactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitembuygiftcreateactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitembuygiftcreateactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 创建活动请求入参
func (r *AlibabahmmarketingitembuygiftcreateactivityAPIRequest) SetParam(_param *ItemBuyGiftActivity) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabahmmarketingitembuygiftcreateactivityAPIRequest) GetParam() *ItemBuyGiftActivity {
	return r._param
}
