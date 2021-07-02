package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkHrworkbenchCdpempsQueryAPIRequest homs员工信息核对查询服务 API请求
// alibaba.wdk.hrworkbench.cdpemps.query
//
// 给盒马可靠软件服务商Cdp系统，做非阿里编员工数据一致性核对检查
type AlibabaWdkHrworkbenchCdpempsQueryAPIRequest struct {
	model.Params
	// 页面大小
	_pageSize int64
	// 业务授权key
	_bizKey string
	// 业务授权code
	_bizCode string
	// 起始页
	_currentPage int64
}

// NewAlibabaWdkHrworkbenchCdpempsQueryRequest 初始化AlibabaWdkHrworkbenchCdpempsQueryAPIRequest对象
func NewAlibabaWdkHrworkbenchCdpempsQueryRequest() *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest {
	return &AlibabaWdkHrworkbenchCdpempsQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.hrworkbench.cdpemps.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPageSize is PageSize Setter
// 页面大小
func (r *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetBizKey is BizKey Setter
// 业务授权key
func (r *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) SetBizKey(_bizKey string) error {
	r._bizKey = _bizKey
	r.Set("biz_key", _bizKey)
	return nil
}

// GetBizKey BizKey Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetBizKey() string {
	return r._bizKey
}

// SetBizCode is BizCode Setter
// 业务授权code
func (r *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) SetBizCode(_bizCode string) error {
	r._bizCode = _bizCode
	r.Set("biz_code", _bizCode)
	return nil
}

// GetBizCode BizCode Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetBizCode() string {
	return r._bizCode
}

// SetCurrentPage is CurrentPage Setter
// 起始页
func (r *AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) SetCurrentPage(_currentPage int64) error {
	r._currentPage = _currentPage
	r.Set("current_page", _currentPage)
	return nil
}

// GetCurrentPage CurrentPage Getter
func (r AlibabaWdkHrworkbenchCdpempsQueryAPIRequest) GetCurrentPage() int64 {
	return r._currentPage
}
