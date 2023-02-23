package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScAdzoneCreateAPIRequest 淘宝客-服务商-创建推广者位 API请求
// taobao.tbk.sc.adzone.create
//
// 提供淘宝客创建广告位
type TaobaoTbkScAdzoneCreateAPIRequest struct {
	model.Params
	// 广告位名称，最大长度64字符
	_adzoneName string
	// 网站ID
	_siteId int64
}

// NewTaobaoTbkScAdzoneCreateRequest 初始化TaobaoTbkScAdzoneCreateAPIRequest对象
func NewTaobaoTbkScAdzoneCreateRequest() *TaobaoTbkScAdzoneCreateAPIRequest {
	return &TaobaoTbkScAdzoneCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScAdzoneCreateAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.adzone.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScAdzoneCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkScAdzoneCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdzoneName is AdzoneName Setter
// 广告位名称，最大长度64字符
func (r *TaobaoTbkScAdzoneCreateAPIRequest) SetAdzoneName(_adzoneName string) error {
	r._adzoneName = _adzoneName
	r.Set("adzone_name", _adzoneName)
	return nil
}

// GetAdzoneName AdzoneName Getter
func (r TaobaoTbkScAdzoneCreateAPIRequest) GetAdzoneName() string {
	return r._adzoneName
}

// SetSiteId is SiteId Setter
// 网站ID
func (r *TaobaoTbkScAdzoneCreateAPIRequest) SetSiteId(_siteId int64) error {
	r._siteId = _siteId
	r.Set("site_id", _siteId)
	return nil
}

// GetSiteId SiteId Getter
func (r TaobaoTbkScAdzoneCreateAPIRequest) GetSiteId() int64 {
	return r._siteId
}
