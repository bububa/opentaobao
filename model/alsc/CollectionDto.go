package alsc

// CollectionDto 结构体
type CollectionDto struct {
	// 任务关系配置信息
	RelationConfigs []RelationConfigDto `json:"relation_configs,omitempty" xml:"relation_configs>relation_config_dto,omitempty"`
}
