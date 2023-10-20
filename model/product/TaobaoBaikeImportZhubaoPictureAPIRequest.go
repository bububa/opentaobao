package product

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBaikeImportZhubaoPictureAPIRequest 百科图片数据导入 API请求
// taobao.baike.import.zhubao.picture
//
// 用于接入外部--图片--录入到商品百科中
type TaobaoBaikeImportZhubaoPictureAPIRequest struct {
	model.Params
	// 图片二进制数据
	_picture *model.File
}

// NewTaobaoBaikeImportZhubaoPictureRequest 初始化TaobaoBaikeImportZhubaoPictureAPIRequest对象
func NewTaobaoBaikeImportZhubaoPictureRequest() *TaobaoBaikeImportZhubaoPictureAPIRequest {
	return &TaobaoBaikeImportZhubaoPictureAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoBaikeImportZhubaoPictureAPIRequest) Reset() {
	r._picture = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaikeImportZhubaoPictureAPIRequest) GetApiMethodName() string {
	return "taobao.baike.import.zhubao.picture"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaikeImportZhubaoPictureAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoBaikeImportZhubaoPictureAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPicture is Picture Setter
// 图片二进制数据
func (r *TaobaoBaikeImportZhubaoPictureAPIRequest) SetPicture(_picture *model.File) error {
	r._picture = _picture
	r.Set("picture", _picture)
	return nil
}

// GetPicture Picture Getter
func (r TaobaoBaikeImportZhubaoPictureAPIRequest) GetPicture() *model.File {
	return r._picture
}

var poolTaobaoBaikeImportZhubaoPictureAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoBaikeImportZhubaoPictureRequest()
	},
}

// GetTaobaoBaikeImportZhubaoPictureRequest 从 sync.Pool 获取 TaobaoBaikeImportZhubaoPictureAPIRequest
func GetTaobaoBaikeImportZhubaoPictureAPIRequest() *TaobaoBaikeImportZhubaoPictureAPIRequest {
	return poolTaobaoBaikeImportZhubaoPictureAPIRequest.Get().(*TaobaoBaikeImportZhubaoPictureAPIRequest)
}

// ReleaseTaobaoBaikeImportZhubaoPictureAPIRequest 将 TaobaoBaikeImportZhubaoPictureAPIRequest 放入 sync.Pool
func ReleaseTaobaoBaikeImportZhubaoPictureAPIRequest(v *TaobaoBaikeImportZhubaoPictureAPIRequest) {
	v.Reset()
	poolTaobaoBaikeImportZhubaoPictureAPIRequest.Put(v)
}
