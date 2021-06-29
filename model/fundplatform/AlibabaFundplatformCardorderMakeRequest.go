package fundplatform

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知制卡商制卡 API请求
alibaba.fundplatform.cardorder.make

该接口由内部定义，外部制卡商实现。当需要制卡商进行制卡操作时，将会调用该接口。
*/
type AlibabaFundplatformCardorderMakeRequest struct {
    model.Params
    // 物流信息
    _logisticsInfo   *AlibabaFundplatformCardorderMakeStruct
    // 卡模板信息列表
    _cardProductInfos   []AlibabaFundplatformCardorderMakeStruct
    // 子制卡单ID
    _cardOrderId   int64
    // 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
    _ownSign   string
}

// 初始化AlibabaFundplatformCardorderMakeRequest对象
func NewAlibabaFundplatformCardorderMakeRequest() *AlibabaFundplatformCardorderMakeRequest{
    return &AlibabaFundplatformCardorderMakeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderMakeRequest) GetApiMethodName() string {
    return "alibaba.fundplatform.cardorder.make"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderMakeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LogisticsInfo Setter
// 物流信息
func (r *AlibabaFundplatformCardorderMakeRequest) SetLogisticsInfo(_logisticsInfo *AlibabaFundplatformCardorderMakeStruct) error {
    r._logisticsInfo = _logisticsInfo
    r.Set("logistics_info", _logisticsInfo)
    return nil
}

// LogisticsInfo Getter
func (r AlibabaFundplatformCardorderMakeRequest) GetLogisticsInfo() *AlibabaFundplatformCardorderMakeStruct {
    return r._logisticsInfo
}
// CardProductInfos Setter
// 卡模板信息列表
func (r *AlibabaFundplatformCardorderMakeRequest) SetCardProductInfos(_cardProductInfos []AlibabaFundplatformCardorderMakeStruct) error {
    r._cardProductInfos = _cardProductInfos
    r.Set("card_product_infos", _cardProductInfos)
    return nil
}

// CardProductInfos Getter
func (r AlibabaFundplatformCardorderMakeRequest) GetCardProductInfos() []AlibabaFundplatformCardorderMakeStruct {
    return r._cardProductInfos
}
// CardOrderId Setter
// 子制卡单ID
func (r *AlibabaFundplatformCardorderMakeRequest) SetCardOrderId(_cardOrderId int64) error {
    r._cardOrderId = _cardOrderId
    r.Set("card_order_id", _cardOrderId)
    return nil
}

// CardOrderId Getter
func (r AlibabaFundplatformCardorderMakeRequest) GetCardOrderId() int64 {
    return r._cardOrderId
}
// OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabaFundplatformCardorderMakeRequest) SetOwnSign(_ownSign string) error {
    r._ownSign = _ownSign
    r.Set("own_sign", _ownSign)
    return nil
}

// OwnSign Getter
func (r AlibabaFundplatformCardorderMakeRequest) GetOwnSign() string {
    return r._ownSign
}
