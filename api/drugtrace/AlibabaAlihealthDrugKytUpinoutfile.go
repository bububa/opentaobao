package drugtrace

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/drugtrace"
)

/* 
上传出入库单据(传文件) 
alibaba.alihealth.drug.kyt.upinoutfile

上传出入库单据(传文件)
*/
func AlibabaAlihealthDrugKytUpinoutfile(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUpinoutfileRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytUpinoutfileAPIResponse, error) {
    var resp drugtrace.AlibabaAlihealthDrugKytUpinoutfileAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
