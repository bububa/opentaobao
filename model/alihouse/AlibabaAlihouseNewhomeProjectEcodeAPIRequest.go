package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectEcodeAPIRequest 楼盘e码更新 API请求
// alibaba.alihouse.newhome.project.ecode
//
// 更新楼盘ecode
type AlibabaAlihouseNewhomeProjectEcodeAPIRequest struct {
	model.Params
	// 楼盘外部id
	_outerId string
	// 楼盘e码
	_eCode string
	// 外部门店ID
	_outerStoreId string
}

// NewAlibabaAlihouseNewhomeProjectEcodeRequest 初始化AlibabaAlihouseNewhomeProjectEcodeAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectEcodeRequest() *AlibabaAlihouseNewhomeProjectEcodeAPIRequest {
	return &AlibabaAlihouseNewhomeProjectEcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.ecode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterId is OuterId Setter
// 楼盘外部id
func (r *AlibabaAlihouseNewhomeProjectEcodeAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetECode is ECode Setter
// 楼盘e码
func (r *AlibabaAlihouseNewhomeProjectEcodeAPIRequest) SetECode(_eCode string) error {
	r._eCode = _eCode
	r.Set("e_code", _eCode)
	return nil
}

// GetECode ECode Getter
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetECode() string {
	return r._eCode
}

// SetOuterStoreId is OuterStoreId Setter
// 外部门店ID
func (r *AlibabaAlihouseNewhomeProjectEcodeAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseNewhomeProjectEcodeAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}
