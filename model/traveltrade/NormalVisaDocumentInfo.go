package traveltrade

import (
	"sync"
)

// NormalVisaDocumentInfo 结构体
type NormalVisaDocumentInfo struct {
	// 特殊必填，上传子材料，如上传护照封面
	SubDocumentInfos []NormalVisaSubDocumentInfo `json:"sub_document_infos,omitempty" xml:"sub_document_infos>normal_visa_sub_document_info,omitempty"`
	// 文件类型
	FileType string `json:"file_type,omitempty" xml:"file_type,omitempty"`
	// base64编码的文件内容
	FileContent string `json:"file_content,omitempty" xml:"file_content,omitempty"`
	// 文档编号，0:护照，1:证件照，2:申请表，3:身份证，4:户口本，5:暂住证，6:在职收入证明，7:营业执照，8:组织机构代码证，9:结婚证，10:个人信息处理同意书，11:退休证，12:保险订单，13:在读证明，14:机票预订证明，15:酒店预订证明，16:财力证明，17:房产证，18:汽车驾驶证，19:社保缴纳记录，20:学校准假证明，21:儿童出生医学证明，22:未成年人亲属关系证明，23:其他材料，24:银行存款证明，25:学生证，26:其他材料2，27:其他材料3，28:居住证，29:车辆登记证，34:保险声明，36:出行同意书，38:职业证明，39:以往申根签证页
	DocType int64 `json:"doc_type,omitempty" xml:"doc_type,omitempty"`
}

var poolNormalVisaDocumentInfo = sync.Pool{
	New: func() any {
		return new(NormalVisaDocumentInfo)
	},
}

// GetNormalVisaDocumentInfo() 从对象池中获取NormalVisaDocumentInfo
func GetNormalVisaDocumentInfo() *NormalVisaDocumentInfo {
	return poolNormalVisaDocumentInfo.Get().(*NormalVisaDocumentInfo)
}

// ReleaseNormalVisaDocumentInfo 释放NormalVisaDocumentInfo
func ReleaseNormalVisaDocumentInfo(v *NormalVisaDocumentInfo) {
	v.SubDocumentInfos = v.SubDocumentInfos[:0]
	v.FileType = ""
	v.FileContent = ""
	v.DocType = 0
	poolNormalVisaDocumentInfo.Put(v)
}
