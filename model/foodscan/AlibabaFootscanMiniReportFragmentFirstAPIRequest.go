package foodscan

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniReportFragmentFirstAPIRequest) GetApiMethodName() string {
	return "alibaba.footscan.mini.report.fragment.first"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniReportFragmentFirstAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniReportFragmentFirstAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r AlibabaFootscanMiniReportFragmentFirstAPIRequest) GetToken() string {
	return r._token
}

// Set is ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportFragmentFirstAPIRequest) SetReqData(_reqData *FilePackageRequest) error {
	r._reqData = _reqData
	r.Set("req_data", _reqData)
	return nil
}

// Get ReqData Getter
func (r AlibabaFootscanMiniReportFragmentFirstAPIRequest) GetReqData() *FilePackageRequest {
	return r._reqData
}
