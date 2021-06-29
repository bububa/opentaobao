package drugtrace

// PlantingInfoDto 
type PlantingInfoDto struct {

    // 种植面积
    
    PlantingArea   string `json:"planting_area,omitempty" xml:"planting_area,omitempty"`
    

    // 种植时间yyyy-MM-dd
    
    PlantingDate   string `json:"planting_date,omitempty" xml:"planting_date,omitempty"`
    

    // 农药使用
    
    PesticideUse   string `json:"pesticide_use,omitempty" xml:"pesticide_use,omitempty"`
    

    // 生物调节剂
    
    BiologicalRegulator   string `json:"biological_regulator,omitempty" xml:"biological_regulator,omitempty"`
    

    // 土壤类型
    
    SoilType   string `json:"soil_type,omitempty" xml:"soil_type,omitempty"`
    

    // 种植管理图片（上传图片）图片建议尺寸：height: 310px;width: 670px;
    
    PlantingPictures   []string `json:"planting_pictures,omitempty" xml:"planting_pictures>string,omitempty"`
    

    // 栽培模式
    
    CultivationMode   string `json:"cultivation_mode,omitempty" xml:"cultivation_mode,omitempty"`
    

}
