package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest
码流向查询 API请求
alibaba.alihealth.drug.code.kyt.querycodeflow

追溯码流向查询 */
type AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest struct {
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

// NewAlibabaAlihealthDrugCodeKytQuerycodeflowRequest 初始化AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest对象
func NewAlibabaAlihealthDrugCodeKytQuerycodeflowRequest() *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest {
	return &AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.querycodeflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// Get RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// Set is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// Get Code Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetCode() string {
	return r._code
}

// Set is Longitude Setter
// 经度
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// Get Longitude Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetLongitude() string {
	return r._longitude
}

// Set is Latitude Setter
// 纬度
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// Get Latitude Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetLatitude() string {
	return r._latitude
}

// Set is QueryProv Setter
// 查询地所在省
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetQueryProv(_queryProv string) error {
	r._queryProv = _queryProv
	r.Set("query_prov", _queryProv)
	return nil
}

// Get QueryProv Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetQueryProv() string {
	return r._queryProv
}

// Set is QueryCity Setter
// 查询地所在市
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetQueryCity(_queryCity string) error {
	r._queryCity = _queryCity
	r.Set("query_city", _queryCity)
	return nil
}

// Get QueryCity Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetQueryCity() string {
	return r._queryCity
}

// Set is QueryArea Setter
// 查询地所在区
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetQueryArea(_queryArea string) error {
	r._queryArea = _queryArea
	r.Set("query_area", _queryArea)
	return nil
}

// Get QueryArea Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetQueryArea() string {
	return r._queryArea
}

// Set is QueryRegionCode Setter
// 查询地所在区域代码
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetQueryRegionCode(_queryRegionCode string) error {
	r._queryRegionCode = _queryRegionCode
	r.Set("query_region_code", _queryRegionCode)
	return nil
}

// Get QueryRegionCode Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetQueryRegionCode() string {
	return r._queryRegionCode
}

// Set is Detail Setter
// 详细地址
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetDetail(_detail string) error {
	r._detail = _detail
	r.Set("detail", _detail)
	return nil
}

// Get Detail Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetDetail() string {
	return r._detail
}
