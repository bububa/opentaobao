package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressSiteTmsSyncAPIRequest 配服务商网点信息同步 API请求
// taobao.logistics.express.site.tms.sync
//
// 配服务商网点信息同步
type TaobaoLogisticsExpressSiteTmsSyncAPIRequest struct {
	model.Params
	// 网点信息
	_siteUpsetRequest *SiteUpsetRequest
}

// NewTaobaoLogisticsExpressSiteTmsSyncRequest 初始化TaobaoLogisticsExpressSiteTmsSyncAPIRequest对象
func NewTaobaoLogisticsExpressSiteTmsSyncRequest() *TaobaoLogisticsExpressSiteTmsSyncAPIRequest {
	return &TaobaoLogisticsExpressSiteTmsSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsExpressSiteTmsSyncAPIRequest) Reset() {
	r._siteUpsetRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressSiteTmsSyncAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.site.tms.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressSiteTmsSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressSiteTmsSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSiteUpsetRequest is SiteUpsetRequest Setter
// 网点信息
func (r *TaobaoLogisticsExpressSiteTmsSyncAPIRequest) SetSiteUpsetRequest(_siteUpsetRequest *SiteUpsetRequest) error {
	r._siteUpsetRequest = _siteUpsetRequest
	r.Set("site_upset_request", _siteUpsetRequest)
	return nil
}

// GetSiteUpsetRequest SiteUpsetRequest Getter
func (r TaobaoLogisticsExpressSiteTmsSyncAPIRequest) GetSiteUpsetRequest() *SiteUpsetRequest {
	return r._siteUpsetRequest
}

var poolTaobaoLogisticsExpressSiteTmsSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsExpressSiteTmsSyncRequest()
	},
}

// GetTaobaoLogisticsExpressSiteTmsSyncRequest 从 sync.Pool 获取 TaobaoLogisticsExpressSiteTmsSyncAPIRequest
func GetTaobaoLogisticsExpressSiteTmsSyncAPIRequest() *TaobaoLogisticsExpressSiteTmsSyncAPIRequest {
	return poolTaobaoLogisticsExpressSiteTmsSyncAPIRequest.Get().(*TaobaoLogisticsExpressSiteTmsSyncAPIRequest)
}

// ReleaseTaobaoLogisticsExpressSiteTmsSyncAPIRequest 将 TaobaoLogisticsExpressSiteTmsSyncAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsExpressSiteTmsSyncAPIRequest(v *TaobaoLogisticsExpressSiteTmsSyncAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsExpressSiteTmsSyncAPIRequest.Put(v)
}
