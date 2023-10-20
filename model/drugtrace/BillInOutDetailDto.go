package drugtrace

import (
	"sync"
)

// BillInOutDetailDto 结构体
type BillInOutDetailDto struct {
	// 单据详情
	BillChkInOutDetailListDTOList []Billchkinoutdetaillistdtolist `json:"bill_chk_in_out_detail_list_d_t_o_list,omitempty" xml:"bill_chk_in_out_detail_list_d_t_o_list>billchkinoutdetaillistdtolist,omitempty"`
	// 单据中的码
	Codes []string `json:"codes,omitempty" xml:"codes>string,omitempty"`
	// 修改时间
	ModDate string `json:"mod_date,omitempty" xml:"mod_date,omitempty"`
	// 处理时间
	ProcessDate string `json:"process_date,omitempty" xml:"process_date,omitempty"`
	// 单据日期
	BillTime string `json:"bill_time,omitempty" xml:"bill_time,omitempty"`
	// 收货企业id
	ToUserId string `json:"to_user_id,omitempty" xml:"to_user_id,omitempty"`
	// 收货企业名称
	ToEntName string `json:"to_ent_name,omitempty" xml:"to_ent_name,omitempty"`
	// 发货企业id
	FromUserId string `json:"from_user_id,omitempty" xml:"from_user_id,omitempty"`
	// 发货企业名称
	FromEntName string `json:"from_ent_name,omitempty" xml:"from_ent_name,omitempty"`
	// 单据类型名称
	BillTypeName string `json:"bill_type_name,omitempty" xml:"bill_type_name,omitempty"`
	// 单据类型
	BillType string `json:"bill_type,omitempty" xml:"bill_type,omitempty"`
	// 单据号码
	BillCode string `json:"bill_code,omitempty" xml:"bill_code,omitempty"`
	// 上传文件名称
	UploadFileName string `json:"upload_file_name,omitempty" xml:"upload_file_name,omitempty"`
}

var poolBillInOutDetailDto = sync.Pool{
	New: func() any {
		return new(BillInOutDetailDto)
	},
}

// GetBillInOutDetailDto() 从对象池中获取BillInOutDetailDto
func GetBillInOutDetailDto() *BillInOutDetailDto {
	return poolBillInOutDetailDto.Get().(*BillInOutDetailDto)
}

// ReleaseBillInOutDetailDto 释放BillInOutDetailDto
func ReleaseBillInOutDetailDto(v *BillInOutDetailDto) {
	v.BillChkInOutDetailListDTOList = v.BillChkInOutDetailListDTOList[:0]
	v.Codes = v.Codes[:0]
	v.ModDate = ""
	v.ProcessDate = ""
	v.BillTime = ""
	v.ToUserId = ""
	v.ToEntName = ""
	v.FromUserId = ""
	v.FromEntName = ""
	v.BillTypeName = ""
	v.BillType = ""
	v.BillCode = ""
	v.UploadFileName = ""
	poolBillInOutDetailDto.Put(v)
}
