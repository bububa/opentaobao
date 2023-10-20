package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterspserviceorderepocuploadAPIRequest 电子保单文件上传接口 API请求
// tmall.servicecenter.spserviceorder.epoc.upload
//
// 电子保单文件上传接口
type TmallservicecenterspserviceorderepocuploadAPIRequest struct {
	model.Params
	// 电子保单文件数据
	_epocFileData string
	// 电子保单文件名
	_epocFileName string
	// t&c文件数据
	_tcFileData string
	// t&c文件名
	_tcFileName string
	// 服务交易订单号
	_bizOrderId int64
}

// NewTmallservicecenterspserviceorderepocuploadRequest 初始化TmallservicecenterspserviceorderepocuploadAPIRequest对象
func NewTmallservicecenterspserviceorderepocuploadRequest() *TmallservicecenterspserviceorderepocuploadAPIRequest {
	return &TmallservicecenterspserviceorderepocuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterspserviceorderepocuploadAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.spserviceorder.epoc.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterspserviceorderepocuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterspserviceorderepocuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEpocFileData is EpocFileData Setter
// 电子保单文件数据
func (r *TmallservicecenterspserviceorderepocuploadAPIRequest) SetEpocFileData(_epocFileData string) error {
	r._epocFileData = _epocFileData
	r.Set("epoc_file_data", _epocFileData)
	return nil
}

// GetEpocFileData EpocFileData Getter
func (r TmallservicecenterspserviceorderepocuploadAPIRequest) GetEpocFileData() string {
	return r._epocFileData
}

// SetEpocFileName is EpocFileName Setter
// 电子保单文件名
func (r *TmallservicecenterspserviceorderepocuploadAPIRequest) SetEpocFileName(_epocFileName string) error {
	r._epocFileName = _epocFileName
	r.Set("epoc_file_name", _epocFileName)
	return nil
}

// GetEpocFileName EpocFileName Getter
func (r TmallservicecenterspserviceorderepocuploadAPIRequest) GetEpocFileName() string {
	return r._epocFileName
}

// SetTcFileData is TcFileData Setter
// t&amp;c文件数据
func (r *TmallservicecenterspserviceorderepocuploadAPIRequest) SetTcFileData(_tcFileData string) error {
	r._tcFileData = _tcFileData
	r.Set("tc_file_data", _tcFileData)
	return nil
}

// GetTcFileData TcFileData Getter
func (r TmallservicecenterspserviceorderepocuploadAPIRequest) GetTcFileData() string {
	return r._tcFileData
}

// SetTcFileName is TcFileName Setter
// t&amp;c文件名
func (r *TmallservicecenterspserviceorderepocuploadAPIRequest) SetTcFileName(_tcFileName string) error {
	r._tcFileName = _tcFileName
	r.Set("tc_file_name", _tcFileName)
	return nil
}

// GetTcFileName TcFileName Getter
func (r TmallservicecenterspserviceorderepocuploadAPIRequest) GetTcFileName() string {
	return r._tcFileName
}

// SetBizOrderId is BizOrderId Setter
// 服务交易订单号
func (r *TmallservicecenterspserviceorderepocuploadAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TmallservicecenterspserviceorderepocuploadAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
