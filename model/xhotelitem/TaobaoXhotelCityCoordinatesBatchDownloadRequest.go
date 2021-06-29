package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下载飞猪国际城市结果 APIRequest
taobao.xhotel.city.coordinates.batch.download

给国际酒店供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程
*/
type TaobaoXhotelCityCoordinatesBatchDownloadRequest struct {
    model.Params

    // 上传的经纬度批次号
    batchId   int64 

}

func NewTaobaoXhotelCityCoordinatesBatchDownloadRequest() *TaobaoXhotelCityCoordinatesBatchDownloadRequest{
    return &TaobaoXhotelCityCoordinatesBatchDownloadRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoXhotelCityCoordinatesBatchDownloadRequest) GetApiMethodName() string {
    return "taobao.xhotel.city.coordinates.batch.download"
}

func (r TaobaoXhotelCityCoordinatesBatchDownloadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoXhotelCityCoordinatesBatchDownloadRequest) SetBatchId(batchId int64) error {
    r.batchId = batchId
    r.Set("batch_id", batchId)
    return nil
}

func (r TaobaoXhotelCityCoordinatesBatchDownloadRequest) GetBatchId() int64 {
    return r.batchId
}

