package simba

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCreativeAddAPIRequest 增加创意 API请求
// taobao.simba.creative.add
//
// 创建一个创意
type TaobaoSimbaCreativeAddAPIRequest struct {
	model.Params
	// 主人昵称
	_nick string
	// 创意标题，最多30个汉字
	_title string
	// 创意图片地址，必须是推广组对应商品的图片之一
	_imgUrl string
	// 广审批准文号
	_adExaminationCode string
	// 推广组Id
	_adgroupId int64
}

// NewTaobaoSimbaCreativeAddRequest 初始化TaobaoSimbaCreativeAddAPIRequest对象
func NewTaobaoSimbaCreativeAddRequest() *TaobaoSimbaCreativeAddAPIRequest {
	return &TaobaoSimbaCreativeAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSimbaCreativeAddAPIRequest) GetApiMethodName() string {
	return "taobao.simba.creative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSimbaCreativeAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetNick is Nick Setter
// 主人昵称
func (r *TaobaoSimbaCreativeAddAPIRequest) SetNick(_nick string) error {
	r._nick = _nick
	r.Set("nick", _nick)
	return nil
}

// GetNick Nick Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetNick() string {
	return r._nick
}

// SetTitle is Title Setter
// 创意标题，最多30个汉字
func (r *TaobaoSimbaCreativeAddAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetTitle() string {
	return r._title
}

// SetImgUrl is ImgUrl Setter
// 创意图片地址，必须是推广组对应商品的图片之一
func (r *TaobaoSimbaCreativeAddAPIRequest) SetImgUrl(_imgUrl string) error {
	r._imgUrl = _imgUrl
	r.Set("img_url", _imgUrl)
	return nil
}

// GetImgUrl ImgUrl Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetImgUrl() string {
	return r._imgUrl
}

// SetAdExaminationCode is AdExaminationCode Setter
// 广审批准文号
func (r *TaobaoSimbaCreativeAddAPIRequest) SetAdExaminationCode(_adExaminationCode string) error {
	r._adExaminationCode = _adExaminationCode
	r.Set("ad_examination_code", _adExaminationCode)
	return nil
}

// GetAdExaminationCode AdExaminationCode Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetAdExaminationCode() string {
	return r._adExaminationCode
}

// SetAdgroupId is AdgroupId Setter
// 推广组Id
func (r *TaobaoSimbaCreativeAddAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoSimbaCreativeAddAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
