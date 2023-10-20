package drugtrace

// AlibabaalihealthdrugkytybgetcoderelationModel 结构体
type AlibabaalihealthdrugkytybgetcoderelationModel struct {
	// 码关联关系DTO
	CodeRelationDtoList []CodeRelationDto `json:"code_relation_dto_list,omitempty" xml:"code_relation_dto_list>code_relation_dto,omitempty"`
	// 包装比例
	PkgRatio string `json:"pkg_ratio,omitempty" xml:"pkg_ratio,omitempty"`
}
