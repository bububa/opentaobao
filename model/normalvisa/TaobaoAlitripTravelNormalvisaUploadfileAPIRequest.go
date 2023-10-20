package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitriptravelnormalvisauploadfileAPIRequest 上传电子签证 API请求
// taobao.alitrip.travel.normalvisa.uploadfile
//
// 上传电子签证
type TaobaoalitriptravelnormalvisauploadfileAPIRequest struct {
	model.Params
	// 文件名：注意文件名请保证和上传的文件一直
	_fileName string
	// 文件
	_fileBytes *model.File
	// 订单id
	_bizOrderId int64
}

// NewTaobaoalitriptravelnormalvisauploadfileRequest 初始化TaobaoalitriptravelnormalvisauploadfileAPIRequest对象
func NewTaobaoalitriptravelnormalvisauploadfileRequest() *TaobaoalitriptravelnormalvisauploadfileAPIRequest {
	return &TaobaoalitriptravelnormalvisauploadfileAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitriptravelnormalvisauploadfileAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.uploadfile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitriptravelnormalvisauploadfileAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitriptravelnormalvisauploadfileAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 文件名：注意文件名请保证和上传的文件一直
func (r *TaobaoalitriptravelnormalvisauploadfileAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r TaobaoalitriptravelnormalvisauploadfileAPIRequest) GetFileName() string {
	return r._fileName
}

// SetFileBytes is FileBytes Setter
// 文件
func (r *TaobaoalitriptravelnormalvisauploadfileAPIRequest) SetFileBytes(_fileBytes *model.File) error {
	r._fileBytes = _fileBytes
	r.Set("file_bytes", _fileBytes)
	return nil
}

// GetFileBytes FileBytes Getter
func (r TaobaoalitriptravelnormalvisauploadfileAPIRequest) GetFileBytes() *model.File {
	return r._fileBytes
}

// SetBizOrderId is BizOrderId Setter
// 订单id
func (r *TaobaoalitriptravelnormalvisauploadfileAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoalitriptravelnormalvisauploadfileAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}
