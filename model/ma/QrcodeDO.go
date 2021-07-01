package ma

// QrcodeDo 结构体
type QrcodeDo struct {
	// 二维码id
	QrcodeId int64 `json:"qrcode_id,omitempty" xml:"qrcode_id,omitempty"`
	// 二维码图片链接
	QrcodeUrl string `json:"qrcode_url,omitempty" xml:"qrcode_url,omitempty"`
	// 二维码扫码后访问的目的地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 二维码的短地址，每一个原始地址都会生成一个短地址
	ShortUrl string `json:"short_url,omitempty" xml:"short_url,omitempty"`
	// 二维码的矢量图下载地址，只有设置入参need_eps为true且style不为官方模板时，才返回值
	EpsUrl string `json:"eps_url,omitempty" xml:"eps_url,omitempty"`
	// 二维码所属渠道ID
	ChannelId int64 `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// 二维码对应的渠道名
	ChannelName string `json:"channel_name,omitempty" xml:"channel_name,omitempty"`
}
