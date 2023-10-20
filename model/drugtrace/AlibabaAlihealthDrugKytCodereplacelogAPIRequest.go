package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugkytcodereplacelogAPIRequest 码替换记录查询 API请求
// alibaba.alihealth.drug.kyt.codereplacelog
//
// 码替换记录查询
type AlibabaalihealthdrugkytcodereplacelogAPIRequest struct {
	model.Params
	// 企业ref_ent_id（码申请企业）
	_refEntId string
	// 开始时间（最大查询区间一年）
	_beginTime string
	// 截至时间（最大查询区间一年）
	_endTime string
	// 追溯码（不区分新码、旧码）
	_code string
	// 药品ID
	_drugEntBaseInfoId string
	// 页数（默认一页20条）
	_page int64
}

// NewAlibabaalihealthdrugkytcodereplacelogRequest 初始化AlibabaalihealthdrugkytcodereplacelogAPIRequest对象
func NewAlibabaalihealthdrugkytcodereplacelogRequest() *AlibabaalihealthdrugkytcodereplacelogAPIRequest {
	return &AlibabaalihealthdrugkytcodereplacelogAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugkytcodereplacelogAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.kyt.codereplacelog"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugkytcodereplacelogAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugkytcodereplacelogAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ref_ent_id（码申请企业）
func (r *AlibabaalihealthdrugkytcodereplacelogAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugkytcodereplacelogAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetBeginTime is BeginTime Setter
// 开始时间（最大查询区间一年）
func (r *AlibabaalihealthdrugkytcodereplacelogAPIRequest) SetBeginTime(_beginTime string) error {
	r._beginTime = _beginTime
	r.Set("begin_time", _beginTime)
	return nil
}

// GetBeginTime BeginTime Getter
func (r AlibabaalihealthdrugkytcodereplacelogAPIRequest) GetBeginTime() string {
	return r._beginTime
}

// SetEndTime is EndTime Setter
// 截至时间（最大查询区间一年）
func (r *AlibabaalihealthdrugkytcodereplacelogAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaalihealthdrugkytcodereplacelogAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetCode is Code Setter
// 追溯码（不区分新码、旧码）
func (r *AlibabaalihealthdrugkytcodereplacelogAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugkytcodereplacelogAPIRequest) GetCode() string {
	return r._code
}

// SetDrugEntBaseInfoId is DrugEntBaseInfoId Setter
// 药品ID
func (r *AlibabaalihealthdrugkytcodereplacelogAPIRequest) SetDrugEntBaseInfoId(_drugEntBaseInfoId string) error {
	r._drugEntBaseInfoId = _drugEntBaseInfoId
	r.Set("drug_ent_base_info_id", _drugEntBaseInfoId)
	return nil
}

// GetDrugEntBaseInfoId DrugEntBaseInfoId Getter
func (r AlibabaalihealthdrugkytcodereplacelogAPIRequest) GetDrugEntBaseInfoId() string {
	return r._drugEntBaseInfoId
}

// SetPage is Page Setter
// 页数（默认一页20条）
func (r *AlibabaalihealthdrugkytcodereplacelogAPIRequest) SetPage(_page int64) error {
	r._page = _page
	r.Set("page", _page)
	return nil
}

// GetPage Page Getter
func (r AlibabaalihealthdrugkytcodereplacelogAPIRequest) GetPage() int64 {
	return r._page
}
