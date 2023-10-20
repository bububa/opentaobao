package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformcardordermakeAPIRequest 通知制卡商制卡 API请求
// alibaba.fundplatform.cardorder.make
//
// 该接口由内部定义，外部制卡商实现。当需要制卡商进行制卡操作时，将会调用该接口。
type AlibabafundplatformcardordermakeAPIRequest struct {
	model.Params
	// 卡模板信息列表
	_cardProductInfos []AlibabafundplatformcardordermakeStruct
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	_ownSign string
	// 物流信息
	_logisticsInfo *AlibabafundplatformcardordermakeStruct
	// 子制卡单ID
	_cardOrderId int64
}

// NewAlibabafundplatformcardordermakeRequest 初始化AlibabafundplatformcardordermakeAPIRequest对象
func NewAlibabafundplatformcardordermakeRequest() *AlibabafundplatformcardordermakeAPIRequest {
	return &AlibabafundplatformcardordermakeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformcardordermakeAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.make"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformcardordermakeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformcardordermakeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCardProductInfos is CardProductInfos Setter
// 卡模板信息列表
func (r *AlibabafundplatformcardordermakeAPIRequest) SetCardProductInfos(_cardProductInfos []AlibabafundplatformcardordermakeStruct) error {
	r._cardProductInfos = _cardProductInfos
	r.Set("card_product_infos", _cardProductInfos)
	return nil
}

// GetCardProductInfos CardProductInfos Getter
func (r AlibabafundplatformcardordermakeAPIRequest) GetCardProductInfos() []AlibabafundplatformcardordermakeStruct {
	return r._cardProductInfos
}

// SetOwnSign is OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabafundplatformcardordermakeAPIRequest) SetOwnSign(_ownSign string) error {
	r._ownSign = _ownSign
	r.Set("own_sign", _ownSign)
	return nil
}

// GetOwnSign OwnSign Getter
func (r AlibabafundplatformcardordermakeAPIRequest) GetOwnSign() string {
	return r._ownSign
}

// SetLogisticsInfo is LogisticsInfo Setter
// 物流信息
func (r *AlibabafundplatformcardordermakeAPIRequest) SetLogisticsInfo(_logisticsInfo *AlibabafundplatformcardordermakeStruct) error {
	r._logisticsInfo = _logisticsInfo
	r.Set("logistics_info", _logisticsInfo)
	return nil
}

// GetLogisticsInfo LogisticsInfo Getter
func (r AlibabafundplatformcardordermakeAPIRequest) GetLogisticsInfo() *AlibabafundplatformcardordermakeStruct {
	return r._logisticsInfo
}

// SetCardOrderId is CardOrderId Setter
// 子制卡单ID
func (r *AlibabafundplatformcardordermakeAPIRequest) SetCardOrderId(_cardOrderId int64) error {
	r._cardOrderId = _cardOrderId
	r.Set("card_order_id", _cardOrderId)
	return nil
}

// GetCardOrderId CardOrderId Getter
func (r AlibabafundplatformcardordermakeAPIRequest) GetCardOrderId() int64 {
	return r._cardOrderId
}
