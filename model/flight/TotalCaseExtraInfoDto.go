package flight

import (
	"sync"
)

// TotalCaseExtraInfoDto 结构体
type TotalCaseExtraInfoDto struct {
	// 加订婴童问题类型extraInfo入参
	CaseExtraAddBabyRequestDtoList []CaseExtraAddBabyRequestDto `json:"case_extra_add_baby_request_dto_list,omitempty" xml:"case_extra_add_baby_request_dto_list>case_extra_add_baby_request_dto,omitempty"`
	// 二次回填问题类型extraInfo入参
	CaseRepeatBackFillExtraInfoRequestDtoList []CaseRepeatBackFillExtraInfoRequestDto `json:"case_repeat_back_fill_extra_info_request_dto_list,omitempty" xml:"case_repeat_back_fill_extra_info_request_dto_list>case_repeat_back_fill_extra_info_request_dto,omitempty"`
	// 改名改证件问题类型extraInfo入参
	CaseChangePassengerExtraInfoRequestDtoList []CaseChangePassengerExtraInfoRequestDto `json:"case_change_passenger_extra_info_request_dto_list,omitempty" xml:"case_change_passenger_extra_info_request_dto_list>case_change_passenger_extra_info_request_dto,omitempty"`
}

var poolTotalCaseExtraInfoDto = sync.Pool{
	New: func() any {
		return new(TotalCaseExtraInfoDto)
	},
}

// GetTotalCaseExtraInfoDto() 从对象池中获取TotalCaseExtraInfoDto
func GetTotalCaseExtraInfoDto() *TotalCaseExtraInfoDto {
	return poolTotalCaseExtraInfoDto.Get().(*TotalCaseExtraInfoDto)
}

// ReleaseTotalCaseExtraInfoDto 释放TotalCaseExtraInfoDto
func ReleaseTotalCaseExtraInfoDto(v *TotalCaseExtraInfoDto) {
	v.CaseExtraAddBabyRequestDtoList = v.CaseExtraAddBabyRequestDtoList[:0]
	v.CaseRepeatBackFillExtraInfoRequestDtoList = v.CaseRepeatBackFillExtraInfoRequestDtoList[:0]
	v.CaseChangePassengerExtraInfoRequestDtoList = v.CaseChangePassengerExtraInfoRequestDtoList[:0]
	poolTotalCaseExtraInfoDto.Put(v)
}
