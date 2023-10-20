package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafootscanminiquerymobilereportAPIRequest 根据scanId查询报告 API请求
// alibaba.footscan.mini.query.mobilereport
//
// 根据scanId查询报告
type AlibabafootscanminiquerymobilereportAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 扫描ID
	_scanId string
}

// NewAlibabafootscanminiquerymobilereportRequest 初始化AlibabafootscanminiquerymobilereportAPIRequest对象
func NewAlibabafootscanminiquerymobilereportRequest() *AlibabafootscanminiquerymobilereportAPIRequest {
	return &AlibabafootscanminiquerymobilereportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafootscanminiquerymobilereportAPIRequest) GetApiMethodName() string {
	return "alibaba.footscan.mini.query.mobilereport"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafootscanminiquerymobilereportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafootscanminiquerymobilereportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 平台分配的token
func (r *AlibabafootscanminiquerymobilereportAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabafootscanminiquerymobilereportAPIRequest) GetToken() string {
	return r._token
}

// SetScanId is ScanId Setter
// 扫描ID
func (r *AlibabafootscanminiquerymobilereportAPIRequest) SetScanId(_scanId string) error {
	r._scanId = _scanId
	r.Set("scan_id", _scanId)
	return nil
}

// GetScanId ScanId Getter
func (r AlibabafootscanminiquerymobilereportAPIRequest) GetScanId() string {
	return r._scanId
}
