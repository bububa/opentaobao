package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest 口碑综合体手机号获取验证码 API请求
// taobao.koubei.tribe.open.verify.code.apply
//
// 口碑综合体通过手机号获取验证码对外开放接口
type TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest struct {
	model.Params
	// 数据集id
	_dataSetId string
	// 手机号
	_phone string
}

// NewTaobaoKoubeiTribeOpenVerifyCodeApplyRequest 初始化TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest对象
func NewTaobaoKoubeiTribeOpenVerifyCodeApplyRequest() *TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest {
	return &TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest) GetApiMethodName() string {
	return "taobao.koubei.tribe.open.verify.code.apply"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DataSetId Setter
// 数据集id
func (r *TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest) SetDataSetId(_dataSetId string) error {
	r._dataSetId = _dataSetId
	r.Set("data_set_id", _dataSetId)
	return nil
}

// Get DataSetId Getter
func (r TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest) GetDataSetId() string {
	return r._dataSetId
}

// Set is Phone Setter
// 手机号
func (r *TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest) SetPhone(_phone string) error {
	r._phone = _phone
	r.Set("phone", _phone)
	return nil
}

// Get Phone Getter
func (r TaobaoKoubeiTribeOpenVerifyCodeApplyAPIRequest) GetPhone() string {
	return r._phone
}
