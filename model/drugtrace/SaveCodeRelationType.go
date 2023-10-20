package drugtrace

import (
	"sync"
)

// SaveCodeRelationType 结构体
type SaveCodeRelationType struct {
	// 操作的icCode
	OperIcCode string `json:"oper_ic_code,omitempty" xml:"oper_ic_code,omitempty"`
	// 用户证书
	UserCert string `json:"user_cert,omitempty" xml:"user_cert,omitempty"`
	// 生产码
	ProdCode string `json:"prod_code,omitempty" xml:"prod_code,omitempty"`
	// 企业名
	EntName string `json:"ent_name,omitempty" xml:"ent_name,omitempty"`
	// 操作员姓名
	OperIcName string `json:"oper_ic_name,omitempty" xml:"oper_ic_name,omitempty"`
	// 上传标记
	UploadFlag string `json:"upload_flag,omitempty" xml:"upload_flag,omitempty"`
	// 上传文件的文件名（建议填写一个，用于后面查询）
	UploadFileName string `json:"upload_file_name,omitempty" xml:"upload_file_name,omitempty"`
	// 上传日期(格式 2018-09-18)
	CrtDate string `json:"crt_date,omitempty" xml:"crt_date,omitempty"`
	// 上传文件的企业ID
	EntSeqNo string `json:"ent_seq_no,omitempty" xml:"ent_seq_no,omitempty"`
	// 1药  3中药饮片  5医疗器材
	BusinessType int64 `json:"business_type,omitempty" xml:"business_type,omitempty"`
}

var poolSaveCodeRelationType = sync.Pool{
	New: func() any {
		return new(SaveCodeRelationType)
	},
}

// GetSaveCodeRelationType() 从对象池中获取SaveCodeRelationType
func GetSaveCodeRelationType() *SaveCodeRelationType {
	return poolSaveCodeRelationType.Get().(*SaveCodeRelationType)
}

// ReleaseSaveCodeRelationType 释放SaveCodeRelationType
func ReleaseSaveCodeRelationType(v *SaveCodeRelationType) {
	v.OperIcCode = ""
	v.UserCert = ""
	v.ProdCode = ""
	v.EntName = ""
	v.OperIcName = ""
	v.UploadFlag = ""
	v.UploadFileName = ""
	v.CrtDate = ""
	v.EntSeqNo = ""
	v.BusinessType = 0
	poolSaveCodeRelationType.Put(v)
}
