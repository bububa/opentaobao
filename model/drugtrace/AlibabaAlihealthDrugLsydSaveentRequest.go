package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售药店往来单位新增 API请求
alibaba.alihealth.drug.lsyd.saveent

新增往来单位企业记录
*/
type AlibabaAlihealthDrugLsydSaveentRequest struct {
    model.Params
    // 添加企业唯一标识
    refEntId   string
    // 新增企业信息
    addEntReq   *AddEntReqDto
    // 图片数据流。图片大小务必控制在2M以内
    licPictureByte   []*model.File
}

// 初始化AlibabaAlihealthDrugLsydSaveentRequest对象
func NewAlibabaAlihealthDrugLsydSaveentRequest() *AlibabaAlihealthDrugLsydSaveentRequest{
    return &AlibabaAlihealthDrugLsydSaveentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugLsydSaveentRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.lsyd.saveent"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugLsydSaveentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 添加企业唯一标识
func (r *AlibabaAlihealthDrugLsydSaveentRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugLsydSaveentRequest) GetRefEntId() string {
    return r.refEntId
}
// AddEntReq Setter
// 新增企业信息
func (r *AlibabaAlihealthDrugLsydSaveentRequest) SetAddEntReq(addEntReq *AddEntReqDto) error {
    r.addEntReq = addEntReq
    r.Set("add_ent_req", addEntReq)
    return nil
}

// AddEntReq Getter
func (r AlibabaAlihealthDrugLsydSaveentRequest) GetAddEntReq() *AddEntReqDto {
    return r.addEntReq
}
// LicPictureByte Setter
// 图片数据流。图片大小务必控制在2M以内
func (r *AlibabaAlihealthDrugLsydSaveentRequest) SetLicPictureByte(licPictureByte []*model.File) error {
    r.licPictureByte = licPictureByte
    r.Set("lic_picture_byte", licPictureByte)
    return nil
}

// LicPictureByte Getter
func (r AlibabaAlihealthDrugLsydSaveentRequest) GetLicPictureByte() []*model.File {
    return r.licPictureByte
}
