package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsmediaresourcesuploadAPIRequest 图片与视频上传 API请求
// taobao.logistics.media.resources.upload
//
// 图片与视频上传
type TaobaologisticsmediaresourcesuploadAPIRequest struct {
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

// NewTaobaologisticsmediaresourcesuploadRequest 初始化TaobaologisticsmediaresourcesuploadAPIRequest对象
func NewTaobaologisticsmediaresourcesuploadRequest() *TaobaologisticsmediaresourcesuploadAPIRequest {
	return &TaobaologisticsmediaresourcesuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsmediaresourcesuploadAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.media.resources.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsmediaresourcesuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsmediaresourcesuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetName is Name Setter
// 图片/视频名称，只支持后缀：.jpg、.png、.mp4
func (r *TaobaologisticsmediaresourcesuploadAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r TaobaologisticsmediaresourcesuploadAPIRequest) GetName() string {
	return r._name
}

// SetSupplierId is SupplierId Setter
// 淘天物流服务商ID
func (r *TaobaologisticsmediaresourcesuploadAPIRequest) SetSupplierId(_supplierId int64) error {
	r._supplierId = _supplierId
	r.Set("supplier_id", _supplierId)
	return nil
}

// GetSupplierId SupplierId Getter
func (r TaobaologisticsmediaresourcesuploadAPIRequest) GetSupplierId() int64 {
	return r._supplierId
}

// SetType is Type Setter
// 1-图片，2-视频
func (r *TaobaologisticsmediaresourcesuploadAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaologisticsmediaresourcesuploadAPIRequest) GetType() int64 {
	return r._type
}

// SetData is Data Setter
// 图片/视频的字节数组
func (r *TaobaologisticsmediaresourcesuploadAPIRequest) SetData(_data *model.File) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaologisticsmediaresourcesuploadAPIRequest) GetData() *model.File {
	return r._data
}
