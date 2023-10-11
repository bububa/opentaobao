package axindata

// FscProjectAddApiResponse 结构体
type FscProjectAddApiResponse struct {
	// 团期计划ID映射
	ProjectMapDTOS []FscProjectMapDto `json:"project_map_d_t_o_s,omitempty" xml:"project_map_d_t_o_s>fsc_project_map_dto,omitempty"`
}
