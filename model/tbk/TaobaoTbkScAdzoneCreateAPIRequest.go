package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscadzonecreateAPIRequest 淘宝客-服务商-创建推广者位 API请求
// taobao.tbk.sc.adzone.create
//
// 提供淘宝客创建广告位
type TaobaotbkscadzonecreateAPIRequest struct {
	model.Params
	// 广告位名称，最大长度64字符
	_adzonename string
	// 网站ID
	_siteid int64
}

// NewTaobaotbkscadzonecreateRequest 初始化TaobaotbkscadzonecreateAPIRequest对象
func NewTaobaotbkscadzonecreateRequest() *TaobaotbkscadzonecreateAPIRequest {
	return &TaobaotbkscadzonecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscadzonecreateAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.adzone.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscadzonecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscadzonecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdzonename is Adzonename Setter
// 广告位名称，最大长度64字符
func (r *TaobaotbkscadzonecreateAPIRequest) SetAdzonename(_adzonename string) error {
	r._adzonename = _adzonename
	r.Set("adzone_name", _adzonename)
	return nil
}

// GetAdzonename Adzonename Getter
func (r TaobaotbkscadzonecreateAPIRequest) GetAdzonename() string {
	return r._adzonename
}

// SetSiteid is Siteid Setter
// 网站ID
func (r *TaobaotbkscadzonecreateAPIRequest) SetSiteid(_siteid int64) error {
	r._siteid = _siteid
	r.Set("site_id", _siteid)
	return nil
}

// GetSiteid Siteid Getter
func (r TaobaotbkscadzonecreateAPIRequest) GetSiteid() int64 {
	return r._siteid
}
