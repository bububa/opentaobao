package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest 下载飞猪国际城市结果 API请求
// taobao.xhotel.city.coordinates.batch.download
//
// 给国际酒店供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程
type TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest struct {
	model.Params
	// 上传的经纬度批次号
	_batchId int64
}

// NewTaobaoXhotelCityCoordinatesBatchDownloadRequest 初始化TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest对象
func NewTaobaoXhotelCityCoordinatesBatchDownloadRequest() *TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest {
	return &TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) Reset() {
	r._batchId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.city.coordinates.batch.download"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchId is BatchId Setter
// 上传的经纬度批次号
func (r *TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) SetBatchId(_batchId int64) error {
	r._batchId = _batchId
	r.Set("batch_id", _batchId)
	return nil
}

// GetBatchId BatchId Getter
func (r TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) GetBatchId() int64 {
	return r._batchId
}

var poolTaobaoXhotelCityCoordinatesBatchDownloadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelCityCoordinatesBatchDownloadRequest()
	},
}

// GetTaobaoXhotelCityCoordinatesBatchDownloadRequest 从 sync.Pool 获取 TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest
func GetTaobaoXhotelCityCoordinatesBatchDownloadAPIRequest() *TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest {
	return poolTaobaoXhotelCityCoordinatesBatchDownloadAPIRequest.Get().(*TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest)
}

// ReleaseTaobaoXhotelCityCoordinatesBatchDownloadAPIRequest 将 TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelCityCoordinatesBatchDownloadAPIRequest(v *TaobaoXhotelCityCoordinatesBatchDownloadAPIRequest) {
	v.Reset()
	poolTaobaoXhotelCityCoordinatesBatchDownloadAPIRequest.Put(v)
}
