package normalvisa

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelNormalvisaUploadfileAPIRequest 上传电子签证 API请求
// taobao.alitrip.travel.normalvisa.uploadfile
//
// 上传电子签证
type TaobaoAlitripTravelNormalvisaUploadfileAPIRequest struct {
	model.Params
	// 文件名：注意文件名请保证和上传的文件一直
	_fileName string
	// 文件
	_fileBytes *model.File
	// 订单id
	_bizOrderId int64
}

// NewTaobaoAlitripTravelNormalvisaUploadfileRequest 初始化TaobaoAlitripTravelNormalvisaUploadfileAPIRequest对象
func NewTaobaoAlitripTravelNormalvisaUploadfileRequest() *TaobaoAlitripTravelNormalvisaUploadfileAPIRequest {
	return &TaobaoAlitripTravelNormalvisaUploadfileAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) Reset() {
	r._fileName = ""
	r._fileBytes = nil
	r._bizOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.travel.normalvisa.uploadfile"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetFileName is FileName Setter
// 文件名：注意文件名请保证和上传的文件一直
func (r *TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// GetFileName FileName Getter
func (r TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) GetFileName() string {
	return r._fileName
}

// SetFileBytes is FileBytes Setter
// 文件
func (r *TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) SetFileBytes(_fileBytes *model.File) error {
	r._fileBytes = _fileBytes
	r.Set("file_bytes", _fileBytes)
	return nil
}

// GetFileBytes FileBytes Getter
func (r TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) GetFileBytes() *model.File {
	return r._fileBytes
}

// SetBizOrderId is BizOrderId Setter
// 订单id
func (r *TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) SetBizOrderId(_bizOrderId int64) error {
	r._bizOrderId = _bizOrderId
	r.Set("biz_order_id", _bizOrderId)
	return nil
}

// GetBizOrderId BizOrderId Getter
func (r TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) GetBizOrderId() int64 {
	return r._bizOrderId
}

var poolTaobaoAlitripTravelNormalvisaUploadfileAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripTravelNormalvisaUploadfileRequest()
	},
}

// GetTaobaoAlitripTravelNormalvisaUploadfileRequest 从 sync.Pool 获取 TaobaoAlitripTravelNormalvisaUploadfileAPIRequest
func GetTaobaoAlitripTravelNormalvisaUploadfileAPIRequest() *TaobaoAlitripTravelNormalvisaUploadfileAPIRequest {
	return poolTaobaoAlitripTravelNormalvisaUploadfileAPIRequest.Get().(*TaobaoAlitripTravelNormalvisaUploadfileAPIRequest)
}

// ReleaseTaobaoAlitripTravelNormalvisaUploadfileAPIRequest 将 TaobaoAlitripTravelNormalvisaUploadfileAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripTravelNormalvisaUploadfileAPIRequest(v *TaobaoAlitripTravelNormalvisaUploadfileAPIRequest) {
	v.Reset()
	poolTaobaoAlitripTravelNormalvisaUploadfileAPIRequest.Put(v)
}
