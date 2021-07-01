package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增往来单位企业 API请求
alibaba.alihealth.drug.kyt.saveent

新增往来单位企业记录
*/
type AlibabaAlihealthDrugKytSaveentAPIRequest struct {
    model.Params
    // 添加企业唯一标识
    _refEntId   string
    // 新增企业信息
    _addEntReq   *AddEntReqDTO
    // 图片数据流。图片大小务必控制在2M以内
    _licPictureByte   *model.File
}

// 初始化AlibabaAlihealthDrugKytSaveentAPIRequest对象
func NewAlibabaAlihealthDrugKytSaveentRequest() *AlibabaAlihealthDrugKytSaveentAPIRequest{
    return &AlibabaAlihealthDrugKytSaveentAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.saveent"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 添加企业唯一标识
func (r *AlibabaAlihealthDrugKytSaveentAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// AddEntReq Setter
// 新增企业信息
func (r *AlibabaAlihealthDrugKytSaveentAPIRequest) SetAddEntReq(_addEntReq *AddEntReqDTO) error {
    r._addEntReq = _addEntReq
    r.Set("add_ent_req", _addEntReq)
    return nil
}

// AddEntReq Getter
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetAddEntReq() *AddEntReqDTO {
    return r._addEntReq
}
// LicPictureByte Setter
// 图片数据流。图片大小务必控制在2M以内
func (r *AlibabaAlihealthDrugKytSaveentAPIRequest) SetLicPictureByte(_licPictureByte *model.File) error {
    r._licPictureByte = _licPictureByte
    r.Set("lic_picture_byte", _licPictureByte)
    return nil
}

// LicPictureByte Getter
func (r AlibabaAlihealthDrugKytSaveentAPIRequest) GetLicPictureByte() *model.File {
    return r._licPictureByte
}
