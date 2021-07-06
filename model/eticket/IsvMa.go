package eticket

// IsvMa 结构体
type IsvMa struct {
	// 串码码值
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 二维码图片文件名。已经申请了上传二维码的码商必填，其它码商无需关心。这个值是taobao.eticket.merchant.img.upload调用后的file_name
	QrImage string `json:"qr_image,omitempty" xml:"qr_image,omitempty"`
	// 码的可核销份数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
}
