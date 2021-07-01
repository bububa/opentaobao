package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaikeImportZhubaoPictureAPIRequest
百科图片数据导入 API请求
taobao.baike.import.zhubao.picture

用于接入外部--图片--录入到商品百科中 */
type TaobaoBaikeImportZhubaoPictureAPIRequest struct {
	model.Params
	// 图片二进制数据
	_picture *model.File
}

// NewTaobaoBaikeImportZhubaoPictureRequest 初始化TaobaoBaikeImportZhubaoPictureAPIRequest对象
func NewTaobaoBaikeImportZhubaoPictureRequest() *TaobaoBaikeImportZhubaoPictureAPIRequest {
	return &TaobaoBaikeImportZhubaoPictureAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoBaikeImportZhubaoPictureAPIRequest) GetApiMethodName() string {
	return "taobao.baike.import.zhubao.picture"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoBaikeImportZhubaoPictureAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Picture Setter
// 图片二进制数据
func (r *TaobaoBaikeImportZhubaoPictureAPIRequest) SetPicture(_picture *model.File) error {
	r._picture = _picture
	r.Set("picture", _picture)
	return nil
}

// Get Picture Getter
func (r TaobaoBaikeImportZhubaoPictureAPIRequest) GetPicture() *model.File {
	return r._picture
}
