package user

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取用户openId APIRequest
taobao.koubei.tribe.open.user.query

口碑综合体通过手机号码获取加密后的用户openId
*/
type TaobaoKoubeiTribeOpenUserQueryRequest struct {
    model.Params

    // 验证码
    verifyCode   string 

    // 手机号
    phone   string 

    // 数据集id
    dataSetId   string 

}

func NewTaobaoKoubeiTribeOpenUserQueryRequest() *TaobaoKoubeiTribeOpenUserQueryRequest{
    return &TaobaoKoubeiTribeOpenUserQueryRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoKoubeiTribeOpenUserQueryRequest) GetApiMethodName() string {
    return "taobao.koubei.tribe.open.user.query"
}

func (r TaobaoKoubeiTribeOpenUserQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoKoubeiTribeOpenUserQueryRequest) SetVerifyCode(verifyCode string) error {
    r.verifyCode = verifyCode
    r.Set("verify_code", verifyCode)
    return nil
}

func (r TaobaoKoubeiTribeOpenUserQueryRequest) GetVerifyCode() string {
    return r.verifyCode
}

func (r *TaobaoKoubeiTribeOpenUserQueryRequest) SetPhone(phone string) error {
    r.phone = phone
    r.Set("phone", phone)
    return nil
}

func (r TaobaoKoubeiTribeOpenUserQueryRequest) GetPhone() string {
    return r.phone
}

func (r *TaobaoKoubeiTribeOpenUserQueryRequest) SetDataSetId(dataSetId string) error {
    r.dataSetId = dataSetId
    r.Set("data_set_id", dataSetId)
    return nil
}

func (r TaobaoKoubeiTribeOpenUserQueryRequest) GetDataSetId() string {
    return r.dataSetId
}

