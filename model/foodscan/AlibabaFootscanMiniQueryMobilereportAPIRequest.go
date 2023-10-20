package foodscan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFootscanMiniQueryMobilereportAPIRequest 根据scanId查询报告 API请求
// alibaba.footscan.mini.query.mobilereport
//
// 根据scanId查询报告
type AlibabaFootscanMiniQueryMobilereportAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 扫描ID
	_scanId string
}

// NewAlibabaFootscanMiniQueryMobilereportRequest 初始化AlibabaFootscanMiniQueryMobilereportAPIRequest对象
func NewAlibabaFootscanMiniQueryMobilereportRequest() *AlibabaFootscanMiniQueryMobilereportAPIRequest {
	return &AlibabaFootscanMiniQueryMobilereportAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFootscanMiniQueryMobilereportAPIRequest) Reset() {
	r._token = ""
	r._scanId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniQueryMobilereportAPIRequest) GetApiMethodName() string {
	return "alibaba.footscan.mini.query.mobilereport"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniQueryMobilereportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFootscanMiniQueryMobilereportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniQueryMobilereportAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaFootscanMiniQueryMobilereportAPIRequest) GetToken() string {
	return r._token
}

// SetScanId is ScanId Setter
// 扫描ID
func (r *AlibabaFootscanMiniQueryMobilereportAPIRequest) SetScanId(_scanId string) error {
	r._scanId = _scanId
	r.Set("scan_id", _scanId)
	return nil
}

// GetScanId ScanId Getter
func (r AlibabaFootscanMiniQueryMobilereportAPIRequest) GetScanId() string {
	return r._scanId
}

var poolAlibabaFootscanMiniQueryMobilereportAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFootscanMiniQueryMobilereportRequest()
	},
}

// GetAlibabaFootscanMiniQueryMobilereportRequest 从 sync.Pool 获取 AlibabaFootscanMiniQueryMobilereportAPIRequest
func GetAlibabaFootscanMiniQueryMobilereportAPIRequest() *AlibabaFootscanMiniQueryMobilereportAPIRequest {
	return poolAlibabaFootscanMiniQueryMobilereportAPIRequest.Get().(*AlibabaFootscanMiniQueryMobilereportAPIRequest)
}

// ReleaseAlibabaFootscanMiniQueryMobilereportAPIRequest 将 AlibabaFootscanMiniQueryMobilereportAPIRequest 放入 sync.Pool
func ReleaseAlibabaFootscanMiniQueryMobilereportAPIRequest(v *AlibabaFootscanMiniQueryMobilereportAPIRequest) {
	v.Reset()
	poolAlibabaFootscanMiniQueryMobilereportAPIRequest.Put(v)
}
