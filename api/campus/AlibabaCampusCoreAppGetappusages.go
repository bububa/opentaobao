package campus

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/campus"
)

/* 
根据应用ID获得应用所在的园区 
alibaba.campus.core.app.getappusages

传入应用的id,  获得用户授权的园区
*/
func AlibabaCampusCoreAppGetappusages(clt *core.SDKClient, req *campus.AlibabaCampusCoreAppGetappusagesAPIRequest, session string) (*campus.AlibabaCampusCoreAppGetappusagesAPIResponse, error) {
    var resp campus.AlibabaCampusCoreAppGetappusagesAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
