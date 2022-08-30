package damaiticklet

// ScriptInfoOpenParam 结构体
type ScriptInfoOpenParam struct {
	// 剧本标签数据
	ThirdScriptTagList []ScriptTagThirdParam `json:"third_script_tag_list,omitempty" xml:"third_script_tag_list>script_tag_third_param,omitempty"`
	// 剧本其他图片
	ScriptOtherImageList []string `json:"script_other_image_list,omitempty" xml:"script_other_image_list>string,omitempty"`
	// 剧本视频
	ScriptVideoList []string `json:"script_video_list,omitempty" xml:"script_video_list>string,omitempty"`
	// 剧本人物海报图
	ScriptPosterList []string `json:"script_poster_list,omitempty" xml:"script_poster_list>string,omitempty"`
	// 发行时间
	ReleaseTime string `json:"release_time,omitempty" xml:"release_time,omitempty"`
	// 是否全息 是1 否0
	HasHolographic string `json:"has_holographic,omitempty" xml:"has_holographic,omitempty"`
	// 剧本文字描述
	ScriptDescribe string `json:"script_describe,omitempty" xml:"script_describe,omitempty"`
	// 监制名称
	SuperviseName string `json:"supervise_name,omitempty" xml:"supervise_name,omitempty"`
	// 有无道具 有1 无0
	HasProp string `json:"has_prop,omitempty" xml:"has_prop,omitempty"`
	// 电子资料
	ElectronicData string `json:"electronic_data,omitempty" xml:"electronic_data,omitempty"`
	// 剧本封面
	ScriptCover string `json:"script_cover,omitempty" xml:"script_cover,omitempty"`
	// 作者
	Author string `json:"author,omitempty" xml:"author,omitempty"`
	// 可否反串 1可 0否
	IsReverse string `json:"is_reverse,omitempty" xml:"is_reverse,omitempty"`
	// 剧本名称
	ScriptName string `json:"script_name,omitempty" xml:"script_name,omitempty"`
	// 出品方
	Producer string `json:"producer,omitempty" xml:"producer,omitempty"`
	// 发行工作室
	DistributionStudio string `json:"distribution_studio,omitempty" xml:"distribution_studio,omitempty"`
	// 有无演绎
	HasDeduction string `json:"has_deduction,omitempty" xml:"has_deduction,omitempty"`
	// 玩家人数下限
	PlayerMinNum int64 `json:"player_min_num,omitempty" xml:"player_min_num,omitempty"`
	// 剧本来源 剧游:2
	Source int64 `json:"source,omitempty" xml:"source,omitempty"`
	// 女生人数
	GirlNum int64 `json:"girl_num,omitempty" xml:"girl_num,omitempty"`
	// 阅读体量
	ReadingVolume int64 `json:"reading_volume,omitempty" xml:"reading_volume,omitempty"`
	// 主持难度
	HostingDifficulty int64 `json:"hosting_difficulty,omitempty" xml:"hosting_difficulty,omitempty"`
	// 男生人数
	BoyNum int64 `json:"boy_num,omitempty" xml:"boy_num,omitempty"`
	// 三方剧本id
	OutId int64 `json:"out_id,omitempty" xml:"out_id,omitempty"`
	// 发售价格（单位分）
	SalePrice int64 `json:"sale_price,omitempty" xml:"sale_price,omitempty"`
	// 发售方式 盒装1 限定2 独家3 待定4  实景5
	SaleMode int64 `json:"sale_mode,omitempty" xml:"sale_mode,omitempty"`
	// 平均游戏时长（分钟）
	ScriptTime int64 `json:"script_time,omitempty" xml:"script_time,omitempty"`
	// 剧本难度 1:新手 2:进阶 3:烧脑
	Difficulty int64 `json:"difficulty,omitempty" xml:"difficulty,omitempty"`
	// 玩家人数模式：固定:0 范围1
	PlayerNumMode int64 `json:"player_num_mode,omitempty" xml:"player_num_mode,omitempty"`
	// 玩家人数上限
	PlayerMaxNum int64 `json:"player_max_num,omitempty" xml:"player_max_num,omitempty"`
	// 1:6+ ，2:12+ ，3:16+ ，4:18+
	RightAge int64 `json:"right_age,omitempty" xml:"right_age,omitempty"`
}
