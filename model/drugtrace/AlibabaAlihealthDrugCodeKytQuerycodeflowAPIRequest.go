package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest 码流向查询 API请求
// alibaba.alihealth.drug.code.kyt.querycodeflow
//
// 追溯码流向查询
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
		Params: model.NewParams(9),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) Reset() {
	r._refEntId = ""
	r._code = ""
	r._longitude = ""
	r._latitude = ""
	r._queryProv = ""
	r._queryCity = ""
	r._queryArea = ""
	r._queryRegionCode = ""
	r._detail = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.code.kyt.querycodeflow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRefEntId is RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetRefEntId() string {
	return r._refEntId
}

// SetCode is Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetCode(_code string) error {
	r._code = _code
	r.Set("code", _code)
	return nil
}

// GetCode Code Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetCode() string {
	return r._code
}

// SetLongitude is Longitude Setter
// 经度
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetLatitude is Latitude Setter
// 纬度
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetQueryProv is QueryProv Setter
// 查询地所在省
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetQueryProv(_queryProv string) error {
	r._queryProv = _queryProv
	r.Set("query_prov", _queryProv)
	return nil
}

// GetQueryProv QueryProv Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetQueryProv() string {
	return r._queryProv
}

// SetQueryCity is QueryCity Setter
// 查询地所在市
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetQueryCity(_queryCity string) error {
	r._queryCity = _queryCity
	r.Set("query_city", _queryCity)
	return nil
}

// GetQueryCity QueryCity Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetQueryCity() string {
	return r._queryCity
}

// SetQueryArea is QueryArea Setter
// 查询地所在区
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetQueryArea(_queryArea string) error {
	r._queryArea = _queryArea
	r.Set("query_area", _queryArea)
	return nil
}

// GetQueryArea QueryArea Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetQueryArea() string {
	return r._queryArea
}

// SetQueryRegionCode is QueryRegionCode Setter
// 查询地所在区域代码
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetQueryRegionCode(_queryRegionCode string) error {
	r._queryRegionCode = _queryRegionCode
	r.Set("query_region_code", _queryRegionCode)
	return nil
}

// GetQueryRegionCode QueryRegionCode Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetQueryRegionCode() string {
	return r._queryRegionCode
}

// SetDetail is Detail Setter
// 详细地址
func (r *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) SetDetail(_detail string) error {
	r._detail = _detail
	r.Set("detail", _detail)
	return nil
}

// GetDetail Detail Getter
func (r AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) GetDetail() string {
	return r._detail
}

var poolAlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugCodeKytQuerycodeflowRequest()
	},
}

// GetAlibabaAlihealthDrugCodeKytQuerycodeflowRequest 从 sync.Pool 获取 AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest
func GetAlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest() *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest {
	return poolAlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest.Get().(*AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest)
}

// ReleaseAlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest 将 AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest(v *AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest.Put(v)
}
