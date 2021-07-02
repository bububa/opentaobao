package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFootscanMiniReportListAPIRequest 查询报告列表 API请求
// alibaba.footscan.mini.report.list
//
// 查询报告列表
type AlibabaFootscanMiniReportListAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData *TobFeetModelMobileReportRequest
}

// NewAlibabaFootscanMiniReportListRequest 初始化AlibabaFootscanMiniReportListAPIRequest对象
func NewAlibabaFootscanMiniReportListRequest() *AlibabaFootscanMiniReportListAPIRequest {
	return &AlibabaFootscanMiniReportListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFootscanMiniReportListAPIRequest) GetApiMethodName() string {
	return "alibaba.footscan.mini.report.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFootscanMiniReportListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Token Setter
// 平台分配的token
func (r *AlibabaFootscanMiniReportListAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r AlibabaFootscanMiniReportListAPIRequest) GetToken() string {
	return r._token
}

// Set is ReqData Setter
// 请求数据
func (r *AlibabaFootscanMiniReportListAPIRequest) SetReqData(_reqData *TobFeetModelMobileReportRequest) error {
	r._reqData = _reqData
	r.Set("req_data", _reqData)
	return nil
}

// Get ReqData Getter
func (r AlibabaFootscanMiniReportListAPIRequest) GetReqData() *TobFeetModelMobileReportRequest {
	return r._reqData
}
