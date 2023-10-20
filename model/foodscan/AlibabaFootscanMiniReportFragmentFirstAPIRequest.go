package foodscan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFootscanMiniReportFragmentFirstAPIRequest 第一只脚生成报告接口 API请求
// alibaba.footscan.mini.report.fragment.first
//
// 第一只脚生成报告接口
type AlibabaFootscanMiniReportFragmentFirstAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData *FilePackageRequest
}

// NewAlibabaFootscanMiniReportFragmentFirstRequest 初始化AlibabaFootscanMiniReportFragmentFirstAPIRequest对象
func NewAlibabaFootscanMiniReportFragmentFirstRequest() *AlibabaFootscanMiniReportFragmentFirstAPIRequest {
	return &AlibabaFootscanMiniReportFragmentFirstAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFootscanMiniReportFragmentFirstAPIRequest) Reset() {
	r._token = ""
	r._reqData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniReportFragmentFirstAPIRequest) GetApiMethodName() string {
	return "alibaba.footscan.mini.report.fragment.first"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniReportFragmentFirstAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFootscanMiniReportFragmentFirstAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniReportFragmentFirstAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaFootscanMiniReportFragmentFirstAPIRequest) GetToken() string {
	return r._token
}

// SetReqData is ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportFragmentFirstAPIRequest) SetReqData(_reqData *FilePackageRequest) error {
	r._reqData = _reqData
	r.Set("req_data", _reqData)
	return nil
}

// GetReqData ReqData Getter
func (r AlibabaFootscanMiniReportFragmentFirstAPIRequest) GetReqData() *FilePackageRequest {
	return r._reqData
}

var poolAlibabaFootscanMiniReportFragmentFirstAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFootscanMiniReportFragmentFirstRequest()
	},
}

// GetAlibabaFootscanMiniReportFragmentFirstRequest 从 sync.Pool 获取 AlibabaFootscanMiniReportFragmentFirstAPIRequest
func GetAlibabaFootscanMiniReportFragmentFirstAPIRequest() *AlibabaFootscanMiniReportFragmentFirstAPIRequest {
	return poolAlibabaFootscanMiniReportFragmentFirstAPIRequest.Get().(*AlibabaFootscanMiniReportFragmentFirstAPIRequest)
}

// ReleaseAlibabaFootscanMiniReportFragmentFirstAPIRequest 将 AlibabaFootscanMiniReportFragmentFirstAPIRequest 放入 sync.Pool
func ReleaseAlibabaFootscanMiniReportFragmentFirstAPIRequest(v *AlibabaFootscanMiniReportFragmentFirstAPIRequest) {
	v.Reset()
	poolAlibabaFootscanMiniReportFragmentFirstAPIRequest.Put(v)
}
