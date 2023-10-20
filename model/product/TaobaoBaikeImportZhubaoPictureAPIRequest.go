package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaobaikeimportzhubaopictureAPIRequest 百科图片数据导入 API请求
// taobao.baike.import.zhubao.picture
//
// 用于接入外部--图片--录入到商品百科中
type TaobaobaikeimportzhubaopictureAPIRequest struct {
	model.Params
	// 图片二进制数据
	_picture *model.File
}

// NewTaobaobaikeimportzhubaopictureRequest 初始化TaobaobaikeimportzhubaopictureAPIRequest对象
func NewTaobaobaikeimportzhubaopictureRequest() *TaobaobaikeimportzhubaopictureAPIRequest {
	return &TaobaobaikeimportzhubaopictureAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaobaikeimportzhubaopictureAPIRequest) GetApiMethodName() string {
	return "taobao.baike.import.zhubao.picture"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaobaikeimportzhubaopictureAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaobaikeimportzhubaopictureAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPicture is Picture Setter
// 图片二进制数据
func (r *TaobaobaikeimportzhubaopictureAPIRequest) SetPicture(_picture *model.File) error {
	r._picture = _picture
	r.Set("picture", _picture)
	return nil
}

// GetPicture Picture Getter
func (r TaobaobaikeimportzhubaopictureAPIRequest) GetPicture() *model.File {
	return r._picture
}
