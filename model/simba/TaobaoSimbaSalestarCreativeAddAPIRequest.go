package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaosimbasalestarcreativeaddAPIRequest （新）新建创意 API请求
// taobao.simba.salestar.creative.add
//
// 创建一个创意
type TaobaosimbasalestarcreativeaddAPIRequest struct {
	model.Params
	// 创意标题，最多20个汉字
	_title string
	// 创意图片地址，必须是推广组对应商品的图片之一
	_imgUrl string
	// 主人昵称
	_nick string
	// 创意广审编号
	_adExaminationCode string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaosimbasalestarcreativeaddRequest 初始化TaobaosimbasalestarcreativeaddAPIRequest对象
func NewTaobaosimbasalestarcreativeaddRequest() *TaobaosimbasalestarcreativeaddAPIRequest {
	return &TaobaosimbasalestarcreativeaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaosimbasalestarcreativeaddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.salestar.creative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaosimbasalestarcreativeaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaosimbasalestarcreativeaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitle is Title Setter
// 创意标题，最多20个汉字
func (r *TaobaosimbasalestarcreativeaddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaosimbasalestarcreativeaddAPIRequest) GetTitle() string {
	return r._title
}

// SetImgUrl is ImgUrl Setter
// 创意图片地址，必须是推广组对应商品的图片之一
func (r *TaobaosimbasalestarcreativeaddAPIRequest) SetImgUrl(_imgUrl string) error {
	r._imgUrl = _imgUrl
	r.Set("img_url", _imgUrl)
	return nil
}

// GetImgUrl ImgUrl Getter
func (r TaobaosimbasalestarcreativeaddAPIRequest) GetImgUrl() string {
	return r._imgUrl
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaosimbasalestarcreativeaddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaosimbasalestarcreativeaddAPIRequest) GetNick() string {
	return r._nick
}

// SetAdExaminationCode is AdExaminationCode Setter
// 创意广审编号
func (r *TaobaosimbasalestarcreativeaddAPIRequest) SetAdExaminationCode(_adExaminationCode string) error {
	r._adExaminationCode = _adExaminationCode
	r.Set("ad_examination_code", _adExaminationCode)
	return nil
}

// GetAdExaminationCode AdExaminationCode Getter
func (r TaobaosimbasalestarcreativeaddAPIRequest) GetAdExaminationCode() string {
	return r._adExaminationCode
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaosimbasalestarcreativeaddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaosimbasalestarcreativeaddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
