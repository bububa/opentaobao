package lstvending

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingCargospaceSaveAPIRequest 自动售卖机货道数据回流 API请求
// alibaba.lst.vending.cargospace.save
//
// 自动售卖机货道数据回流接口，ISV通过调用此接口上传售卖机货道信息。
type AlibabaLstVendingCargospaceSaveAPIRequest struct {
	model.Params
	// 货道信息
	_cargoSpaceDTOList []VendingCargoSpaceDto
}

// NewAlibabaLstVendingCargospaceSaveRequest 初始化AlibabaLstVendingCargospaceSaveAPIRequest对象
func NewAlibabaLstVendingCargospaceSaveRequest() *AlibabaLstVendingCargospaceSaveAPIRequest {
	return &AlibabaLstVendingCargospaceSaveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLstVendingCargospaceSaveAPIRequest) Reset() {
	r._cargoSpaceDTOList = r._cargoSpaceDTOList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingCargospaceSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.vending.cargospace.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingCargospaceSaveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstVendingCargospaceSaveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCargoSpaceDTOList is CargoSpaceDTOList Setter
// 货道信息
func (r *AlibabaLstVendingCargospaceSaveAPIRequest) SetCargoSpaceDTOList(_cargoSpaceDTOList []VendingCargoSpaceDto) error {
	r._cargoSpaceDTOList = _cargoSpaceDTOList
	r.Set("cargo_space_d_t_o_list", _cargoSpaceDTOList)
	return nil
}

// GetCargoSpaceDTOList CargoSpaceDTOList Getter
func (r AlibabaLstVendingCargospaceSaveAPIRequest) GetCargoSpaceDTOList() []VendingCargoSpaceDto {
	return r._cargoSpaceDTOList
}

var poolAlibabaLstVendingCargospaceSaveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLstVendingCargospaceSaveRequest()
	},
}

// GetAlibabaLstVendingCargospaceSaveRequest 从 sync.Pool 获取 AlibabaLstVendingCargospaceSaveAPIRequest
func GetAlibabaLstVendingCargospaceSaveAPIRequest() *AlibabaLstVendingCargospaceSaveAPIRequest {
	return poolAlibabaLstVendingCargospaceSaveAPIRequest.Get().(*AlibabaLstVendingCargospaceSaveAPIRequest)
}

// ReleaseAlibabaLstVendingCargospaceSaveAPIRequest 将 AlibabaLstVendingCargospaceSaveAPIRequest 放入 sync.Pool
func ReleaseAlibabaLstVendingCargospaceSaveAPIRequest(v *AlibabaLstVendingCargospaceSaveAPIRequest) {
	v.Reset()
	poolAlibabaLstVendingCargospaceSaveAPIRequest.Put(v)
}
