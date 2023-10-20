package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest 获取盲底文件中的批次信息 API请求
// alibaba.alihealth.drugcode.drugfactory.blindfile.getbatchinfo
//
// 获取盲底文件中的批次信息
type AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest struct {
	model.Params
	// 企业Id
	_refEntId string
	// 药品子类编码
	_subTypeNo string
	// 盲底文件上传开始时间（yyyy-MM-dd）
	_blindFileStartDate string
	// 盲底文件上传结束时间（yyyy-MM-dd）
	_blindFileEndDate string
}

// NewAlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoRequest 初始化AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest对象
func NewAlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoRequest() *AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest {
	return &AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.drugfactory.blindfile.getbatchinfo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业Id
func (r *AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetSubTypeNo is SubTypeNo Setter
// 药品子类编码
func (r *AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) SetSubTypeNo(_subTypeNo string) error {
	r._subTypeNo = _subTypeNo
	r.Set("sub_type_no", _subTypeNo)
	return nil
}

// GetSubTypeNo SubTypeNo Getter
func (r AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) GetSubTypeNo() string {
	return r._subTypeNo
}

// SetBlindFileStartDate is BlindFileStartDate Setter
// 盲底文件上传开始时间（yyyy-MM-dd）
func (r *AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) SetBlindFileStartDate(_blindFileStartDate string) error {
	r._blindFileStartDate = _blindFileStartDate
	r.Set("blind_file_start_date", _blindFileStartDate)
	return nil
}

// GetBlindFileStartDate BlindFileStartDate Getter
func (r AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) GetBlindFileStartDate() string {
	return r._blindFileStartDate
}

// SetBlindFileEndDate is BlindFileEndDate Setter
// 盲底文件上传结束时间（yyyy-MM-dd）
func (r *AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) SetBlindFileEndDate(_blindFileEndDate string) error {
	r._blindFileEndDate = _blindFileEndDate
	r.Set("blind_file_end_date", _blindFileEndDate)
	return nil
}

// GetBlindFileEndDate BlindFileEndDate Getter
func (r AlibabaalihealthdrugcodedrugfactoryblindfilegetbatchinfoAPIRequest) GetBlindFileEndDate() string {
	return r._blindFileEndDate
}
