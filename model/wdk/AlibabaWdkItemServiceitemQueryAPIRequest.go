package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkItemServiceitemQueryAPIRequest 查询服务商品 API请求
// alibaba.wdk.item.serviceitem.query
//
// 查询服务商品
type AlibabaWdkItemServiceitemQueryAPIRequest struct {
	model.Params
	// 后台类目
	_hemaCategoryId string
	// 机构编码
	_orgNo string
}

// NewAlibabaWdkItemServiceitemQueryRequest 初始化AlibabaWdkItemServiceitemQueryAPIRequest对象
func NewAlibabaWdkItemServiceitemQueryRequest() *AlibabaWdkItemServiceitemQueryAPIRequest {
	return &AlibabaWdkItemServiceitemQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkItemServiceitemQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.item.serviceitem.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkItemServiceitemQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetHemaCategoryId is HemaCategoryId Setter
// 后台类目
func (r *AlibabaWdkItemServiceitemQueryAPIRequest) SetHemaCategoryId(_hemaCategoryId string) error {
	r._hemaCategoryId = _hemaCategoryId
	r.Set("hema_category_id", _hemaCategoryId)
	return nil
}

// GetHemaCategoryId HemaCategoryId Getter
func (r AlibabaWdkItemServiceitemQueryAPIRequest) GetHemaCategoryId() string {
	return r._hemaCategoryId
}

// SetOrgNo is OrgNo Setter
// 机构编码
func (r *AlibabaWdkItemServiceitemQueryAPIRequest) SetOrgNo(_orgNo string) error {
	r._orgNo = _orgNo
	r.Set("org_no", _orgNo)
	return nil
}

// GetOrgNo OrgNo Getter
func (r AlibabaWdkItemServiceitemQueryAPIRequest) GetOrgNo() string {
	return r._orgNo
}
