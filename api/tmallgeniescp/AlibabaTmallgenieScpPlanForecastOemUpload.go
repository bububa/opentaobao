package tmallgeniescp

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallgeniescp"
)

/* 
16-供应商预测（OEM-成品）回传接口 
alibaba.tmallgenie.scp.plan.forecast.oem.upload

供应商预测（OEM-成品）回传接口
*/
func AlibabaTmallgenieScpPlanForecastOemUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanForecastOemUploadRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse, error) {
    var resp tmallgeniescp.AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
