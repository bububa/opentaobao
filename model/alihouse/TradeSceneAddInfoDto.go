package alihouse

import (
	"sync"
)

// TradeSceneAddInfoDto 结构体
type TradeSceneAddInfoDto struct {
	// 外部场景ID
	OuterTradeSceneConfigId string `json:"outer_trade_scene_config_id,omitempty" xml:"outer_trade_scene_config_id,omitempty"`
	// 场景编码
	SceneCode int64 `json:"scene_code,omitempty" xml:"scene_code,omitempty"`
	// 经营主体ID
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
	// 结算账户ID	否
	BankId int64 `json:"bank_id,omitempty" xml:"bank_id,omitempty"`
	// 电子签章ID
	SignatureId int64 `json:"signature_id,omitempty" xml:"signature_id,omitempty"`
	// 协议模版ID
	AgreementId int64 `json:"agreement_id,omitempty" xml:"agreement_id,omitempty"`
	// 代收主体id�
	InsteadMerchantOpenId int64 `json:"instead_merchant_open_id,omitempty" xml:"instead_merchant_open_id,omitempty"`
	// 代收结算账户id�
	InsteadBankId int64 `json:"instead_bank_id,omitempty" xml:"instead_bank_id,omitempty"`
	// 收费模式 1-直收 2-代收�
	CollectionType int64 `json:"collection_type,omitempty" xml:"collection_type,omitempty"`
	// 是否测试 1-是 0-否
	IsTest int64 `json:"is_test,omitempty" xml:"is_test,omitempty"`
}

var poolTradeSceneAddInfoDto = sync.Pool{
	New: func() any {
		return new(TradeSceneAddInfoDto)
	},
}

// GetTradeSceneAddInfoDto() 从对象池中获取TradeSceneAddInfoDto
func GetTradeSceneAddInfoDto() *TradeSceneAddInfoDto {
	return poolTradeSceneAddInfoDto.Get().(*TradeSceneAddInfoDto)
}

// ReleaseTradeSceneAddInfoDto 释放TradeSceneAddInfoDto
func ReleaseTradeSceneAddInfoDto(v *TradeSceneAddInfoDto) {
	v.OuterTradeSceneConfigId = ""
	v.SceneCode = 0
	v.MerchantOpenId = 0
	v.BankId = 0
	v.SignatureId = 0
	v.AgreementId = 0
	v.InsteadMerchantOpenId = 0
	v.InsteadBankId = 0
	v.CollectionType = 0
	v.IsTest = 0
	poolTradeSceneAddInfoDto.Put(v)
}
