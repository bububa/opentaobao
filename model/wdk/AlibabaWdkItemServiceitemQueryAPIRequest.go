package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkitemserviceitemqueryAPIRequest 查询服务商品 API请求
// alibaba.wdk.item.serviceitem.query
//
// 查询服务商品
type AlibabawdkitemserviceitemqueryAPIRequest struct {
	model.Params
	// 后台类目
	_hemaCategoryId string
	// 机构编码
	_orgNo string
}

// NewAlibabawdkitemserviceitemqueryRequest 初始化AlibabawdkitemserviceitemqueryAPIRequest对象
func NewAlibabawdkitemserviceitemqueryRequest() *AlibabawdkitemserviceitemqueryAPIRequest {
	return &AlibabawdkitemserviceitemqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkitemserviceitemqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.serviceitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkitemserviceitemqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkitemserviceitemqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHemaCategoryId is HemaCategoryId Setter
// 后台类目
func (r *AlibabawdkitemserviceitemqueryAPIRequest) SetHemaCategoryId(_hemaCategoryId string) error {
	r._hemaCategoryId = _hemaCategoryId
	r.Set("hema_category_id", _hemaCategoryId)
	return nil
}

// GetHemaCategoryId HemaCategoryId Getter
func (r AlibabawdkitemserviceitemqueryAPIRequest) GetHemaCategoryId() string {
	return r._hemaCategoryId
}

// SetOrgNo is OrgNo Setter
// 机构编码
func (r *AlibabawdkitemserviceitemqueryAPIRequest) SetOrgNo(_orgNo string) error {
	r._orgNo = _orgNo
	r.Set("org_no", _orgNo)
	return nil
}

// GetOrgNo OrgNo Getter
func (r AlibabawdkitemserviceitemqueryAPIRequest) GetOrgNo() string {
	return r._orgNo
}
