package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodekytquerycodeflowAPIRequest 码流向查询 API请求
// alibaba.alihealth.drug.code.kyt.querycodeflow
//
// 追溯码流向查询
type AlibabaalihealthdrugcodekytquerycodeflowAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 追溯码
	_code string
	// 经度
	_longitude string
	// 纬度
	_latitude string
	// 查询地所在省
	_queryProv string
	// 查询地所在市
	_queryCity string
	// 查询地所在区
	_queryArea string
	// 查询地所在区域代码
	_queryRegionCode string
	// 详细地址
	_detail string
}

// NewAlibabaalihealthdrugcodekytquerycodeflowRequest 初始化AlibabaalihealthdrugcodekytquerycodeflowAPIRequest对象
func NewAlibabaalihealthdrugcodekytquerycodeflowRequest() *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest {
	return &AlibabaalihealthdrugcodekytquerycodeflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.querycodeflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetCode() string {
	return r._code
}

// SetLongitude is Longitude Setter
// 经度
func (r *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 纬度
func (r *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetQueryProv is QueryProv Setter
// 查询地所在省
func (r *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) SetQueryProv(_queryProv string) error {
	r._queryProv = _queryProv
	r.Set("query_prov", _queryProv)
	return nil
}

// GetQueryProv QueryProv Getter
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetQueryProv() string {
	return r._queryProv
}

// SetQueryCity is QueryCity Setter
// 查询地所在市
func (r *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) SetQueryCity(_queryCity string) error {
	r._queryCity = _queryCity
	r.Set("query_city", _queryCity)
	return nil
}

// GetQueryCity QueryCity Getter
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetQueryCity() string {
	return r._queryCity
}

// SetQueryArea is QueryArea Setter
// 查询地所在区
func (r *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) SetQueryArea(_queryArea string) error {
	r._queryArea = _queryArea
	r.Set("query_area", _queryArea)
	return nil
}

// GetQueryArea QueryArea Getter
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetQueryArea() string {
	return r._queryArea
}

// SetQueryRegionCode is QueryRegionCode Setter
// 查询地所在区域代码
func (r *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) SetQueryRegionCode(_queryRegionCode string) error {
	r._queryRegionCode = _queryRegionCode
	r.Set("query_region_code", _queryRegionCode)
	return nil
}

// GetQueryRegionCode QueryRegionCode Getter
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetQueryRegionCode() string {
	return r._queryRegionCode
}

// SetDetail is Detail Setter
// 详细地址
func (r *AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) SetDetail(_detail string) error {
	r._detail = _detail
	r.Set("detail", _detail)
	return nil
}

// GetDetail Detail Getter
func (r AlibabaalihealthdrugcodekytquerycodeflowAPIRequest) GetDetail() string {
	return r._detail
}
