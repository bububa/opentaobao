package tmallcar

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallcar"
)

/* 
爱卡车系图片数据接入 
tmall.car.xcar.synchronize.car.line.pics.data

爱卡车系图片数据同步天猫汽车
*/
func TmallCarXcarSynchronizeCarLinePicsData(clt *core.SDKClient, req *tmallcar.TmallCarXcarSynchronizeCarLinePicsDataAPIRequest, session string) (*tmallcar.TmallCarXcarSynchronizeCarLinePicsDataAPIResponse, error) {
    var resp tmallcar.TmallCarXcarSynchronizeCarLinePicsDataAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
