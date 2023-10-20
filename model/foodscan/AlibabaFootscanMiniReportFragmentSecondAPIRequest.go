package foodscan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafootscanminireportfragmentsecondAPIRequest 第二只脚生成报告接口 API请求
// alibaba.footscan.mini.report.fragment.second
//
// 第二只脚生成报告接口
type AlibabafootscanminireportfragmentsecondAPIRequest struct {
	model.Params
	// 平台分配的token
	_token string
	// 请求数据
	_reqData *FilePackageBasicReq
}

// NewAlibabafootscanminireportfragmentsecondRequest 初始化AlibabafootscanminireportfragmentsecondAPIRequest对象
func NewAlibabafootscanminireportfragmentsecondRequest() *AlibabafootscanminireportfragmentsecondAPIRequest {
	return &AlibabafootscanminireportfragmentsecondAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafootscanminireportfragmentsecondAPIRequest) GetApiMethodName() string {
	return "alibaba.footscan.mini.report.fragment.second"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafootscanminireportfragmentsecondAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafootscanminireportfragmentsecondAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// 平台分配的token
func (r *AlibabafootscanminireportfragmentsecondAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r AlibabafootscanminireportfragmentsecondAPIRequest) GetToken() string {
	return r._token
}

// SetReqData is ReqData Setter
// 请求数据
func (r *AlibabafootscanminireportfragmentsecondAPIRequest) SetReqData(_reqData *FilePackageBasicReq) error {
	r._reqData = _reqData
	r.Set("req_data", _reqData)
	return nil
}

// GetReqData ReqData Getter
func (r AlibabafootscanminireportfragmentsecondAPIRequest) GetReqData() *FilePackageBasicReq {
	return r._reqData
}
