package lstvending

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstvendingcargospacesaveAPIRequest 自动售卖机货道数据回流 API请求
// alibaba.lst.vending.cargospace.save
//
// 自动售卖机货道数据回流接口，ISV通过调用此接口上传售卖机货道信息。
type AlibabalstvendingcargospacesaveAPIRequest struct {
	model.Params
	// 货道信息
	_cargoSpaceDTOList []VendingCargoSpaceDto
}

// NewAlibabalstvendingcargospacesaveRequest 初始化AlibabalstvendingcargospacesaveAPIRequest对象
func NewAlibabalstvendingcargospacesaveRequest() *AlibabalstvendingcargospacesaveAPIRequest {
	return &AlibabalstvendingcargospacesaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstvendingcargospacesaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.cargospace.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstvendingcargospacesaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstvendingcargospacesaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCargoSpaceDTOList is CargoSpaceDTOList Setter
// 货道信息
func (r *AlibabalstvendingcargospacesaveAPIRequest) SetCargoSpaceDTOList(_cargoSpaceDTOList []VendingCargoSpaceDto) error {
	r._cargoSpaceDTOList = _cargoSpaceDTOList
	r.Set("cargo_space_d_t_o_list", _cargoSpaceDTOList)
	return nil
}

// GetCargoSpaceDTOList CargoSpaceDTOList Getter
func (r AlibabalstvendingcargospacesaveAPIRequest) GetCargoSpaceDTOList() []VendingCargoSpaceDto {
	return r._cargoSpaceDTOList
}
