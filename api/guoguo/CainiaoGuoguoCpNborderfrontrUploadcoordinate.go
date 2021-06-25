package guoguo

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/guoguo"
)

/* 
上传小件员GPS位置信息 
cainiao.guoguo.cp.nborderfrontr.uploadcoordinate

上传小件员GPS位置信息
*/
func CainiaoGuoguoCpNborderfrontrUploadcoordinate(clt *core.SDKClient, req *guoguo.CainiaoGuoguoCpNborderfrontrUploadcoordinateRequest, session string) (*guoguo.CainiaoGuoguoCpNborderfrontrUploadcoordinateResponse, error) {
    var resp guoguo.CainiaoGuoguoCpNborderfrontrUploadcoordinateAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
