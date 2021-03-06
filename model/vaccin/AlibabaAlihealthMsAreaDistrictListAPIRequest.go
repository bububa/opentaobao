package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMsAreaDistrictListAPIRequest 疫苗预约地市信息查询 API请求
// alibaba.alihealth.ms.area.district.list
//
// 微信小程序疫苗预约地市信息查询
type AlibabaAlihealthMsAreaDistrictListAPIRequest struct {
	model.Params
	// 省份ID
	_divisionId int64
}

// NewAlibabaAlihealthMsAreaDistrictListRequest 初始化AlibabaAlihealthMsAreaDistrictListAPIRequest对象
func NewAlibabaAlihealthMsAreaDistrictListRequest() *AlibabaAlihealthMsAreaDistrictListAPIRequest {
	return &AlibabaAlihealthMsAreaDistrictListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthMsAreaDistrictListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.ms.area.district.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthMsAreaDistrictListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDivisionId is DivisionId Setter
// 省份ID
func (r *AlibabaAlihealthMsAreaDistrictListAPIRequest) SetDivisionId(_divisionId int64) error {
	r._divisionId = _divisionId
	r.Set("division_id", _divisionId)
	return nil
}

// GetDivisionId DivisionId Getter
func (r AlibabaAlihealthMsAreaDistrictListAPIRequest) GetDivisionId() int64 {
	return r._divisionId
}
