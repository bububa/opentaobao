package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingdiscountitemaddasyncAPIRequest 特价批量新增商品 API请求
// alibaba.hm.marketing.discount.item.add.async
//
// 新分组模型下新增商品
type AlibabahmmarketingdiscountitemaddasyncAPIRequest struct {
	model.Params
	// sku信息
	_param0 []ItemDiscountSku
	// 系统自动生成
	_param1 *CommonActivityParam
	// alibaba.wdk.marketing.version.generate接口生成
	_version int64
}

// NewAlibabahmmarketingdiscountitemaddasyncRequest 初始化AlibabahmmarketingdiscountitemaddasyncAPIRequest对象
func NewAlibabahmmarketingdiscountitemaddasyncRequest() *AlibabahmmarketingdiscountitemaddasyncAPIRequest {
	return &AlibabahmmarketingdiscountitemaddasyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingdiscountitemaddasyncAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.discount.item.add.async"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingdiscountitemaddasyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingdiscountitemaddasyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// sku信息
func (r *AlibabahmmarketingdiscountitemaddasyncAPIRequest) SetParam0(_param0 []ItemDiscountSku) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabahmmarketingdiscountitemaddasyncAPIRequest) GetParam0() []ItemDiscountSku {
	return r._param0
}

// SetParam1 is Param1 Setter
// 系统自动生成
func (r *AlibabahmmarketingdiscountitemaddasyncAPIRequest) SetParam1(_param1 *CommonActivityParam) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r AlibabahmmarketingdiscountitemaddasyncAPIRequest) GetParam1() *CommonActivityParam {
	return r._param1
}

// SetVersion is Version Setter
// alibaba.wdk.marketing.version.generate接口生成
func (r *AlibabahmmarketingdiscountitemaddasyncAPIRequest) SetVersion(_version int64) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r AlibabahmmarketingdiscountitemaddasyncAPIRequest) GetVersion() int64 {
	return r._version
}
