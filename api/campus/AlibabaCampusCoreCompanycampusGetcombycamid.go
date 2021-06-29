package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据园区ID获取运营公司信息 
alibaba.campus.core.companycampus.getcombycamid

根据园区ID获取运营公司信息
*/
func AlibabaCampusCoreCompanycampusGetcombycamid(clt *core.SDKClient, req *campus.AlibabaCampusCoreCompanycampusGetcombycamidRequest, session string) (*campus.AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse, error) {
    var resp campus.AlibabaCampusCoreCompanycampusGetcombycamidAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}