package xhotelitem

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
下载飞猪国际城市结果 API请求
taobao.xhotel.city.coordinates.batch.download

给国际酒店供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程
*/
type TaobaoXhotelCityCoordinatesBatchDownloadRequest struct {
    model.Params
    // 上传的经纬度批次号
    _batchId   int64
}

// 初始化TaobaoXhotelCityCoordinatesBatchDownloadRequest对象
func NewTaobaoXhotelCityCoordinatesBatchDownloadRequest() *TaobaoXhotelCityCoordinatesBatchDownloadRequest{
    return &TaobaoXhotelCityCoordinatesBatchDownloadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCityCoordinatesBatchDownloadRequest) GetApiMethodName() string {
    return "taobao.xhotel.city.coordinates.batch.download"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCityCoordinatesBatchDownloadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchId Setter
// 上传的经纬度批次号
func (r *TaobaoXhotelCityCoordinatesBatchDownloadRequest) SetBatchId(_batchId int64) error {
    r._batchId = _batchId
    r.Set("batch_id", _batchId)
    return nil
}

// BatchId Getter
func (r TaobaoXhotelCityCoordinatesBatchDownloadRequest) GetBatchId() int64 {
    return r._batchId
}
