package xhotelitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest 下载飞猪国际城市结果 API请求
// taobao.xhotel.city.coordinates.batch.download
//
// 给国际酒店供应商提供计算对应飞猪城市的服务，免去城市名称匹配流程，加快对接流程
type TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest struct {
	model.Params
	// 上传的经纬度批次号
	_batchId int64
}

// NewTaobaoxhotelcitycoordinatesbatchdownloadRequest 初始化TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest对象
func NewTaobaoxhotelcitycoordinatesbatchdownloadRequest() *TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest {
	return &TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.city.coordinates.batch.download"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBatchId is BatchId Setter
// 上传的经纬度批次号
func (r *TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest) SetBatchId(_batchId int64) error {
	r._batchId = _batchId
	r.Set("batch_id", _batchId)
	return nil
}

// GetBatchId BatchId Getter
func (r TaobaoxhotelcitycoordinatesbatchdownloadAPIRequest) GetBatchId() int64 {
	return r._batchId
}
