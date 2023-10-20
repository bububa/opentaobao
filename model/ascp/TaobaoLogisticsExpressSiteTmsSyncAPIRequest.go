package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpresssitetmssyncAPIRequest 配服务商网点信息同步 API请求
// taobao.logistics.express.site.tms.sync
//
// 配服务商网点信息同步
type TaobaologisticsexpresssitetmssyncAPIRequest struct {
	model.Params
	// 网点信息
	_siteUpsetRequest *SiteUpsetRequest
}

// NewTaobaologisticsexpresssitetmssyncRequest 初始化TaobaologisticsexpresssitetmssyncAPIRequest对象
func NewTaobaologisticsexpresssitetmssyncRequest() *TaobaologisticsexpresssitetmssyncAPIRequest {
	return &TaobaologisticsexpresssitetmssyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpresssitetmssyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.site.tms.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpresssitetmssyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpresssitetmssyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSiteUpsetRequest is SiteUpsetRequest Setter
// 网点信息
func (r *TaobaologisticsexpresssitetmssyncAPIRequest) SetSiteUpsetRequest(_siteUpsetRequest *SiteUpsetRequest) error {
	r._siteUpsetRequest = _siteUpsetRequest
	r.Set("site_upset_request", _siteUpsetRequest)
	return nil
}

// GetSiteUpsetRequest SiteUpsetRequest Getter
func (r TaobaologisticsexpresssitetmssyncAPIRequest) GetSiteUpsetRequest() *SiteUpsetRequest {
	return r._siteUpsetRequest
}
