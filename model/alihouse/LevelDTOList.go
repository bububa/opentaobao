package alihouse

import (
	"sync"
)

// LevelDTOList 结构体
type LevelDTOList struct {
	// 权益列表
	GradeItemDTOList []GradeItemDTOList `json:"grade_item_d_t_o_list,omitempty" xml:"grade_item_d_t_o_list>grade_item_dto_list,omitempty"`
	// 等级编号
	LevelCode string `json:"level_code,omitempty" xml:"level_code,omitempty"`
	// 外部门店ID
	OuterStoreId string `json:"outer_store_id,omitempty" xml:"outer_store_id,omitempty"`
	// 等级名称
	LevelName string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	// 总分
	TotalScore int64 `json:"total_score,omitempty" xml:"total_score,omitempty"`
}

var poolLevelDTOList = sync.Pool{
	New: func() any {
		return new(LevelDTOList)
	},
}

// GetLevelDTOList() 从对象池中获取LevelDTOList
func GetLevelDTOList() *LevelDTOList {
	return poolLevelDTOList.Get().(*LevelDTOList)
}

// ReleaseLevelDTOList 释放LevelDTOList
func ReleaseLevelDTOList(v *LevelDTOList) {
	v.GradeItemDTOList = v.GradeItemDTOList[:0]
	v.LevelCode = ""
	v.OuterStoreId = ""
	v.LevelName = ""
	v.TotalScore = 0
	poolLevelDTOList.Put(v)
}
