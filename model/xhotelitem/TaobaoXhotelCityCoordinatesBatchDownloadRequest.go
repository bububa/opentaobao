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
type TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest struct {
    model.Params
    // 上传的经纬度批次号
    _batchId   int64
}

// 初始化TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest对象
func NewTaobaoXhotelCityCoordinatesBatchDownloadRequest() *TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest{
    return &TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) GetApiMethodName() string {
    return "taobao.xhotel.city.coordinates.batch.download"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BatchId Setter
// 上传的经纬度批次号
func (r *TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) SetBatchId(_batchId int64) error {
    r._batchId = _batchId
    r.Set("batch_id", _batchId)
    return nil
}

// BatchId Getter
func (r TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) GetBatchId() int64 {
    return r._batchId
}
