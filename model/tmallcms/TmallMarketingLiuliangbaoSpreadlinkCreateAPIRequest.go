package tmallcms

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest 创建流量宝活动链接 API请求
// tmall.marketing.liuliangbao.spreadlink.create
//
// 通过源活动链接和商家ID，创建流量宝活动链接
type TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest struct {
	model.Params
	// 活动链接，必须为淘系链接
	_url string
}

// NewTmallMarketingLiuliangbaoSpreadlinkCreateRequest 初始化TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest对象
func NewTmallMarketingLiuliangbaoSpreadlinkCreateRequest() *TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest {
	return &TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest) GetApiMethodName() string {
	return "tmall.marketing.liuliangbao.spreadlink.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUrl is Url Setter
// 活动链接，必须为淘系链接
func (r *TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest) SetUrl(_url string) error {
	r._url = _url
	r.Set("url", _url)
	return nil
}

// GetUrl Url Getter
func (r TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest) GetUrl() string {
	return r._url
}
