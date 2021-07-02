package foodscan

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniReportFragmentSecondAPIRequest) GetApiMethodName() string {
	return "alibaba.footscan.mini.report.fragment.second"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniReportFragmentSecondAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniReportFragmentSecondAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r AlibabaFootscanMiniReportFragmentSecondAPIRequest) GetToken() string {
	return r._token
}

// Set is ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportFragmentSecondAPIRequest) SetReqData(_reqData *FilePackageBasicReq) error {
	r._reqData = _reqData
	r.Set("req_data", _reqData)
	return nil
}

// Get ReqData Getter
func (r AlibabaFootscanMiniReportFragmentSecondAPIRequest) GetReqData() *FilePackageBasicReq {
	return r._reqData
}
