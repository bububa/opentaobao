package hotelhstdf

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelHotelMessageReceiveAPIRequest 接收道消息接口 API请求
// taobao.xhotel.hotel.message.receive
//
// 接收道消息接口
type TaobaoXhotelHotelMessageReceiveAPIRequest struct {
	model.Params
	// 文件在OSS的存储路径，不包含文件名，只包含文件夹路径
	_ossPath string
	// 上传的文件个数
	_fileCount int64
	// 文件记录总数量
	_recordCount int64
	// 数据的类型，枚举值见文档
	_dataType int64
}

// NewTaobaoXhotelHotelMessageReceiveRequest 初始化TaobaoXhotelHotelMessageReceiveAPIRequest对象
func NewTaobaoXhotelHotelMessageReceiveRequest() *TaobaoXhotelHotelMessageReceiveAPIRequest {
	return &TaobaoXhotelHotelMessageReceiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelHotelMessageReceiveAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.hotel.message.receive"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelHotelMessageReceiveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelHotelMessageReceiveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOssPath is OssPath Setter
// 文件在OSS的存储路径，不包含文件名，只包含文件夹路径
func (r *TaobaoXhotelHotelMessageReceiveAPIRequest) SetOssPath(_ossPath string) error {
	r._ossPath = _ossPath
	r.Set("oss_path", _ossPath)
	return nil
}

// GetOssPath OssPath Getter
func (r TaobaoXhotelHotelMessageReceiveAPIRequest) GetOssPath() string {
	return r._ossPath
}

// SetFileCount is FileCount Setter
// 上传的文件个数
func (r *TaobaoXhotelHotelMessageReceiveAPIRequest) SetFileCount(_fileCount int64) error {
	r._fileCount = _fileCount
	r.Set("file_count", _fileCount)
	return nil
}

// GetFileCount FileCount Getter
func (r TaobaoXhotelHotelMessageReceiveAPIRequest) GetFileCount() int64 {
	return r._fileCount
}

// SetRecordCount is RecordCount Setter
// 文件记录总数量
func (r *TaobaoXhotelHotelMessageReceiveAPIRequest) SetRecordCount(_recordCount int64) error {
	r._recordCount = _recordCount
	r.Set("record_count", _recordCount)
	return nil
}

// GetRecordCount RecordCount Getter
func (r TaobaoXhotelHotelMessageReceiveAPIRequest) GetRecordCount() int64 {
	return r._recordCount
}

// SetDataType is DataType Setter
// 数据的类型，枚举值见文档
func (r *TaobaoXhotelHotelMessageReceiveAPIRequest) SetDataType(_dataType int64) error {
	r._dataType = _dataType
	r.Set("data_type", _dataType)
	return nil
}

// GetDataType DataType Getter
func (r TaobaoXhotelHotelMessageReceiveAPIRequest) GetDataType() int64 {
	return r._dataType
}
