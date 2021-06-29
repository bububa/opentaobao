package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增往来单位企业 APIRequest
alibaba.alihealth.drug.kyt.saveent

新增往来单位企业记录
*/
type AlibabaAlihealthDrugKytSaveentRequest struct {
    model.Params

    // 添加企业唯一标识
    refEntId   string 

    // 新增企业信息
    addEntReq   *AddEntReqDto 

    // 图片数据流。图片大小务必控制在2M以内
    licPictureByte   []*model.File 

}

func NewAlibabaAlihealthDrugKytSaveentRequest() *AlibabaAlihealthDrugKytSaveentRequest{
    return &AlibabaAlihealthDrugKytSaveentRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytSaveentRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.saveent"
}

func (r AlibabaAlihealthDrugKytSaveentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytSaveentRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytSaveentRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytSaveentRequest) SetAddEntReq(addEntReq *AddEntReqDto) error {
    r.addEntReq = addEntReq
    r.Set("add_ent_req", addEntReq)
    return nil
}

func (r AlibabaAlihealthDrugKytSaveentRequest) GetAddEntReq() *AddEntReqDto {
    return r.addEntReq
}

func (r *AlibabaAlihealthDrugKytSaveentRequest) SetLicPictureByte(licPictureByte []*model.File) error {
    r.licPictureByte = licPictureByte
    r.Set("lic_picture_byte", licPictureByte)
    return nil
}

func (r AlibabaAlihealthDrugKytSaveentRequest) GetLicPictureByte() []*model.File {
    return r.licPictureByte
}

