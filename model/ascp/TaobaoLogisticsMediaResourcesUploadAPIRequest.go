package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsMediaResourcesUploadAPIRequest 图片与视频上传 API请求
// taobao.logistics.media.resources.upload
//
// 图片与视频上传
type TaobaoLogisticsMediaResourcesUploadAPIRequest struct {
	model.Params
	// 图片/视频名称，只支持后缀：.jpg、.png、.mp4
	_name string
	// 淘天物流服务商ID
	_supplierId int64
	// 1-图片，2-视频
	_type int64
	// 图片/视频的字节数组
	_data *model.File
}

// NewTaobaoLogisticsMediaResourcesUploadRequest 初始化TaobaoLogisticsMediaResourcesUploadAPIRequest对象
func NewTaobaoLogisticsMediaResourcesUploadRequest() *TaobaoLogisticsMediaResourcesUploadAPIRequest {
	return &TaobaoLogisticsMediaResourcesUploadAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsMediaResourcesUploadAPIRequest) Reset() {
	r._name = ""
	r._supplierId = 0
	r._type = 0
	r._data = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsMediaResourcesUploadAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.media.resources.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsMediaResourcesUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsMediaResourcesUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 图片/视频名称，只支持后缀：.jpg、.png、.mp4
func (r *TaobaoLogisticsMediaResourcesUploadAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaoLogisticsMediaResourcesUploadAPIRequest) GetName() string {
	return r._name
}

// SetSupplierId is SupplierId Setter
// 淘天物流服务商ID
func (r *TaobaoLogisticsMediaResourcesUploadAPIRequest) SetSupplierId(_supplierId int64) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// GetSupplierId SupplierId Getter
func (r TaobaoLogisticsMediaResourcesUploadAPIRequest) GetSupplierId() int64 {
	return r._supplierId
}

// SetType is Type Setter
// 1-图片，2-视频
func (r *TaobaoLogisticsMediaResourcesUploadAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoLogisticsMediaResourcesUploadAPIRequest) GetType() int64 {
	return r._type
}

// SetData is Data Setter
// 图片/视频的字节数组
func (r *TaobaoLogisticsMediaResourcesUploadAPIRequest) SetData(_data *model.File) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaoLogisticsMediaResourcesUploadAPIRequest) GetData() *model.File {
	return r._data
}

var poolTaobaoLogisticsMediaResourcesUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsMediaResourcesUploadRequest()
	},
}

// GetTaobaoLogisticsMediaResourcesUploadRequest 从 sync.Pool 获取 TaobaoLogisticsMediaResourcesUploadAPIRequest
func GetTaobaoLogisticsMediaResourcesUploadAPIRequest() *TaobaoLogisticsMediaResourcesUploadAPIRequest {
	return poolTaobaoLogisticsMediaResourcesUploadAPIRequest.Get().(*TaobaoLogisticsMediaResourcesUploadAPIRequest)
}

// ReleaseTaobaoLogisticsMediaResourcesUploadAPIRequest 将 TaobaoLogisticsMediaResourcesUploadAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsMediaResourcesUploadAPIRequest(v *TaobaoLogisticsMediaResourcesUploadAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsMediaResourcesUploadAPIRequest.Put(v)
}
