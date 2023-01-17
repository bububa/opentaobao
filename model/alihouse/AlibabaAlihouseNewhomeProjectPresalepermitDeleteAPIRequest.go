package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest 删除楼盘预售证 API请求
// alibaba.alihouse.newhome.project.presalepermit.delete
//
// 删除楼盘预售证信息
type AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest struct {
	model.Params
	// 外部顾问ID
	_outerPermitId string
	// 楼盘外部id
	_outerId string
	// 店铺外部id
	_outerStoreId string
}

// NewAlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest 初始化AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectPresalepermitDeleteRequest() *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest {
	return &AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.presalepermit.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterPermitId is OuterPermitId Setter
// 外部顾问ID
func (r *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) SetOuterPermitId(_outerPermitId string) error {
	r._outerPermitId = _outerPermitId
	r.Set("outer_permit_id", _outerPermitId)
	return nil
}

// GetOuterPermitId OuterPermitId Getter
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) GetOuterPermitId() string {
	return r._outerPermitId
}

// SetOuterId is OuterId Setter
// 楼盘外部id
func (r *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterStoreId is OuterStoreId Setter
// 店铺外部id
func (r *AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaAlihouseNewhomeProjectPresalepermitDeleteAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}
