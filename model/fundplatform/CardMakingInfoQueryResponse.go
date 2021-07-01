package fundplatform

// CardMakingInfoQueryResponse 结构体
type CardMakingInfoQueryResponse struct {
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 卡信息列表。cardNo卡号,cardPassword卡密,qrCode二维码链接,shortQrCode短连接二维码,templateNo卡商品编码,status卡状态
	CardMakingInfos []string `json:"card_making_infos,omitempty" xml:"card_making_infos>string,omitempty"`
}
