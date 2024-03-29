package tmallcms

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest) Reset() {
	r._url = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest) GetApiMethodName() string {
	return "tmall.marketing.liuliangbao.spreadlink.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallMarketingLiuliangbaoSpreadlinkCreateRequest()
	},
}

// GetTmallMarketingLiuliangbaoSpreadlinkCreateRequest 从 sync.Pool 获取 TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest
func GetTmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest() *TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest {
	return poolTmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest.Get().(*TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest)
}

// ReleaseTmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest 将 TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest 放入 sync.Pool
func ReleaseTmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest(v *TmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest) {
	v.Reset()
	poolTmallMarketingLiuliangbaoSpreadlinkCreateAPIRequest.Put(v)
}
