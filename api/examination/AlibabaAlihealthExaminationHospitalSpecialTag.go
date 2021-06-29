package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
体检机构获取特色服务标签 
alibaba.alihealth.examination.hospital.special.tag

体检机构获取特色服务标签列表
*/
func AlibabaAlihealthExaminationHospitalSpecialTag(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationHospitalSpecialTagRequest, session string) (*examination.AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationHospitalSpecialTagAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
