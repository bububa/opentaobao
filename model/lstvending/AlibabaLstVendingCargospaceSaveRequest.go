package lstvending

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
自动售卖机货道数据回流 API请求
alibaba.lst.vending.cargospace.save

自动售卖机货道数据回流接口，ISV通过调用此接口上传售卖机货道信息。
*/
type AlibabaLstVendingCargospaceSaveRequest struct {
    model.Params
    // 货道信息
    cargoSpaceDTOList   []VendingCargoSpaceDto
}

// 初始化AlibabaLstVendingCargospaceSaveRequest对象
func NewAlibabaLstVendingCargospaceSaveRequest() *AlibabaLstVendingCargospaceSaveRequest{
    return &AlibabaLstVendingCargospaceSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaLstVendingCargospaceSaveRequest) GetApiMethodName() string {
    return "alibaba.lst.vending.cargospace.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaLstVendingCargospaceSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CargoSpaceDTOList Setter
// 货道信息
func (r *AlibabaLstVendingCargospaceSaveRequest) SetCargoSpaceDTOList(cargoSpaceDTOList []VendingCargoSpaceDto) error {
    r.cargoSpaceDTOList = cargoSpaceDTOList
    r.Set("cargo_space_d_t_o_list", cargoSpaceDTOList)
    return nil
}

// CargoSpaceDTOList Getter
func (r AlibabaLstVendingCargospaceSaveRequest) GetCargoSpaceDTOList() []VendingCargoSpaceDto {
    return r.cargoSpaceDTOList
}
