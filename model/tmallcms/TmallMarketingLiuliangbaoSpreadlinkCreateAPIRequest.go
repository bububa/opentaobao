package tmallcms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmarketingliuliangbaospreadlinkcreateAPIRequest 创建流量宝活动链接 API请求
// tmall.marketing.liuliangbao.spreadlink.create
//
// 通过源活动链接和商家ID，创建流量宝活动链接
type TmallmarketingliuliangbaospreadlinkcreateAPIRequest struct {
	model.Params
	// 活动链接，必须为淘系链接
	_url string
}

// NewTmallmarketingliuliangbaospreadlinkcreateRequest 初始化TmallmarketingliuliangbaospreadlinkcreateAPIRequest对象
func NewTmallmarketingliuliangbaospreadlinkcreateRequest() *TmallmarketingliuliangbaospreadlinkcreateAPIRequest {
	return &TmallmarketingliuliangbaospreadlinkcreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmarketingliuliangbaospreadlinkcreateAPIRequest) GetApiMethodName() string {
	return "tmall.marketing.liuliangbao.spreadlink.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmarketingliuliangbaospreadlinkcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmarketingliuliangbaospreadlinkcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUrl is Url Setter
// 活动链接，必须为淘系链接
func (r *TmallmarketingliuliangbaospreadlinkcreateAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TmallmarketingliuliangbaospreadlinkcreateAPIRequest) GetUrl() string {
	return r._url
}
