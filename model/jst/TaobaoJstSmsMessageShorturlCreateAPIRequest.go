package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstSmsMessageShorturlCreateAPIRequest
聚石塔营销效果短链生成 API请求
taobao.jst.sms.message.shorturl.create

聚石塔生成淘短链接口 */
type TaobaoJstSmsMessageShorturlCreateAPIRequest struct {
	model.Params
	// 是否需要https前缀： true-要  false-不要
	_needHttpsPrefix bool
	// 人群标签
	_tag string
	// 商品或者店铺的H5地址，只支持长链
	_url string
	// 批次号
	_batchNumber string
}

// NewTaobaoJstSmsMessageShorturlCreateRequest 初始化TaobaoJstSmsMessageShorturlCreateAPIRequest对象
func NewTaobaoJstSmsMessageShorturlCreateRequest() *TaobaoJstSmsMessageShorturlCreateAPIRequest {
	return &TaobaoJstSmsMessageShorturlCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.message.shorturl.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is NeedHttpsPrefix Setter
// 是否需要https前缀： true-要  false-不要
func (r *TaobaoJstSmsMessageShorturlCreateAPIRequest) SetNeedHttpsPrefix(_needHttpsPrefix bool) error {
	r._needHttpsPrefix = _needHttpsPrefix
	r.Set("need_https_prefix", _needHttpsPrefix)
	return nil
}

// Get NeedHttpsPrefix Getter
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetNeedHttpsPrefix() bool {
	return r._needHttpsPrefix
}

// Set is Tag Setter
// 人群标签
func (r *TaobaoJstSmsMessageShorturlCreateAPIRequest) SetTag(_tag string) error {
	r._tag = _tag
	r.Set("tag", _tag)
	return nil
}

// Get Tag Getter
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetTag() string {
	return r._tag
}

// Set is Url Setter
// 商品或者店铺的H5地址，只支持长链
func (r *TaobaoJstSmsMessageShorturlCreateAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// Get Url Getter
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetUrl() string {
	return r._url
}

// Set is BatchNumber Setter
// 批次号
func (r *TaobaoJstSmsMessageShorturlCreateAPIRequest) SetBatchNumber(_batchNumber string) error {
	r._batchNumber = _batchNumber
	r.Set("batch_number", _batchNumber)
	return nil
}

// Get BatchNumber Getter
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetBatchNumber() string {
	return r._batchNumber
}
