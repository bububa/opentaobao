package examination

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构获取特色服务标签 API请求
alibaba.alihealth.examination.hospital.special.tag

体检机构获取特色服务标签列表
*/
type AlibabaAlihealthExaminationHospitalSpecialTagRequest struct {
    model.Params
}

// 初始化AlibabaAlihealthExaminationHospitalSpecialTagRequest对象
func NewAlibabaAlihealthExaminationHospitalSpecialTagRequest() *AlibabaAlihealthExaminationHospitalSpecialTagRequest{
    return &AlibabaAlihealthExaminationHospitalSpecialTagRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthExaminationHospitalSpecialTagRequest) GetApiMethodName() string {
    return "alibaba.alihealth.examination.hospital.special.tag"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthExaminationHospitalSpecialTagRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
