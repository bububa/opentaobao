package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest 删除楼盘预售证 API请求
// alibaba.alihouse.newhome.project.presalepermit.delete
//
// 删除楼盘预售证信息
type AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest struct {
	model.Params
	// 外部顾问ID
	_outerPermitId string
	// 楼盘外部id
	_outerId string
	// 店铺外部id
	_outerStoreId string
}

// NewAlibabaalihousenewhomeprojectpresalepermitdeleteRequest 初始化AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest对象
func NewAlibabaalihousenewhomeprojectpresalepermitdeleteRequest() *AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest {
	return &AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.presalepermit.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterPermitId is OuterPermitId Setter
// 外部顾问ID
func (r *AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest) SetOuterPermitId(_outerPermitId string) error {
	r._outerPermitId = _outerPermitId
	r.Set("outer_permit_id", _outerPermitId)
	return nil
}

// GetOuterPermitId OuterPermitId Getter
func (r AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest) GetOuterPermitId() string {
	return r._outerPermitId
}

// SetOuterId is OuterId Setter
// 楼盘外部id
func (r *AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest) SetOuterId(_outerId string) error {
	r._outerId = _outerId
	r.Set("outer_id", _outerId)
	return nil
}

// GetOuterId OuterId Getter
func (r AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest) GetOuterId() string {
	return r._outerId
}

// SetOuterStoreId is OuterStoreId Setter
// 店铺外部id
func (r *AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest) SetOuterStoreId(_outerStoreId string) error {
	r._outerStoreId = _outerStoreId
	r.Set("outer_store_id", _outerStoreId)
	return nil
}

// GetOuterStoreId OuterStoreId Getter
func (r AlibabaalihousenewhomeprojectpresalepermitdeleteAPIRequest) GetOuterStoreId() string {
	return r._outerStoreId
}
