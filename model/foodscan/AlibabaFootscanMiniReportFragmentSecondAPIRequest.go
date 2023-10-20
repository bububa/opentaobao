package foodscan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFootscanMiniReportFragmentSecondAPIRequest 第二只脚生成报告接口 API请求
// alibaba.footscan.mini.report.fragment.second
//
// 第二只脚生成报告接口
type AlibabaFootscanMiniReportFragmentSecondAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData *FilePackageBasicReq
}

// NewAlibabaFootscanMiniReportFragmentSecondRequest 初始化AlibabaFootscanMiniReportFragmentSecondAPIRequest对象
func NewAlibabaFootscanMiniReportFragmentSecondRequest() *AlibabaFootscanMiniReportFragmentSecondAPIRequest {
	return &AlibabaFootscanMiniReportFragmentSecondAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFootscanMiniReportFragmentSecondAPIRequest) Reset() {
	r._token = ""
	r._reqData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniReportFragmentSecondAPIRequest) GetApiMethodName() string {
	return "alibaba.footscan.mini.report.fragment.second"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniReportFragmentSecondAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFootscanMiniReportFragmentSecondAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniReportFragmentSecondAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabaFootscanMiniReportFragmentSecondAPIRequest) GetToken() string {
	return r._token
}

// SetReqData is ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportFragmentSecondAPIRequest) SetReqData(_reqData *FilePackageBasicReq) error {
	r._reqData = _reqData
	r.Set("req_data", _reqData)
	return nil
}

// GetReqData ReqData Getter
func (r AlibabaFootscanMiniReportFragmentSecondAPIRequest) GetReqData() *FilePackageBasicReq {
	return r._reqData
}

var poolAlibabaFootscanMiniReportFragmentSecondAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFootscanMiniReportFragmentSecondRequest()
	},
}

// GetAlibabaFootscanMiniReportFragmentSecondRequest 从 sync.Pool 获取 AlibabaFootscanMiniReportFragmentSecondAPIRequest
func GetAlibabaFootscanMiniReportFragmentSecondAPIRequest() *AlibabaFootscanMiniReportFragmentSecondAPIRequest {
	return poolAlibabaFootscanMiniReportFragmentSecondAPIRequest.Get().(*AlibabaFootscanMiniReportFragmentSecondAPIRequest)
}

// ReleaseAlibabaFootscanMiniReportFragmentSecondAPIRequest 将 AlibabaFootscanMiniReportFragmentSecondAPIRequest 放入 sync.Pool
func ReleaseAlibabaFootscanMiniReportFragmentSecondAPIRequest(v *AlibabaFootscanMiniReportFragmentSecondAPIRequest) {
	v.Reset()
	poolAlibabaFootscanMiniReportFragmentSecondAPIRequest.Put(v)
}
