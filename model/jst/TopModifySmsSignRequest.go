package jst

// TopModifySmsSignRequest 结构体
type TopModifySmsSignRequest struct {
	// 上传的证明文件
	SignFileList []SmsFileContentDto `json:"sign_file_list,omitempty" xml:"sign_file_list>sms_file_content_dto,omitempty"`
	// 要修改的签名,不能修改签名
	SignName string `json:"sign_name,omitempty" xml:"sign_name,omitempty"`
	// 请输入签名用途（必填）、企业官网链接（可提升通过率）
	Remark string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 0--企事业单位的全程或简称   1--已备案网站  2--已上线APP  3--公众号或小程序 4--电商平台店铺名 5--已注册商标名
	SignSource int64 `json:"sign_source,omitempty" xml:"sign_source,omitempty"`
}
