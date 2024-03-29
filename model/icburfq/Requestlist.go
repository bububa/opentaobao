package icburfq

import (
	"sync"
)

// Requestlist 结构体
type Requestlist struct {
	// 附件
	AnnexFiles []RfqAnnexFileRemoteDto `json:"annex_files,omitempty" xml:"annex_files>rfq_annex_file_remote_dto,omitempty"`
	// RFQID
	RfqId string `json:"rfq_id,omitempty" xml:"rfq_id,omitempty"`
	// RFQ标题
	Subject string `json:"subject,omitempty" xml:"subject,omitempty"`
	// RFQ内容
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 数量单位
	QuantityUnit string `json:"quantity_unit,omitempty" xml:"quantity_unit,omitempty"`
	// 国家简称
	CountrySimple string `json:"country_simple,omitempty" xml:"country_simple,omitempty"`
	// 附件名称
	AnnexNames string `json:"annex_names,omitempty" xml:"annex_names,omitempty"`
	// 语种
	LangSrc string `json:"lang_src,omitempty" xml:"lang_src,omitempty"`
	// 图片地址
	ImageUrl string `json:"image_url,omitempty" xml:"image_url,omitempty"`
	// 供应商国家
	SupplierCountrys string `json:"supplier_countrys,omitempty" xml:"supplier_countrys,omitempty"`
	// 唯一加密RFQID
	UniqueRfqId string `json:"unique_rfq_id,omitempty" xml:"unique_rfq_id,omitempty"`
	// 类目ID
	CategoryId int64 `json:"category_id,omitempty" xml:"category_id,omitempty"`
	// 数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 剩余报价
	LeftCount int64 `json:"left_count,omitempty" xml:"left_count,omitempty"`
	// 过期时间
	ExpirateTime int64 `json:"expirate_time,omitempty" xml:"expirate_time,omitempty"`
	// 开始时间
	OpenTime int64 `json:"open_time,omitempty" xml:"open_time,omitempty"`
}

var poolRequestlist = sync.Pool{
	New: func() any {
		return new(Requestlist)
	},
}

// GetRequestlist() 从对象池中获取Requestlist
func GetRequestlist() *Requestlist {
	return poolRequestlist.Get().(*Requestlist)
}

// ReleaseRequestlist 释放Requestlist
func ReleaseRequestlist(v *Requestlist) {
	v.AnnexFiles = v.AnnexFiles[:0]
	v.RfqId = ""
	v.Subject = ""
	v.Description = ""
	v.QuantityUnit = ""
	v.CountrySimple = ""
	v.AnnexNames = ""
	v.LangSrc = ""
	v.ImageUrl = ""
	v.SupplierCountrys = ""
	v.UniqueRfqId = ""
	v.CategoryId = 0
	v.Quantity = 0
	v.LeftCount = 0
	v.ExpirateTime = 0
	v.OpenTime = 0
	poolRequestlist.Put(v)
}
