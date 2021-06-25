package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
tv支付申请设备授权 APIRequest
taobao.tvpay.auth.apply

为用户在指定设备上申请支付授权
*/
type TaobaoTvpayAuthApplyRequest struct {
    model.Params

    // 设备id
    deviceId   string 

    // 请求来源
    from   string 

    // 场景
    bizScene   string 

    // 商品名称
    itemName   string 

    // 授权类型
    operateType   string 

    // 外部订单号
    outApproveId   string 

    // 金额
    totalFee   string 

}

func NewTaobaoTvpayAuthApplyRequest() *TaobaoTvpayAuthApplyRequest{
    return &TaobaoTvpayAuthApplyRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTvpayAuthApplyRequest) GetApiMethodName() string {
    return "taobao.tvpay.auth.apply"
}

func (r TaobaoTvpayAuthApplyRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTvpayAuthApplyRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r TaobaoTvpayAuthApplyRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *TaobaoTvpayAuthApplyRequest) SetFrom(from string) error {
    r.from = from
    r.Set("from", from)
    return nil
}

func (r TaobaoTvpayAuthApplyRequest) GetFrom() string {
    return r.from
}

func (r *TaobaoTvpayAuthApplyRequest) SetBizScene(bizScene string) error {
    r.bizScene = bizScene
    r.Set("biz_scene", bizScene)
    return nil
}

func (r TaobaoTvpayAuthApplyRequest) GetBizScene() string {
    return r.bizScene
}

func (r *TaobaoTvpayAuthApplyRequest) SetItemName(itemName string) error {
    r.itemName = itemName
    r.Set("item_name", itemName)
    return nil
}

func (r TaobaoTvpayAuthApplyRequest) GetItemName() string {
    return r.itemName
}

func (r *TaobaoTvpayAuthApplyRequest) SetOperateType(operateType string) error {
    r.operateType = operateType
    r.Set("operate_type", operateType)
    return nil
}

func (r TaobaoTvpayAuthApplyRequest) GetOperateType() string {
    return r.operateType
}

func (r *TaobaoTvpayAuthApplyRequest) SetOutApproveId(outApproveId string) error {
    r.outApproveId = outApproveId
    r.Set("out_approve_id", outApproveId)
    return nil
}

func (r TaobaoTvpayAuthApplyRequest) GetOutApproveId() string {
    return r.outApproveId
}

func (r *TaobaoTvpayAuthApplyRequest) SetTotalFee(totalFee string) error {
    r.totalFee = totalFee
    r.Set("total_fee", totalFee)
    return nil
}

func (r TaobaoTvpayAuthApplyRequest) GetTotalFee() string {
    return r.totalFee
}

