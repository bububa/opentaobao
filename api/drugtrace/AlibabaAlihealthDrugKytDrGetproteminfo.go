package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
疫苗，获取生产企业的存储和运输温度 
alibaba.alihealth.drug.kyt.dr.getproteminfo

疫苗，获取生产企业的存储和运输温度
*/
func AlibabaAlihealthDrugKytDrGetproteminfo(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrGetproteminfoRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytDrGetproteminfoAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytDrGetproteminfoAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
