package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOmniorderStoreSdtconsignAPIRequest
通知菜鸟裹裹发货 API请求
taobao.omniorder.store.sdtconsign

ISV取完单号后通知菜鸟裹裹发货 */
type TaobaoOmniorderStoreSdtconsignAPIRequest struct {
	model.Params
	// 取号接口返回的包裹id
	_packageId string
	// 发货标签号
	_tagCode string
}

// NewTaobaoOmniorderStoreSdtconsignRequest 初始化TaobaoOmniorderStoreSdtconsignAPIRequest对象
func NewTaobaoOmniorderStoreSdtconsignRequest() *TaobaoOmniorderStoreSdtconsignAPIRequest {
	return &TaobaoOmniorderStoreSdtconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSdtconsignAPIRequest) GetApiMethodName() string {
	return "taobao.omniorder.store.sdtconsign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSdtconsignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is PackageId Setter
// 取号接口返回的包裹id
func (r *TaobaoOmniorderStoreSdtconsignAPIRequest) SetPackageId(_packageId string) error {
	r._packageId = _packageId
	r.Set("package_id", _packageId)
	return nil
}

// Get PackageId Getter
func (r TaobaoOmniorderStoreSdtconsignAPIRequest) GetPackageId() string {
	return r._packageId
}

// Set is TagCode Setter
// 发货标签号
func (r *TaobaoOmniorderStoreSdtconsignAPIRequest) SetTagCode(_tagCode string) error {
	r._tagCode = _tagCode
	r.Set("tag_code", _tagCode)
	return nil
}

// Get TagCode Getter
func (r TaobaoOmniorderStoreSdtconsignAPIRequest) GetTagCode() string {
	return r._tagCode
}
