package xhotelitem

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/xhotelitem"
)

/* 
下载飞猪国际城市结果 
taobao.xhotel.city.coordinates.batch.download

给国际酒店供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程
*/
func TaobaoXhotelCityCoordinatesBatchDownload(clt *core.SDKClient, req *xhotelitem.TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest, session string) (*xhotelitem.TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse, error) {
    var resp xhotelitem.TaobaoXhotelCityCoordinatesBatchDownloadAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
