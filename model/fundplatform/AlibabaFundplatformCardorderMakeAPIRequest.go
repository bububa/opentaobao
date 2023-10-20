package fundplatform

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFundplatformCardorderMakeAPIRequest 通知制卡商制卡 API请求
// alibaba.fundplatform.cardorder.make
//
// 该接口由内部定义，外部制卡商实现。当需要制卡商进行制卡操作时，将会调用该接口。
type AlibabaFundplatformCardorderMakeAPIRequest struct {
	model.Params
	// 卡模板信息列表
	_cardProductInfos []AlibabaFundplatformCardorderMakeStruct
	// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
	_ownSign string
	// 物流信息
	_logisticsInfo *AlibabaFundplatformCardorderMakeStruct
	// 子制卡单ID
	_cardOrderId int64
}

// NewAlibabaFundplatformCardorderMakeRequest 初始化AlibabaFundplatformCardorderMakeAPIRequest对象
func NewAlibabaFundplatformCardorderMakeRequest() *AlibabaFundplatformCardorderMakeAPIRequest {
	return &AlibabaFundplatformCardorderMakeAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaFundplatformCardorderMakeAPIRequest) Reset() {
	r._cardProductInfos = r._cardProductInfos[:0]
	r._ownSign = ""
	r._logisticsInfo = nil
	r._cardOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaFundplatformCardorderMakeAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.cardorder.make"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaFundplatformCardorderMakeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaFundplatformCardorderMakeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCardProductInfos is CardProductInfos Setter
// 卡模板信息列表
func (r *AlibabaFundplatformCardorderMakeAPIRequest) SetCardProductInfos(_cardProductInfos []AlibabaFundplatformCardorderMakeStruct) error {
	r._cardProductInfos = _cardProductInfos
	r.Set("card_product_infos", _cardProductInfos)
	return nil
}

// GetCardProductInfos CardProductInfos Getter
func (r AlibabaFundplatformCardorderMakeAPIRequest) GetCardProductInfos() []AlibabaFundplatformCardorderMakeStruct {
	return r._cardProductInfos
}

// SetOwnSign is OwnSign Setter
// 环境变量值，该字段为枚举值：daily（日常），pre（预发），online（线上）
func (r *AlibabaFundplatformCardorderMakeAPIRequest) SetOwnSign(_ownSign string) error {
	r._ownSign = _ownSign
	r.Set("own_sign", _ownSign)
	return nil
}

// GetOwnSign OwnSign Getter
func (r AlibabaFundplatformCardorderMakeAPIRequest) GetOwnSign() string {
	return r._ownSign
}

// SetLogisticsInfo is LogisticsInfo Setter
// 物流信息
func (r *AlibabaFundplatformCardorderMakeAPIRequest) SetLogisticsInfo(_logisticsInfo *AlibabaFundplatformCardorderMakeStruct) error {
	r._logisticsInfo = _logisticsInfo
	r.Set("logistics_info", _logisticsInfo)
	return nil
}

// GetLogisticsInfo LogisticsInfo Getter
func (r AlibabaFundplatformCardorderMakeAPIRequest) GetLogisticsInfo() *AlibabaFundplatformCardorderMakeStruct {
	return r._logisticsInfo
}

// SetCardOrderId is CardOrderId Setter
// 子制卡单ID
func (r *AlibabaFundplatformCardorderMakeAPIRequest) SetCardOrderId(_cardOrderId int64) error {
	r._cardOrderId = _cardOrderId
	r.Set("card_order_id", _cardOrderId)
	return nil
}

// GetCardOrderId CardOrderId Getter
func (r AlibabaFundplatformCardorderMakeAPIRequest) GetCardOrderId() int64 {
	return r._cardOrderId
}

var poolAlibabaFundplatformCardorderMakeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaFundplatformCardorderMakeRequest()
	},
}

// GetAlibabaFundplatformCardorderMakeRequest 从 sync.Pool 获取 AlibabaFundplatformCardorderMakeAPIRequest
func GetAlibabaFundplatformCardorderMakeAPIRequest() *AlibabaFundplatformCardorderMakeAPIRequest {
	return poolAlibabaFundplatformCardorderMakeAPIRequest.Get().(*AlibabaFundplatformCardorderMakeAPIRequest)
}

// ReleaseAlibabaFundplatformCardorderMakeAPIRequest 将 AlibabaFundplatformCardorderMakeAPIRequest 放入 sync.Pool
func ReleaseAlibabaFundplatformCardorderMakeAPIRequest(v *AlibabaFundplatformCardorderMakeAPIRequest) {
	v.Reset()
	poolAlibabaFundplatformCardorderMakeAPIRequest.Put(v)
}
