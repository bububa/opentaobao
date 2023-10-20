package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest 删除商品特价活动 API请求
// alibaba.wdk.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
type AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest struct {
	model.Params
	// 需要删除的活动的信息
	_param *CommonActivityRequest
}

// NewAlibabawdkmarketingitemdiscountdeleteactivityRequest 初始化AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest对象
func NewAlibabawdkmarketingitemdiscountdeleteactivityRequest() *AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest {
	return &AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itemdiscount.deleteactivity"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 需要删除的活动的信息
func (r *AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest) SetParam(_param *CommonActivityRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabawdkmarketingitemdiscountdeleteactivityAPIRequest) GetParam() *CommonActivityRequest {
	return r._param
}
