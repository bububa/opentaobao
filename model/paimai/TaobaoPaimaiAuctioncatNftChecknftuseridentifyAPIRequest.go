package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest 根据用户数字id和身份证号校验该用户是否已实名认证成功 API请求
// taobao.paimai.auctioncat.nft.checknftuseridentify
//
// 根据用户数字id和身份证号校验该用户是否已实名认证成功
type TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest struct {
	model.Params
	// 用户数字id
	_thirdId string
	// 身份证号的hashCode
	_idNumber string
}

// NewTaobaopaimaiauctioncatnftchecknftuseridentifyRequest 初始化TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest对象
func NewTaobaopaimaiauctioncatnftchecknftuseridentifyRequest() *TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest {
	return &TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest) GetApiMethodName() string {
	return "taobao.paimai.auctioncat.nft.checknftuseridentify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetThirdId is ThirdId Setter
// 用户数字id
func (r *TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest) SetThirdId(_thirdId string) error {
	r._thirdId = _thirdId
	r.Set("third_id", _thirdId)
	return nil
}

// GetThirdId ThirdId Getter
func (r TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest) GetThirdId() string {
	return r._thirdId
}

// SetIdNumber is IdNumber Setter
// 身份证号的hashCode
func (r *TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r TaobaopaimaiauctioncatnftchecknftuseridentifyAPIRequest) GetIdNumber() string {
	return r._idNumber
}
