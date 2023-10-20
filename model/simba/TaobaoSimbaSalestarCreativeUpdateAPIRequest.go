package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestarcreativeupdateAPIRequest 销量明星更新创意相关接口 API请求
// taobao.simba.salestar.creative.update
//
// 更新一个创意的信息，可以设置创意标题、创意图片
type TaobaosimbasalestarcreativeupdateAPIRequest struct {
	model.Params
	// 创意标题，最多20个汉字
	_title string
	// 创意图片地址，必须是推广组对应商品的图片之一
	_imgUrl string
	// 推广组Id
	_adgroupId int64
	// 创意Id
	_creativeId int64
	// 如果用户开通了创意本地上传图片功能的，可以使用该用户图片空间的图片来修改创意，pictureId为图片空间中图片的pictureId，img_url为图片空间中图片链接地址，如果是使用的主图或副图修改创意，则pictureId必须为空
	_pictureId int64
}

// NewTaobaosimbasalestarcreativeupdateRequest 初始化TaobaosimbasalestarcreativeupdateAPIRequest对象
func NewTaobaosimbasalestarcreativeupdateRequest() *TaobaosimbasalestarcreativeupdateAPIRequest {
	return &TaobaosimbasalestarcreativeupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbasalestarcreativeupdateAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.creative.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbasalestarcreativeupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbasalestarcreativeupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitle is Title Setter
// 创意标题，最多20个汉字
func (r *TaobaosimbasalestarcreativeupdateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaosimbasalestarcreativeupdateAPIRequest) GetTitle() string {
	return r._title
}

// SetImgUrl is ImgUrl Setter
// 创意图片地址，必须是推广组对应商品的图片之一
func (r *TaobaosimbasalestarcreativeupdateAPIRequest) SetImgUrl(_imgUrl string) error {
	r._imgUrl = _imgUrl
	r.Set("img_url", _imgUrl)
	return nil
}

// GetImgUrl ImgUrl Getter
func (r TaobaosimbasalestarcreativeupdateAPIRequest) GetImgUrl() string {
	return r._imgUrl
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosimbasalestarcreativeupdateAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbasalestarcreativeupdateAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

// SetCreativeId is CreativeId Setter
// 创意Id
func (r *TaobaosimbasalestarcreativeupdateAPIRequest) SetCreativeId(_creativeId int64) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaosimbasalestarcreativeupdateAPIRequest) GetCreativeId() int64 {
	return r._creativeId
}

// SetPictureId is PictureId Setter
// 如果用户开通了创意本地上传图片功能的，可以使用该用户图片空间的图片来修改创意，pictureId为图片空间中图片的pictureId，img_url为图片空间中图片链接地址，如果是使用的主图或副图修改创意，则pictureId必须为空
func (r *TaobaosimbasalestarcreativeupdateAPIRequest) SetPictureId(_pictureId int64) error {
	r._pictureId = _pictureId
	r.Set("picture_id", _pictureId)
	return nil
}

// GetPictureId PictureId Getter
func (r TaobaosimbasalestarcreativeupdateAPIRequest) GetPictureId() int64 {
	return r._pictureId
}
