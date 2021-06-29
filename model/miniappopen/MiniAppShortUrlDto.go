package miniappopen

// MiniAppShortUrlDTO 
type MiniAppShortUrlDTO struct {
    // 短链接地址【特别注意：短链接有效期为30天，超过时效短链接将无法正常跳转到原始链接地址，请勿将短链接投放或装修到长期存在的入口】
    ShortUrl   string `json:"short_url,omitempty" xml:"short_url,omitempty"`
}
