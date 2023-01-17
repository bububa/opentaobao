package paimai

// NftCertificateApplyCallbackDto 结构体
type NftCertificateApplyCallbackDto struct {
	// 业务ID申请证书编号时记录 回调传入
	BizOuterId string `json:"biz_outer_id,omitempty" xml:"biz_outer_id,omitempty"`
	// 著作权人证件号
	CopyrightOwnerNumber string `json:"copyright_owner_number,omitempty" xml:"copyright_owner_number,omitempty"`
	// 证书图片下载地址
	CertificatePic string `json:"certificate_pic,omitempty" xml:"certificate_pic,omitempty"`
	// 证书号
	CertificateNumber string `json:"certificate_number,omitempty" xml:"certificate_number,omitempty"`
	// 著作权人姓名
	CopyrightOwnerName string `json:"copyright_owner_name,omitempty" xml:"copyright_owner_name,omitempty"`
	// 证书申请号 申请证书编号时记录 回调传入
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// 证书颁发时间
	CertificateReleaseTime string `json:"certificate_release_time,omitempty" xml:"certificate_release_time,omitempty"`
}
