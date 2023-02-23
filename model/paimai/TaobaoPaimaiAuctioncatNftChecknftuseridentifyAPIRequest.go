package paimai

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest 根据用户数字id和身份证号校验该用户是否已实名认证成功 API请求
// taobao.paimai.auctioncat.nft.checknftuseridentify
//
// 根据用户数字id和身份证号校验该用户是否已实名认证成功
type TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest struct {
	model.Params
	// 用户数字id
	_thirdId string
	// 身份证号的hashCode
	_idNumber string
}

// NewTaobaoPaimaiAuctioncatNftChecknftuseridentifyRequest 初始化TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest对象
func NewTaobaoPaimaiAuctioncatNftChecknftuseridentifyRequest() *TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest {
	return &TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest) GetApiMethodName() string {
	return "taobao.paimai.auctioncat.nft.checknftuseridentify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetThirdId is ThirdId Setter
// 用户数字id
func (r *TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest) SetThirdId(_thirdId string) error {
	r._thirdId = _thirdId
	r.Set("third_id", _thirdId)
	return nil
}

// GetThirdId ThirdId Getter
func (r TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest) GetThirdId() string {
	return r._thirdId
}

// SetIdNumber is IdNumber Setter
// 身份证号的hashCode
func (r *TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest) SetIdNumber(_idNumber string) error {
	r._idNumber = _idNumber
	r.Set("id_number", _idNumber)
	return nil
}

// GetIdNumber IdNumber Getter
func (r TaobaoPaimaiAuctioncatNftChecknftuseridentifyAPIRequest) GetIdNumber() string {
	return r._idNumber
}
