package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
口碑综合体手机号获取验证码 API请求
taobao.koubei.tribe.open.verify.code.apply

口碑综合体通过手机号获取验证码对外开放接口
*/
type TaobaoKoubeiTribeOpenVerifyCodeApplyRequest struct {
    model.Params
    // 数据集id
    _dataSetId   string
    // 手机号
    _phone   string
}

// 初始化TaobaoKoubeiTribeOpenVerifyCodeApplyRequest对象
func NewTaobaoKoubeiTribeOpenVerifyCodeApplyRequest() *TaobaoKoubeiTribeOpenVerifyCodeApplyRequest{
    return &TaobaoKoubeiTribeOpenVerifyCodeApplyRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) GetApiMethodName() string {
    return "taobao.koubei.tribe.open.verify.code.apply"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DataSetId Setter
// 数据集id
func (r *TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) SetDataSetId(_dataSetId string) error {
    r._dataSetId = _dataSetId
    r.Set("data_set_id", _dataSetId)
    return nil
}

// DataSetId Getter
func (r TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) GetDataSetId() string {
    return r._dataSetId
}
// Phone Setter
// 手机号
func (r *TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) SetPhone(_phone string) error {
    r._phone = _phone
    r.Set("phone", _phone)
    return nil
}

// Phone Getter
func (r TaobaoKoubeiTribeOpenVerifyCodeApplyRequest) GetPhone() string {
    return r._phone
}
