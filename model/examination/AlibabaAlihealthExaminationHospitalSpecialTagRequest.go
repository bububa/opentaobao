package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构获取特色服务标签 APIRequest
alibaba.alihealth.examination.hospital.special.tag

体检机构获取特色服务标签列表
*/
type AlibabaAlihealthExaminationHospitalSpecialTagRequest struct {
    model.Params

}

func NewAlibabaAlihealthExaminationHospitalSpecialTagRequest() *AlibabaAlihealthExaminationHospitalSpecialTagRequest{
    return &AlibabaAlihealthExaminationHospitalSpecialTagRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthExaminationHospitalSpecialTagRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.hospital.special.tag"
}

func (r AlibabaAlihealthExaminationHospitalSpecialTagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


