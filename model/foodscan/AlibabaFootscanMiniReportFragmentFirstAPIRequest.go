package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafootscanminireportfragmentfirstAPIRequest 第一只脚生成报告接口 API请求
// alibaba.footscan.mini.report.fragment.first
//
// 第一只脚生成报告接口
type AlibabafootscanminireportfragmentfirstAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData *FilePackageRequest
}

// NewAlibabafootscanminireportfragmentfirstRequest 初始化AlibabafootscanminireportfragmentfirstAPIRequest对象
func NewAlibabafootscanminireportfragmentfirstRequest() *AlibabafootscanminireportfragmentfirstAPIRequest {
	return &AlibabafootscanminireportfragmentfirstAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafootscanminireportfragmentfirstAPIRequest) GetApiMethodName() string {
	return "alibaba.footscan.mini.report.fragment.first"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafootscanminireportfragmentfirstAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafootscanminireportfragmentfirstAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 平台分配的token
func (r *AlibabafootscanminireportfragmentfirstAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabafootscanminireportfragmentfirstAPIRequest) GetToken() string {
	return r._token
}

// SetReqData is ReqData Setter
// 请求数据
func (r *AlibabafootscanminireportfragmentfirstAPIRequest) SetReqData(_reqData *FilePackageRequest) error {
	r._reqData = _reqData
	r.Set("req_data", _reqData)
	return nil
}

// GetReqData ReqData Getter
func (r AlibabafootscanminireportfragmentfirstAPIRequest) GetReqData() *FilePackageRequest {
	return r._reqData
}
