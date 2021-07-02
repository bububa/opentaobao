package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiTribeOpenUserQueryAPIRequest 获取用户openId API请求
// taobao.koubei.tribe.open.user.query
//
// 口碑综合体通过手机号码获取加密后的用户openId
type TaobaoKoubeiTribeOpenUserQueryAPIRequest struct {
	model.Params
	// 验证码
	_verifyCode string
	// 手机号
	_phone string
	// 数据集id
	_dataSetId string
}

// NewTaobaoKoubeiTribeOpenUserQueryRequest 初始化TaobaoKoubeiTribeOpenUserQueryAPIRequest对象
func NewTaobaoKoubeiTribeOpenUserQueryRequest() *TaobaoKoubeiTribeOpenUserQueryAPIRequest {
	return &TaobaoKoubeiTribeOpenUserQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiTribeOpenUserQueryAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.tribe.open.user.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiTribeOpenUserQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetVerifyCode is VerifyCode Setter
// 验证码
func (r *TaobaoKoubeiTribeOpenUserQueryAPIRequest) SetVerifyCode(_verifyCode string) error {
	r._verifyCode = _verifyCode
	r.Set("verify_code", _verifyCode)
	return nil
}

// GetVerifyCode VerifyCode Getter
func (r TaobaoKoubeiTribeOpenUserQueryAPIRequest) GetVerifyCode() string {
	return r._verifyCode
}

// SetPhone is Phone Setter
// 手机号
func (r *TaobaoKoubeiTribeOpenUserQueryAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// GetPhone Phone Getter
func (r TaobaoKoubeiTribeOpenUserQueryAPIRequest) GetPhone() string {
	return r._phone
}

// SetDataSetId is DataSetId Setter
// 数据集id
func (r *TaobaoKoubeiTribeOpenUserQueryAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// GetDataSetId DataSetId Getter
func (r TaobaoKoubeiTribeOpenUserQueryAPIRequest) GetDataSetId() string {
	return r._dataSetId
}
