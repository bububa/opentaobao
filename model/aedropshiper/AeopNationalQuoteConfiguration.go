package aedropshiper

// AeopNationalQuoteConfiguration 
type AeopNationalQuoteConfiguration struct {
    // 分国家定价规则类型[percentage：基于基准价格按比例配置]
    ConfigurationType   string `json:"configuration_type,omitempty" xml:"configuration_type,omitempty"`
    // jsonArray格式的分国家定价规则数据。 1)基于基准价格按比例配置的数据格式：[{"shiptoCountry":"US","percentage":"5"},{"shiptoCountry":"RU","percentage":"-2"}] 其中shiptoCountry：ISO两位的国家编码（目前支持11个国家：RU,US,CA,ES,FR,UK,NL,IL,BR,CL,AU）， percentage：相对于基准价的调价比例（百分比整数，支持负数，当前限制>=-30 && <=100）
    ConfigurationData   string `json:"configuration_data,omitempty" xml:"configuration_data,omitempty"`
}
