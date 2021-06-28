package ma

// QrCodeStyle 
type QrCodeStyle struct {

    // 可选参数，二维背景色颜色值，接入业务点配置为准，未配置，默认为白
    
    BgColor   int64 `json:"bg_color,omitempty" xml:"bg_color,omitempty"`
    

    // 可选参数，logo的淘宝tfs地址，默认无
    
    Logo   string `json:"logo,omitempty" xml:"logo,omitempty"`
    

    // 可选参数，二维码纠错级别 0=~7%,1=~15%,2=~25%,3=~30%
    
    Level   int64 `json:"level,omitempty" xml:"level,omitempty"`
    

    // 可选参数，二维码深色点颜色值，接入业务点配置为准，未配置，默认为黑
    
    Color   int64 `json:"color,omitempty" xml:"color,omitempty"`
    

    // 可选参数，二维码的边框，默认大小1个单位点，便于扫码
    
    Margin   int64 `json:"margin,omitempty" xml:"margin,omitempty"`
    

    // 可选参数，二维码大小，值60～600，默认185pix
    
    Size   int64 `json:"size,omitempty" xml:"size,omitempty"`
    

}
