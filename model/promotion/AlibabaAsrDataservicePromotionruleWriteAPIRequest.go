package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAsrDataservicePromotionruleWriteAPIRequest 业务优惠规则写入 API请求
// alibaba.asr.dataservice.promotionrule.write
//
// 星巴克优惠规则写入
type AlibabaAsrDataservicePromotionruleWriteAPIRequest struct {
	model.Params
	// 入参对象
	_poskeyPromotionRuleDto *PosKeyPromotionRuleDto
}

// NewAlibabaAsrDataservicePromotionruleWriteRequest 初始化AlibabaAsrDataservicePromotionruleWriteAPIRequest对象
func NewAlibabaAsrDataservicePromotionruleWriteRequest() *AlibabaAsrDataservicePromotionruleWriteAPIRequest {
	return &AlibabaAsrDataservicePromotionruleWriteAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAsrDataservicePromotionruleWriteAPIRequest) Reset() {
	r._poskeyPromotionRuleDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAsrDataservicePromotionruleWriteAPIRequest) GetApiMethodName() string {
	return "alibaba.asr.dataservice.promotionrule.write"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAsrDataservicePromotionruleWriteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAsrDataservicePromotionruleWriteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPoskeyPromotionRuleDto is PoskeyPromotionRuleDto Setter
// 入参对象
func (r *AlibabaAsrDataservicePromotionruleWriteAPIRequest) SetPoskeyPromotionRuleDto(_poskeyPromotionRuleDto *PosKeyPromotionRuleDto) error {
	r._poskeyPromotionRuleDto = _poskeyPromotionRuleDto
	r.Set("poskey_promotion_rule_dto", _poskeyPromotionRuleDto)
	return nil
}

// GetPoskeyPromotionRuleDto PoskeyPromotionRuleDto Getter
func (r AlibabaAsrDataservicePromotionruleWriteAPIRequest) GetPoskeyPromotionRuleDto() *PosKeyPromotionRuleDto {
	return r._poskeyPromotionRuleDto
}

var poolAlibabaAsrDataservicePromotionruleWriteAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAsrDataservicePromotionruleWriteRequest()
	},
}

// GetAlibabaAsrDataservicePromotionruleWriteRequest 从 sync.Pool 获取 AlibabaAsrDataservicePromotionruleWriteAPIRequest
func GetAlibabaAsrDataservicePromotionruleWriteAPIRequest() *AlibabaAsrDataservicePromotionruleWriteAPIRequest {
	return poolAlibabaAsrDataservicePromotionruleWriteAPIRequest.Get().(*AlibabaAsrDataservicePromotionruleWriteAPIRequest)
}

// ReleaseAlibabaAsrDataservicePromotionruleWriteAPIRequest 将 AlibabaAsrDataservicePromotionruleWriteAPIRequest 放入 sync.Pool
func ReleaseAlibabaAsrDataservicePromotionruleWriteAPIRequest(v *AlibabaAsrDataservicePromotionruleWriteAPIRequest) {
	v.Reset()
	poolAlibabaAsrDataservicePromotionruleWriteAPIRequest.Put(v)
}
