package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest 获取盲底文件处理结果 API请求
// alibaba.alihealth.drugcode.drugfactory.getblindresult
//
// 获取盲底文件处理结果
type AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 盲底文件名称
	_blindFileName string
}

// NewAlibabaalihealthdrugcodedrugfactorygetblindresultRequest 初始化AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest对象
func NewAlibabaalihealthdrugcodedrugfactorygetblindresultRequest() *AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest {
	return &AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.getblindresult"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业id
func (r *AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBlindFileName is BlindFileName Setter
// 盲底文件名称
func (r *AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest) SetBlindFileName(_blindFileName string) error {
	r._blindFileName = _blindFileName
	r.Set("blind_file_name", _blindFileName)
	return nil
}

// GetBlindFileName BlindFileName Getter
func (r AlibabaalihealthdrugcodedrugfactorygetblindresultAPIRequest) GetBlindFileName() string {
	return r._blindFileName
}
