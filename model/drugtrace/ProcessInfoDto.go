package drugtrace

// ProcessInfoDto 
type ProcessInfoDto struct {
    // 采收日期yyyy-MM-dd
    HarvestDate   string `json:"harvest_date,omitempty" xml:"harvest_date,omitempty"`
    // 采收部位
    HarvestPosition   string `json:"harvest_position,omitempty" xml:"harvest_position,omitempty"`
    // 采收数量
    HarvestNum   string `json:"harvest_num,omitempty" xml:"harvest_num,omitempty"`
    // 初加工方法
    ProcessMethod   string `json:"process_method,omitempty" xml:"process_method,omitempty"`
    // 初加工设备
    ProcessMachine   string `json:"process_machine,omitempty" xml:"process_machine,omitempty"`
    // 产地初加工管理图片（上传图片）图片建议尺寸：height: 310px;width: 670px;
    HarvestPictures   []string `json:"harvest_pictures,omitempty" xml:"harvest_pictures>string,omitempty"`
    // 采收地块
    HarvestedPlot   string `json:"harvested_plot,omitempty" xml:"harvested_plot,omitempty"`
}
