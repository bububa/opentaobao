package examination

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/examination"
)

/* 
体检机构对接_门店发布／更新 
alibaba.alihealth.examination.hospital.publish

第三方B端有新的门店发布，或者老的门店更新的时候，使用这个接口
*/
func AlibabaAlihealthExaminationHospitalPublish(clt *core.SDKClient, req *examination.AlibabaAlihealthExaminationHospitalPublishRequest, session string) (*examination.AlibabaAlihealthExaminationHospitalPublishAPIResponse, error) {
    var resp examination.AlibabaAlihealthExaminationHospitalPublishAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
