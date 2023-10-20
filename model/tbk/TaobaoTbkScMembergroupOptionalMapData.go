package tbk

// TaobaotbkscmembergroupoptionalMapData 结构体
type TaobaotbkscmembergroupoptionalMapData struct {
	// 组内的成员ID
	Tbnumids string `json:"tb_num_ids,omitempty" xml:"tb_num_ids,omitempty"`
	// 创建时间
	Createtime string `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 更新时间
	Updatetime string `json:"update_time,omitempty" xml:"update_time,omitempty"`
	// 组ID
	Membergroupid int64 `json:"member_group_id,omitempty" xml:"member_group_id,omitempty"`
}
