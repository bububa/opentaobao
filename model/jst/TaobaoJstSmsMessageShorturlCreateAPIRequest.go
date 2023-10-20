package jst

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsMessageShorturlCreateAPIRequest 聚石塔营销效果短链生成 API请求
// taobao.jst.sms.message.shorturl.create
//
// 聚石塔生成淘短链接口
type TaobaoJstSmsMessageShorturlCreateAPIRequest struct {
	model.Params
	// 人群标签
	_tag string
	// 商品或者店铺的H5地址，只支持长链
	_url string
	// 批次号
	_batchNumber string
	// 是否需要https前缀： true-要  false-不要
	_needHttpsPrefix bool
}

// NewTaobaoJstSmsMessageShorturlCreateRequest 初始化TaobaoJstSmsMessageShorturlCreateAPIRequest对象
func NewTaobaoJstSmsMessageShorturlCreateRequest() *TaobaoJstSmsMessageShorturlCreateAPIRequest {
	return &TaobaoJstSmsMessageShorturlCreateAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoJstSmsMessageShorturlCreateAPIRequest) Reset() {
	r._tag = ""
	r._url = ""
	r._batchNumber = ""
	r._needHttpsPrefix = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.message.shorturl.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTag is Tag Setter
// 人群标签
func (r *TaobaoJstSmsMessageShorturlCreateAPIRequest) SetTag(_tag string) error {
	r._tag = _tag
	r.Set("tag", _tag)
	return nil
}

// GetTag Tag Getter
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetTag() string {
	return r._tag
}

// SetUrl is Url Setter
// 商品或者店铺的H5地址，只支持长链
func (r *TaobaoJstSmsMessageShorturlCreateAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetUrl() string {
	return r._url
}

// SetBatchNumber is BatchNumber Setter
// 批次号
func (r *TaobaoJstSmsMessageShorturlCreateAPIRequest) SetBatchNumber(_batchNumber string) error {
	r._batchNumber = _batchNumber
	r.Set("batch_number", _batchNumber)
	return nil
}

// GetBatchNumber BatchNumber Getter
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetBatchNumber() string {
	return r._batchNumber
}

// SetNeedHttpsPrefix is NeedHttpsPrefix Setter
// 是否需要https前缀： true-要  false-不要
func (r *TaobaoJstSmsMessageShorturlCreateAPIRequest) SetNeedHttpsPrefix(_needHttpsPrefix bool) error {
	r._needHttpsPrefix = _needHttpsPrefix
	r.Set("need_https_prefix", _needHttpsPrefix)
	return nil
}

// GetNeedHttpsPrefix NeedHttpsPrefix Getter
func (r TaobaoJstSmsMessageShorturlCreateAPIRequest) GetNeedHttpsPrefix() bool {
	return r._needHttpsPrefix
}

var poolTaobaoJstSmsMessageShorturlCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoJstSmsMessageShorturlCreateRequest()
	},
}

// GetTaobaoJstSmsMessageShorturlCreateRequest 从 sync.Pool 获取 TaobaoJstSmsMessageShorturlCreateAPIRequest
func GetTaobaoJstSmsMessageShorturlCreateAPIRequest() *TaobaoJstSmsMessageShorturlCreateAPIRequest {
	return poolTaobaoJstSmsMessageShorturlCreateAPIRequest.Get().(*TaobaoJstSmsMessageShorturlCreateAPIRequest)
}

// ReleaseTaobaoJstSmsMessageShorturlCreateAPIRequest 将 TaobaoJstSmsMessageShorturlCreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoJstSmsMessageShorturlCreateAPIRequest(v *TaobaoJstSmsMessageShorturlCreateAPIRequest) {
	v.Reset()
	poolTaobaoJstSmsMessageShorturlCreateAPIRequest.Put(v)
}
