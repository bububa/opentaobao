package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest 接收盲底文件删除日志 API请求
// alibaba.alihealth.drugcode.drugfactory.blindfiledellog
//
// 临床用药试验-接收盲底文件删除日志
type AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest struct {
	model.Params
	// 药厂企业id
	_refEntId string
	// 盲底文件名称，多个盲底文件用,分隔
	_blindFileNames string
	// 操作人
	_operator string
	// 盲底文件删除时间
	_blindFileDeleteTime string
}

// NewAlibabaalihealthdrugcodedrugfactoryblindfiledellogRequest 初始化AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest对象
func NewAlibabaalihealthdrugcodedrugfactoryblindfiledellogRequest() *AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest {
	return &AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.blindfiledellog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 药厂企业id
func (r *AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBlindFileNames is BlindFileNames Setter
// 盲底文件名称，多个盲底文件用,分隔
func (r *AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) SetBlindFileNames(_blindFileNames string) error {
	r._blindFileNames = _blindFileNames
	r.Set("blind_file_names", _blindFileNames)
	return nil
}

// GetBlindFileNames BlindFileNames Getter
func (r AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) GetBlindFileNames() string {
	return r._blindFileNames
}

// SetOperator is Operator Setter
// 操作人
func (r *AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) GetOperator() string {
	return r._operator
}

// SetBlindFileDeleteTime is BlindFileDeleteTime Setter
// 盲底文件删除时间
func (r *AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) SetBlindFileDeleteTime(_blindFileDeleteTime string) error {
	r._blindFileDeleteTime = _blindFileDeleteTime
	r.Set("blind_file_delete_time", _blindFileDeleteTime)
	return nil
}

// GetBlindFileDeleteTime BlindFileDeleteTime Getter
func (r AlibabaalihealthdrugcodedrugfactoryblindfiledellogAPIRequest) GetBlindFileDeleteTime() string {
	return r._blindFileDeleteTime
}
